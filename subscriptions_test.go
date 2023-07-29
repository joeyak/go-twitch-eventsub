package twitch_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"github.com/joeyak/go-twitch-eventsub/v2"
)

func TestEventVersion(t *testing.T) {
	testCases := []struct {
		Name     string
		Version  string
		Expected string
	}{
		{"Override", "-1", "-1"},
		{"Default", "", "2"},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			assertEventOccured(t, func(ch chan struct{}) {
				listener, err := net.Listen("tcp", "127.0.0.1:0")
				if err != nil {
					t.Error(err)
				}

				mux := http.NewServeMux()
				mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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

					if subscription.Version != tc.Expected {
						t.Error("versions did not match")
					}

					close(ch)
				})

				go http.Serve(listener, mux)

				twitch.SubscribeEventUrl(twitch.SubscribeRequest{
					Event:           twitch.SubChannelUpdate,
					VersionOverride: tc.Version,
				}, fmt.Sprintf("http://%s", listener.Addr().String()))
			})
		})
	}
}
