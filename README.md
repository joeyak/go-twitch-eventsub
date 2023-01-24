# Twitch Eventsub

[![Go Report Card](https://goreportcard.com/badge/github.com/joeyak/go-twitch-eventsub)](https://goreportcard.com/report/github.com/joeyak/go-twitch-eventsub)
![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)
[![GoDoc](https://godoc.org/github.com/joeyak/go-twitch-eventsub?status.svg)](https://godoc.org/github.com/joeyak/go-twitch-eventsub)
![tests](https://github.com/joeyak/go-twitch-eventsub/actions/workflows/main.yaml/badge.svg)


Implements a Twitch EventSub Websocket connection

If a websocket connection has no subscriptions, then it will close automatically on twitch's end so call `client.OnWelcome` and subscribe there after getting the subscription ID.

## Example

```go
package main

import (
	"fmt"

	"github.com/joeyak/go-twitch-eventsub"
)

var (
	userID = "<USERID>"
	accessToken = "<ACCESSTOKEN>"
	clientID = "<CLIENTID>"
)

func main() {
	client := twitch.NewClient()

	client.OnError(func(err error) {
		fmt.Printf("ERROR: %v\n", err)
	})
	client.OnWelcome(func(message twitch.WelcomeMessage) {
		fmt.Printf("WELCOME: %v\n", message)

		events := []twitch.EventSubscription{
			twitch.SubStreamOnline,
			twitch.SubStreamOffline,
		}

		for _, event := range events {
			fmt.Printf("subscribing to %s\n", event)
			_, err := twitch.SubscribeEvent(twitch.SubscribeRequest{
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
	client.OnNotification(func(message twitch.NotificationMessage) {
		fmt.Printf("NOTIFICATION: %s: %#v\n", message.Payload.Subscription.Type, message.Payload.Event)
	})
	client.OnKeepAlive(func(message twitch.KeepAliveMessage) {
		fmt.Printf("KEEPALIVE: %v\n", message)
	})
	client.OnRevoke(func(message twitch.RevokeMessage) {
		fmt.Printf("REVOKE: %v\n", message)
	})
	client.OnRawEvent(func(event string, metadata MessageMetadata, eventType EventSubscription) {
		fmt.Printf("EVENT[%s]: %s: %s\n", eventType, metadata, event)
	})

	err := client.Connect()
	if err != nil {
		fmt.Printf("Could not connect client: %v\n", err)
	}
}
```
