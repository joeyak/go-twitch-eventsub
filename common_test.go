package twitch_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"
	"time"

	_ "embed"

	"github.com/google/uuid"
	"github.com/joeyak/go-twitch-eventsub/v2"
	"nhooyr.io/websocket"
)

//go:embed testEvents.json
var testEvents []byte

type messageDataGenerator func() ([][]byte, bool, error)

func getTestEventData(eventType twitch.EventSubscription, suffixes ...string) messageDataGenerator {
	return func() ([][]byte, bool, error) {
		var events map[string]json.RawMessage
		if err := json.Unmarshal(testEvents, &events); err != nil {
			return nil, false, fmt.Errorf("could not parse event json file: %w", err)
		}

		key := strings.Join(append([]string{string(eventType)}, suffixes...), "-")
		eventData, ok := events[key]
		if !ok {
			return nil, false, fmt.Errorf("could not find %s in testEvents", key)
		}

		data, err := json.Marshal(twitch.NotificationMessage{
			Metadata: newMetadata("notification"),
			Payload: struct {
				Subscription twitch.PayloadSubscription "json:\"subscription\""
				Event        *json.RawMessage           "json:\"event\""
			}{
				Event: &eventData,
				Subscription: twitch.PayloadSubscription{
					SubscriptionRequest: twitch.SubscriptionRequest{
						Type:      eventType,
						Version:   "1",
						Condition: map[string]string{},
						Transport: twitch.SubscriptionTransport{
							Method:    "websocket",
							SessionID: "",
						},
					},
					Status:   "enabled",
					Cost:     1,
					CreateAt: time.Now(),
				},
			},
		})
		return [][]byte{data}, true, err
	}
}

type TestServer struct {
	Address            string
	conn               *websocket.Conn
	sendInSubscription bool
	data               [][]byte
}

func newTestServer(gen messageDataGenerator) (TestServer, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return TestServer{}, fmt.Errorf("could not listen on random port: %w", err)
	}

	data, sendInSubscription, err := gen()
	if err != nil {
		return TestServer{}, fmt.Errorf("could not get generate message data: %w", err)
	}

	for i := range data {
		for _, r := range "\t\r\n" {
			data[i] = bytes.ReplaceAll(data[i], []byte{byte(r)}, nil)
		}
	}

	server := TestServer{
		Address:            listener.Addr().String(),
		sendInSubscription: sendInSubscription,
		data:               data,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", server.handleWebsocket)
	if server.sendInSubscription {
		mux.HandleFunc("/subscriptions", server.handleSubscription)
	}

	go http.Serve(listener, mux)
	return server, nil
}

func (s *TestServer) handleWebsocket(w http.ResponseWriter, r *http.Request) {
	var err error
	s.conn, err = websocket.Accept(w, r, nil)
	if err != nil {
		panic(err)
	}

	err = s.sendWelcome(r.Context())
	if err != nil {
		panic(err)
	}

	if !s.sendInSubscription {
		for _, data := range s.data {
			s.conn.Write(r.Context(), websocket.MessageText, data)
		}
	}

	// Read so it can close
	s.conn.Read(r.Context())
}

func (s *TestServer) handleSubscription(w http.ResponseWriter, r *http.Request) {
	request, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	r.Body.Close()

	var subscription twitch.SubscriptionRequest
	err = json.Unmarshal(request, &subscription)
	if err != nil {
		panic(err)
	}

	response, _ := json.Marshal(twitch.SubscribeResponse{})
	w.WriteHeader(http.StatusAccepted)
	w.Write(response)

	for _, data := range s.data {
		err = s.conn.Write(r.Context(), websocket.MessageText, data)
		if err != nil {
			panic(err)
		}
	}
}

func (s *TestServer) sendWelcome(ctx context.Context) error {
	welcome := twitch.WelcomeMessage{
		Metadata: newMetadata("session_welcome"),
		Payload: struct {
			Session twitch.PayloadSession `json:"session"`
		}{
			Session: twitch.PayloadSession{
				ID:                      strings.ReplaceAll(uuid.NewString(), "-", ""),
				Status:                  "connected",
				ConnectedAt:             time.Now(),
				KeepaliveTimeoutSeconds: 10,
				ReconnectUrl:            "",
			},
		},
	}

	data, err := json.Marshal(welcome)
	if err != nil {
		return fmt.Errorf("could not marshal welcome message: %w", err)
	}

	return s.conn.Write(ctx, websocket.MessageText, data)
}

func newMetadata(msgType string) twitch.MessageMetadata {
	return twitch.MessageMetadata{
		MessageID:        uuid.NewString(),
		MessageType:      msgType,
		MessageTimestamp: time.Now(),
	}
}

func newClient(t *testing.T, gen messageDataGenerator) *twitch.Client {
	server, err := newTestServer(gen)
	if err != nil {
		t.Fatal(err)
	}

	client := twitch.NewClientWithUrl(fmt.Sprintf("http://%s/%s", server.Address, "ws"))
	client.OnError(func(err error) {
		t.Fatalf("client registered an error: %v", err)
	})
	client.OnWelcome(func(message twitch.WelcomeMessage) {})

	return client
}

func newClientWithWelcome(t *testing.T, version string, event twitch.EventSubscription, gen messageDataGenerator) *twitch.Client {
	client := newClient(t, gen)

	client.OnWelcome(func(message twitch.WelcomeMessage) {
		_, err := twitch.SubscribeEventUrl(twitch.SubscribeRequest{
			SessionID:       message.Payload.Session.ID,
			ClientID:        "",
			AccessToken:     "",
			VersionOverride: version,
			Event:           event,
			Condition:       map[string]string{},
		}, strings.ReplaceAll(client.Address, "/ws", "/subscriptions"))
		if err != nil {
			t.Errorf("could not subscribe: %v", err)
		}
	})
	return client
}

func connect(t *testing.T, client *twitch.Client) {
	err := client.Connect()
	if err != nil {
		t.Errorf("could not connect client: %v", err)
	}
}
