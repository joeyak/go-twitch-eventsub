# Twitch Eventsub

[![Go Report Card](https://goreportcard.com/badge/github.com/joeyak/go-twitch-eventsub)](https://goreportcard.com/report/github.com/joeyak/go-twitch-eventsub)
![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)
[![GoDoc](https://godoc.org/github.com/joeyak/go-twitch-eventsub?status.svg)](https://godoc.org/github.com/joeyak/go-twitch-eventsub/v2)
![tests](https://github.com/joeyak/go-twitch-eventsub/actions/workflows/main.yaml/badge.svg)


Implements a Twitch EventSub Websocket connection

If a websocket connection has no subscriptions, then it will close automatically on twitch's end so call `client.OnWelcome` and subscribe there after getting the subscription ID.

## Major Version Changes

v2 changes `OnRawEvent` from passing `EventSubscription` to `PayloadSubscription`. This allows extra information to be passed in the event instead of just the type.

## Authorization

For authorization, a user access token must be used. An app access token will cause an error. See the Authorization section in the [Twitch Docs](https://dev.twitch.tv/docs/eventsub/manage-subscriptions/#subscribing-to-events)

> When subscribing to events using WebSockets, you must use a user access token only. The request fails if you use an app access token.
>
> ...
>
> When subscribing to events using webhooks, you must use an app access token. The request fails if you use a user access token.

If the error below occurs, it's likely an app access token is being used instead of a user app token.

```
ERROR: could not subscribe to event: 400 Bad Request: {"error":"Bad Request","status":400,"message":"invalid transport and auth combination"}
```

## Example

```go
package main

import (
	"fmt"

	"github.com/joeyak/go-twitch-eventsub/v2"
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
	client.OnRawEvent(func(event string, metadata twitch.MessageMetadata, subscription twitch.PayloadSubscription) {
		fmt.Printf("EVENT[%s]: %s: %s\n", subscription.Type, metadata, event)
	})

	err := client.Connect()
	if err != nil {
		fmt.Printf("Could not connect client: %v\n", err)
	}
}
```

## Events that won't be handled

Events that are in beta will not be handled since it could change, thus possibly breaking code.

The goals event will not be handled because there is no subscription type to request it.
