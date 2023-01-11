package twitch_test

import (
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
	"github.com/joeyak/go-twitch-eventsub"
	"nhooyr.io/websocket"
)

//go:embed testEvents.json
var testEvents []byte

func getTestEventData(eventType twitch.EventSubscription, suffixes ...string) func() ([]byte, error) {
	return func() ([]byte, error) {
		var events map[string]json.RawMessage
		if err := json.Unmarshal(testEvents, &events); err != nil {
			return nil, fmt.Errorf("could not parse event json file: %w", err)
		}

		key := strings.Join(append([]string{string(eventType)}, suffixes...), "-")
		eventData, ok := events[key]
		if !ok {
			return nil, fmt.Errorf("could not find %s in testEvents", key)
		}

		return json.Marshal(twitch.NotificationMessage{
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
	}
}

type messageDataGenerator func() ([]byte, error)

type TestServer struct {
	Address string
	conn    *websocket.Conn
}

func newTestServer(gen messageDataGenerator) (TestServer, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return TestServer{}, fmt.Errorf("could not listen on random port: %w", err)
	}

	server := TestServer{Address: listener.Addr().String()}

	data, err := gen()
	if err != nil {
		return TestServer{}, fmt.Errorf("could not get generate message data: %w", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/ws", server.handleWebsocket)
	mux.HandleFunc("/subscriptions", server.handleSubscription(data))

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

	// Read so it can close
	s.conn.Read(r.Context())
}

func (s *TestServer) handleSubscription(notification []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		err = s.conn.Write(r.Context(), websocket.MessageText, notification)
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

func newClient(t *testing.T, event twitch.EventSubscription, gen messageDataGenerator) *twitch.Client {
	server, err := newTestServer(gen)
	if err != nil {
		t.Fatal(err)
	}

	client := twitch.NewClientWithUrl(fmt.Sprintf("http://%s/%s", server.Address, "ws"))
	client.OnError(func(err error) {
		t.Fatalf("client registered an error: %v", err)
	})
	client.OnWelcome(func(message twitch.WelcomeMessage) {
		_, err := twitch.SubscribeEventUrl(twitch.SubscribeRequest{
			SessionID:   message.Payload.Session.ID,
			ClientID:    "",
			AccessToken: "",
			Event:       event,
			Condition:   map[string]string{},
		}, fmt.Sprintf("http://%s/%s", server.Address, "subscriptions"))
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
