package twitch_test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/joeyak/go-twitch-eventsub"
)

func newClient(t *testing.T, event twitch.EventSubscription, suffixes ...string) *twitch.Client {
	server, err := NewTestServer(event, suffixes...)
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
	client := newClient(t, twitch.SubChannelUpdate)

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
	client := newClient(t, twitch.SubChannelUpdate)

	ctx, cancel := context.WithCancel(context.Background())
	client.OnRawEvent(func(event string, metadata twitch.MessageMetadata, eventType twitch.EventSubscription) {
		cancel()
	})

	err := client.ConnectWithContext(ctx)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestEventChannelUpdate(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelUpdate)

	ch := make(chan struct{})
	client.OnEventChannelUpdate(func(event twitch.EventChannelUpdate) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelFollow(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelFollow)

	ch := make(chan struct{})
	client.OnEventChannelFollow(func(event twitch.EventChannelFollow) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelSubscribe(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelSubscribe)

	ch := make(chan struct{})
	client.OnEventChannelSubscribe(func(event twitch.EventChannelSubscribe) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelSubscriptionEnd(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelSubscriptionEnd)

	ch := make(chan struct{})
	client.OnEventChannelSubscriptionEnd(func(event twitch.EventChannelSubscriptionEnd) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelSubscriptionGift(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelSubscriptionGift)

	ch := make(chan struct{})
	client.OnEventChannelSubscriptionGift(func(event twitch.EventChannelSubscriptionGift) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelSubscriptionMessage(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelSubscriptionMessage)

	ch := make(chan struct{})
	client.OnEventChannelSubscriptionMessage(func(event twitch.EventChannelSubscriptionMessage) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelCheer(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelCheer)

	ch := make(chan struct{})
	client.OnEventChannelCheer(func(event twitch.EventChannelCheer) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelRaid(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelRaid)

	ch := make(chan struct{})
	client.OnEventChannelRaid(func(event twitch.EventChannelRaid) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelBan(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelBan)

	ch := make(chan struct{})
	client.OnEventChannelBan(func(event twitch.EventChannelBan) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelUnban(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelUnban)

	ch := make(chan struct{})
	client.OnEventChannelUnban(func(event twitch.EventChannelUnban) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelModeratorAdd(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelModeratorAdd)

	ch := make(chan struct{})
	client.OnEventChannelModeratorAdd(func(event twitch.EventChannelModeratorAdd) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelModeratorRemove(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelModeratorRemove)

	ch := make(chan struct{})
	client.OnEventChannelModeratorRemove(func(event twitch.EventChannelModeratorRemove) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelChannelPointsCustomRewardAdd(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelChannelPointsCustomRewardAdd)

	ch := make(chan struct{})
	client.OnEventChannelChannelPointsCustomRewardAdd(func(event twitch.EventChannelChannelPointsCustomRewardAdd) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelChannelPointsCustomRewardUpdate(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelChannelPointsCustomRewardUpdate)

	ch := make(chan struct{})
	client.OnEventChannelChannelPointsCustomRewardUpdate(func(event twitch.EventChannelChannelPointsCustomRewardUpdate) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelChannelPointsCustomRewardRemove(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelChannelPointsCustomRewardRemove)

	ch := make(chan struct{})
	client.OnEventChannelChannelPointsCustomRewardRemove(func(event twitch.EventChannelChannelPointsCustomRewardRemove) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelChannelPointsCustomRewardRedemptionAdd(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelChannelPointsCustomRewardRedemptionAdd)

	ch := make(chan struct{})
	client.OnEventChannelChannelPointsCustomRewardRedemptionAdd(func(event twitch.EventChannelChannelPointsCustomRewardRedemptionAdd) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelChannelPointsCustomRewardRedemptionUpdate(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelChannelPointsCustomRewardRedemptionUpdate)

	ch := make(chan struct{})
	client.OnEventChannelChannelPointsCustomRewardRedemptionUpdate(func(event twitch.EventChannelChannelPointsCustomRewardRedemptionUpdate) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelPollBegin(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelPollBegin)

	ch := make(chan struct{})
	client.OnEventChannelPollBegin(func(event twitch.EventChannelPollBegin) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelPollProgress(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelPollProgress)

	ch := make(chan struct{})
	client.OnEventChannelPollProgress(func(event twitch.EventChannelPollProgress) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelPollEnd(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelPollEnd)

	ch := make(chan struct{})
	client.OnEventChannelPollEnd(func(event twitch.EventChannelPollEnd) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelPredictionBegin(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelPredictionBegin)

	ch := make(chan struct{})
	client.OnEventChannelPredictionBegin(func(event twitch.EventChannelPredictionBegin) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelPredictionProgress(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelPredictionProgress)

	ch := make(chan struct{})
	client.OnEventChannelPredictionProgress(func(event twitch.EventChannelPredictionProgress) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelPredictionLock(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelPredictionLock)

	ch := make(chan struct{})
	client.OnEventChannelPredictionLock(func(event twitch.EventChannelPredictionLock) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelPredictionEnd(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelPredictionEnd)

	ch := make(chan struct{})
	client.OnEventChannelPredictionEnd(func(event twitch.EventChannelPredictionEnd) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventDropEntitlementGrant(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubDropEntitlementGrant)

	ch := make(chan struct{})
	client.OnEventDropEntitlementGrant(func(event []twitch.EventDropEntitlementGrant) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventExtensionBitsTransactionCreate(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubExtensionBitsTransactionCreate)

	ch := make(chan struct{})
	client.OnEventExtensionBitsTransactionCreate(func(event twitch.EventExtensionBitsTransactionCreate) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelGoalBegin(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelGoalBegin)

	ch := make(chan struct{})
	client.OnEventChannelGoalBegin(func(event twitch.EventChannelGoalBegin) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelGoalProgress(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelGoalProgress)

	ch := make(chan struct{})
	client.OnEventChannelGoalProgress(func(event twitch.EventChannelGoalProgress) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelGoalEnd(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelGoalEnd)

	ch := make(chan struct{})
	client.OnEventChannelGoalEnd(func(event twitch.EventChannelGoalEnd) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelHypeTrainBegin(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelHypeTrainBegin)

	ch := make(chan struct{})
	client.OnEventChannelHypeTrainBegin(func(event twitch.EventChannelHypeTrainBegin) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelHypeTrainProgress(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelHypeTrainProgress)

	ch := make(chan struct{})
	client.OnEventChannelHypeTrainProgress(func(event twitch.EventChannelHypeTrainProgress) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventChannelHypeTrainEnd(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelHypeTrainEnd)

	ch := make(chan struct{})
	client.OnEventChannelHypeTrainEnd(func(event twitch.EventChannelHypeTrainEnd) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventStreamOnline(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubStreamOnline)

	ch := make(chan struct{})
	client.OnEventStreamOnline(func(event twitch.EventStreamOnline) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventStreamOffline(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubStreamOffline)

	ch := make(chan struct{})
	client.OnEventStreamOffline(func(event twitch.EventStreamOffline) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventUserAuthorizationGrant(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubUserAuthorizationGrant)

	ch := make(chan struct{})
	client.OnEventUserAuthorizationGrant(func(event twitch.EventUserAuthorizationGrant) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventUserAuthorizationRevoke(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubUserAuthorizationRevoke)

	ch := make(chan struct{})
	client.OnEventUserAuthorizationRevoke(func(event twitch.EventUserAuthorizationRevoke) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}

func TestEventUserUpdate(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubUserUpdate)

	ch := make(chan struct{})
	client.OnEventUserUpdate(func(event twitch.EventUserUpdate) {
		close(ch)
	})

	go connect(t, client)

	select {
	case <-ch:
	case <-time.After(time.Second):
		t.Error("event did not occur")
	}
}
