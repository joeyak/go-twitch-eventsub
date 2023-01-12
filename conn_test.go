package twitch_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/joeyak/go-twitch-eventsub"
	"github.com/stretchr/testify/assert"
)

func noDataGen() ([]byte, bool, error) {
	return nil, false, nil
}

func keepAliveGen() ([]byte, bool, error) {
	return []byte(`{
		"metadata": {
			"message_id": "84c1e79a-2a4b-4c13-ba0b-4312293e9308",
			"message_type": "session_keepalive",
			"message_timestamp": "2019-11-16T10:11:12.634234626Z"
		},
		"payload": {}
	}`), false, nil
}

func revokeGen() ([]byte, bool, error) {
	return []byte(`{
		"metadata": {
			"message_id": "84c1e79a-2a4b-4c13-ba0b-4312293e9308",
			"message_type": "revocation",
			"message_timestamp": "2019-11-16T10:11:12.464757833Z",
			"subscription_type": "channel.follow",
			"subscription_version": "1"
		},
		"payload": {
			"subscription": {
				"id": "f1c2a387-161a-49f9-a165-0f21d7a4e1c4",
				"status": "authorization_revoked",
				"type": "channel.follow",
				"version": "1",
				"cost": 1,
				"condition": {
					"broadcaster_user_id": "12826"
				},
				"transport": {
					"method": "websocket",
					"session_id": "AQoQexAWVYKSTIu4ec_2VAxyuhAB"
				},
				"created_at": "2019-11-16T10:11:12.464757833Z"
			}
		}
	}`), false, nil
}

func genReconnectGen(url string) func() ([]byte, bool, error) {
	return func() ([]byte, bool, error) {
		return []byte(fmt.Sprintf(`{
		"metadata": {
			"message_id": "84c1e79a-2a4b-4c13-ba0b-4312293e9308",
			"message_type": "session_reconnect",
			"message_timestamp": "2019-11-18T09:10:11.634234626Z"
		},
		"payload": {
			"session": {
				"id": "AQoQexAWVYKSTIu4ec_2VAxyuhAB",
				"status": "reconnecting",
				"keepalive_timeout_seconds": null,
				"reconnect_url": "%s",
				"connected_at": "2019-11-16T10:11:12.634234626Z"
			}
		}
	}`, url)), false, nil
	}
}

func assertEventOccured(t *testing.T, f func(ch chan struct{})) {
	ch := make(chan struct{})

	f(ch)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestNoWelcome(t *testing.T) {
	t.Parallel()

	client := twitch.NewClientWithUrl("")
	err := client.Connect()
	assert.ErrorIs(t, err, twitch.ErrNilOnWelcome)
}

func TestOnClose(t *testing.T) {
	t.Parallel()
	client := newClient(t, noDataGen)

	go func() {
		time.Sleep(50 * time.Millisecond)
		client.Close()
	}()

	err := client.Connect()
	assert.NoError(t, err)
}

func TestOnCloseWithContext(t *testing.T) {
	t.Parallel()
	client := newClient(t, noDataGen)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	err := client.ConnectWithContext(ctx)

	assert.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestOnKeepAlive(t *testing.T) {
	t.Parallel()

	assertEventOccured(t, func(ch chan struct{}) {
		client := newClient(t, keepAliveGen)
		client.OnKeepAlive(func(message twitch.KeepAliveMessage) {
			close(ch)
		})

		go connect(t, client)
	})
}

func TestOnReconnect(t *testing.T) {
	t.Parallel()

	assertEventOccured(t, func(ch chan struct{}) {
		client := newClient(t, genReconnectGen(""))
		client.OnReconnect(func(message twitch.ReconnectMessage) {
			close(ch)
		})

		go connect(t, client)
	})
}

func TestOnRevoke(t *testing.T) {
	t.Parallel()

	assertEventOccured(t, func(ch chan struct{}) {
		client := newClient(t, revokeGen)
		client.OnRevoke(func(message twitch.RevokeMessage) {
			close(ch)
		})

		go connect(t, client)
	})
}

func TestOnError(t *testing.T) {
	t.Parallel()

	assertEventOccured(t, func(ch chan struct{}) {
		client := newClient(t, func() ([]byte, bool, error) {
			return []byte(`{}`), false, nil
		})
		client.OnError(func(err error) {
			close(ch)
		})

		go connect(t, client)
	})
}

func TestInvalidJson(t *testing.T) {
	t.Parallel()

	assertEventOccured(t, func(ch chan struct{}) {
		client := newClient(t, func() ([]byte, bool, error) {
			return []byte(`{`), false, nil
		})
		client.OnError(func(err error) {
			close(ch)
		})

		go connect(t, client)
	})
}

func TestReconnectInEvent(t *testing.T) {
	t.Parallel()

	reconnectServer, err := newTestServer(keepAliveGen)
	if err != nil {
		t.Fatalf("could not create reconnect server: %v", err)
	}
	reconnectUrl := fmt.Sprintf("http://%s/%s", reconnectServer.Address, "ws")

	client := newClient(t, genReconnectGen(reconnectUrl))
	client.OnKeepAlive(func(message twitch.KeepAliveMessage) {
		client.Close()
	})
	client.OnReconnect(func(message twitch.ReconnectMessage) {
		err := client.Reconnect(message.Payload.Session.ReconnectUrl)
		if err != nil {
			t.Errorf("could not reconnect: %v", err)
		}
	})

	err = client.Connect()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if client.Address != reconnectUrl {
		t.Fatalf("expected Address %s, got %s", reconnectUrl, client.Address)
	}
}

func TestReconnectOutsideEvent(t *testing.T) {
	t.Parallel()

	reconnectServer, err := newTestServer(keepAliveGen)
	if err != nil {
		t.Fatalf("could not create reconnect server: %v", err)
	}
	reconnectUrl := fmt.Sprintf("http://%s/%s", reconnectServer.Address, "ws")

	client := newClient(t, noDataGen)
	client.OnKeepAlive(func(message twitch.KeepAliveMessage) {
		client.Close()
	})

	go func() {
		time.Sleep(50 * time.Millisecond)
		err := client.Reconnect(reconnectUrl)
		if err != nil {
			t.Error(err)
		}
	}()

	err = client.Connect()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if client.Address != reconnectUrl {
		t.Fatalf("expected Address %s, got %s", reconnectUrl, client.Address)
	}
}
