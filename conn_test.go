package twitch_test

import (
	"context"
	"errors"
	"testing"

	"github.com/joeyak/go-twitch-eventsub"
)

func TestClientNoWelcome(t *testing.T) {
	t.Parallel()

	client := twitch.NewClientWithUrl("")
	err := client.Connect()
	if !errors.Is(err, twitch.ErrNilOnWelcome) {
		t.Fatalf("expected ErrNilOnWelcome, actual %#v", err)
	}
}

func TestClientClose(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubStreamOffline, getTestEventData(twitch.SubStreamOffline))

	client.OnRawEvent(func(event string, metadata twitch.MessageMetadata, eventType twitch.EventSubscription) {
		client.Close()
	})

	err := client.Connect()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestClientCloseWithContext(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubStreamOffline, getTestEventData(twitch.SubStreamOffline))

	ctx, cancel := context.WithCancel(context.Background())
	client.OnRawEvent(func(event string, metadata twitch.MessageMetadata, eventType twitch.EventSubscription) {
		cancel()
	})

	err := client.ConnectWithContext(ctx)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
