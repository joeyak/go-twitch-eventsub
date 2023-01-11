package twitch_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/joeyak/go-twitch-eventsub"
)

func noDataGen() ([]byte, bool, error) {
	return nil, false, nil
}

func TestNoWelcome(t *testing.T) {
	t.Parallel()

	client := twitch.NewClientWithUrl("")
	err := client.Connect()
	if !errors.Is(err, twitch.ErrNilOnWelcome) {
		t.Fatalf("expected ErrNilOnWelcome, actual %#v", err)
	}
}

func TestClose(t *testing.T) {
	t.Parallel()
	client := newClient(t, noDataGen)
	go func() {
		time.Sleep(50 * time.Millisecond)
		client.Close()
	}()

	err := client.Connect()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestCloseWithContext(t *testing.T) {
	t.Parallel()
	client := newClient(t, noDataGen)

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	err := client.ConnectWithContext(ctx)
	if err != nil && !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestKeepAlive(t *testing.T) {
	t.Parallel()
	client := newClient(t, func() ([]byte, bool, error) {
		return []byte(`{
			"metadata": {
				"message_id": "84c1e79a-2a4b-4c13-ba0b-4312293e9308",
				"message_type": "session_keepalive",
				"message_timestamp": "2019-11-16T10:11:12.634234626Z"
			},
			"payload": {}
		}`), false, nil
	})

	ch := make(chan struct{})
	client.OnKeepAlive(func(message twitch.KeepAliveMessage) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("keepalive did not occur")
	}
}

func TestReconnect(t *testing.T) {
	t.Parallel()
	client := newClient(t, func() ([]byte, bool, error) {
		return []byte(`{
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
					"reconnect_url": "wss://eventsub-beta.wss.twitch.tv?...",
					"connected_at": "2019-11-16T10:11:12.634234626Z"
				}
			}
		}`), false, nil
	})

	ch := make(chan struct{})
	client.OnReconnect(func(message twitch.ReconnectMessage) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("reconnect did not occur")
	}
}

func TestRevoke(t *testing.T) {
	t.Parallel()
	client := newClient(t, func() ([]byte, bool, error) {
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
	})

	ch := make(chan struct{})
	client.OnRevoke(func(message twitch.RevokeMessage) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("revoke did not occur")
	}
}

func TestOnError(t *testing.T) {
	t.Parallel()
	client := newClient(t, func() ([]byte, bool, error) {
		return []byte(`{}`), false, nil
	})

	ch := make(chan struct{})
	client.OnError(func(err error) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("onerror did not occur")
	}
}

func TestInvalidJson(t *testing.T) {
	t.Parallel()
	client := newClient(t, func() ([]byte, bool, error) {
		return []byte(`{`), false, nil
	})

	ch := make(chan struct{})
	client.OnError(func(err error) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("onerror did not occur")
	}
}
