# Twitch Eventsub

Implements a Twitch EventSub Websocket connection

If a websocket connection has no subscriptions, then it will close automatically on twitch's end

## Example

```go
package main

import (
	"fmt"

	eventsub "github.com/joeyak/go-twitch-eventsub"
)

var (
	userID = "<USERID>"
	accessToken = "<ACCESSTOKEN>"
	clientID = "<CLIENTID>"
)

func main() {
	conn := eventsub.NewConn()

	conn.OnError(func(err error) {
		fmt.Printf("ERROR: %v\n", err)
	})
	conn.OnWelcome(func(message eventsub.WelcomeMessage) {
		fmt.Printf("WELCOME: %v\n", message)

		events := []eventsub.EventSubscription{
			eventsub.SubStreamOnline,
			eventsub.SubStreamOffline,
		}

		for _, event := range events {
			fmt.Printf("subscribing to %s\n", event)
			_, err := eventsub.SubscribeEvent(eventsub.SubscribeRequest{
				SessionID:   message.Payload.Session.ID,
				ClientID:    clientID,
				AccessToken: accessToken,
				Event:       event,
				Condition: map[string]string{
					"broadcaster_user_id": userID,
				},
			})
			if err != nil {
				fmt.Printf("ERROR: %v\n", err)
				return
			}
		}
	})
	conn.OnNotification(func(message eventsub.NotificationMessage) {
		fmt.Printf("NOTIFICATION: %s: %#v\n", message.Payload.Subscription.Type, message.Payload.Event)
	})
	conn.OnKeepAlive(func(message eventsub.KeepAliveMessage) {
		fmt.Printf("KEEPALIVE: %v\n", message)
	})
	conn.OnRevoke(func(message eventsub.RevokeMessage) {
		fmt.Printf("REVOKE: %v\n", message)
	})

	err := conn.Connect()
	if err != nil {
		fmt.Printf("Could not connect eventsub: %v\n", err)
	}
}
```