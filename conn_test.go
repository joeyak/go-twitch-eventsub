package twitch_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/joeyak/go-twitch-eventsub/v2"
	"github.com/stretchr/testify/assert"
)

func noDataGen() ([][]byte, bool, error) {
	return nil, false, nil
}

func keepAliveGen() ([][]byte, bool, error) {
	return [][]byte{[]byte(`{
		"metadata": {
			"message_id": "84c1e79a-2a4b-4c13-ba0b-4312293e9308",
			"message_type": "session_keepalive",
			"message_timestamp": "2019-11-16T10:11:12.634234626Z"
		},
		"payload": {}
	}`)}, false, nil
}

func revokeGen() ([][]byte, bool, error) {
	return [][]byte{[]byte(`{
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
	}`)}, false, nil
}

func genReconnectGen(url string, gens ...messageDataGenerator) messageDataGenerator {
	return func() ([][]byte, bool, error) {
		events := [][]byte{[]byte(fmt.Sprintf(`{
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
		}`, url))}

		for _, gen := range gens {
			newEvents, _, _ := gen()
			events = append(events, newEvents...)
		}

		return events, false, nil
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

	client.OnWelcome(func(message twitch.WelcomeMessage) {
		go func() {
			time.Sleep(50 * time.Millisecond)
			client.Close()
		}()
	})

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

// func TestOnReconnect(t *testing.T) {
// 	t.Parallel()

// 	assertEventOccured(t, func(ch chan struct{}) {
// 		client := newClient(t, genReconnectGen(""))
// 		client.OnReconnect(func(message twitch.ReconnectMessage) {
// 			close(ch)
// 		})

// 		go connect(t, client)
// 	})
// }

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
		client := newClient(t, func() ([][]byte, bool, error) {
			return [][]byte{[]byte(`{}`)}, false, nil
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
		client := newClient(t, func() ([][]byte, bool, error) {
			return [][]byte{[]byte(`{`)}, false, nil
		})
		client.OnError(func(err error) {
			close(ch)
		})

		go connect(t, client)
	})
}

func TestReconnectEvent(t *testing.T) {
	t.Parallel()

	reconnectServer, err := newTestServer(keepAliveGen)
	if err != nil {
		t.Fatalf("could not create reconnect server: %v", err)
	}
	reconnectUrl := fmt.Sprintf("http://%s/%s", reconnectServer.Address, "ws")

	client := newClient(t, genReconnectGen(reconnectUrl, revokeGen))

	var keepAliveOccured bool
	client.OnKeepAlive(func(message twitch.KeepAliveMessage) {
		keepAliveOccured = true
		client.Close()
	})

	var revokeOccured bool
	client.OnRevoke(func(message twitch.RevokeMessage) { revokeOccured = true })

	err = client.Connect()
	assert.NoError(t, err)
	assert.Equal(t, reconnectUrl, client.Address, "addresses should match")
	assert.True(t, revokeOccured, "revoke did not fire")
	assert.True(t, keepAliveOccured, "keepalive did not fire")
}
