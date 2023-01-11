package twitch_test

import (
	"testing"
	"time"

	"github.com/joeyak/go-twitch-eventsub"
)

func TestEventChannelUpdate(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelUpdate, getTestEventData(twitch.SubChannelUpdate))

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
	client := newClient(t, twitch.SubChannelFollow, getTestEventData(twitch.SubChannelFollow))

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
	client := newClient(t, twitch.SubChannelSubscribe, getTestEventData(twitch.SubChannelSubscribe))

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
	client := newClient(t, twitch.SubChannelSubscriptionEnd, getTestEventData(twitch.SubChannelSubscriptionEnd))

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
	client := newClient(t, twitch.SubChannelSubscriptionGift, getTestEventData(twitch.SubChannelSubscriptionGift))

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

func TestEventChannelSubscriptionGiftAnon(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelSubscriptionGift, getTestEventData(twitch.SubChannelSubscriptionGift, "anon"))

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
	client := newClient(t, twitch.SubChannelSubscriptionMessage, getTestEventData(twitch.SubChannelSubscriptionMessage))

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

func TestEventChannelSubscriptionMessageNoStreak(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelSubscriptionMessage, getTestEventData(twitch.SubChannelSubscriptionMessage, "nostreak"))

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
	client := newClient(t, twitch.SubChannelCheer, getTestEventData(twitch.SubChannelCheer))

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

func TestEventChannelCheerAnon(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubChannelCheer, getTestEventData(twitch.SubChannelCheer, "anon"))

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
	client := newClient(t, twitch.SubChannelRaid, getTestEventData(twitch.SubChannelRaid))

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
	client := newClient(t, twitch.SubChannelBan, getTestEventData(twitch.SubChannelBan))

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
	client := newClient(t, twitch.SubChannelUnban, getTestEventData(twitch.SubChannelUnban))

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
	client := newClient(t, twitch.SubChannelModeratorAdd, getTestEventData(twitch.SubChannelModeratorAdd))

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
	client := newClient(t, twitch.SubChannelModeratorRemove, getTestEventData(twitch.SubChannelModeratorRemove))

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
	client := newClient(t, twitch.SubChannelChannelPointsCustomRewardAdd, getTestEventData(twitch.SubChannelChannelPointsCustomRewardAdd))

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
	client := newClient(t, twitch.SubChannelChannelPointsCustomRewardUpdate, getTestEventData(twitch.SubChannelChannelPointsCustomRewardUpdate))

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
	client := newClient(t, twitch.SubChannelChannelPointsCustomRewardRemove, getTestEventData(twitch.SubChannelChannelPointsCustomRewardRemove))

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
	client := newClient(t, twitch.SubChannelChannelPointsCustomRewardRedemptionAdd, getTestEventData(twitch.SubChannelChannelPointsCustomRewardRedemptionAdd))

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
	client := newClient(t, twitch.SubChannelChannelPointsCustomRewardRedemptionUpdate, getTestEventData(twitch.SubChannelChannelPointsCustomRewardRedemptionUpdate))

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
	client := newClient(t, twitch.SubChannelPollBegin, getTestEventData(twitch.SubChannelPollBegin))

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
	client := newClient(t, twitch.SubChannelPollProgress, getTestEventData(twitch.SubChannelPollProgress))

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
	client := newClient(t, twitch.SubChannelPollEnd, getTestEventData(twitch.SubChannelPollEnd))

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
	client := newClient(t, twitch.SubChannelPredictionBegin, getTestEventData(twitch.SubChannelPredictionBegin))

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
	client := newClient(t, twitch.SubChannelPredictionProgress, getTestEventData(twitch.SubChannelPredictionProgress))

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
	client := newClient(t, twitch.SubChannelPredictionLock, getTestEventData(twitch.SubChannelPredictionLock))

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
	client := newClient(t, twitch.SubChannelPredictionEnd, getTestEventData(twitch.SubChannelPredictionEnd))

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
	client := newClient(t, twitch.SubDropEntitlementGrant, getTestEventData(twitch.SubDropEntitlementGrant))

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
	client := newClient(t, twitch.SubExtensionBitsTransactionCreate, getTestEventData(twitch.SubExtensionBitsTransactionCreate))

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
	client := newClient(t, twitch.SubChannelGoalBegin, getTestEventData(twitch.SubChannelGoalBegin))

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
	client := newClient(t, twitch.SubChannelGoalProgress, getTestEventData(twitch.SubChannelGoalProgress))

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
	client := newClient(t, twitch.SubChannelGoalEnd, getTestEventData(twitch.SubChannelGoalEnd))

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
	client := newClient(t, twitch.SubChannelHypeTrainBegin, getTestEventData(twitch.SubChannelHypeTrainBegin))

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
	client := newClient(t, twitch.SubChannelHypeTrainProgress, getTestEventData(twitch.SubChannelHypeTrainProgress))

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
	client := newClient(t, twitch.SubChannelHypeTrainEnd, getTestEventData(twitch.SubChannelHypeTrainEnd))

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
	client := newClient(t, twitch.SubStreamOnline, getTestEventData(twitch.SubStreamOnline))

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
	client := newClient(t, twitch.SubStreamOffline, getTestEventData(twitch.SubStreamOffline))

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
	client := newClient(t, twitch.SubUserAuthorizationGrant, getTestEventData(twitch.SubUserAuthorizationGrant))

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
	client := newClient(t, twitch.SubUserAuthorizationRevoke, getTestEventData(twitch.SubUserAuthorizationRevoke))

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

func TestEventUserAuthorizationRevokeNoUser(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubUserAuthorizationRevoke, getTestEventData(twitch.SubUserAuthorizationRevoke, "nouser"))

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
	client := newClient(t, twitch.SubUserUpdate, getTestEventData(twitch.SubUserUpdate))

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

func TestEventUserUpdateNoEmail(t *testing.T) {
	t.Parallel()
	client := newClient(t, twitch.SubUserUpdate, getTestEventData(twitch.SubUserUpdate, "noemail"))

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
