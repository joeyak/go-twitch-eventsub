package twitch_test

import (
	"testing"

	"github.com/joeyak/go-twitch-eventsub"
)

func assertSpecificEventOccured(t *testing.T, register func(client *twitch.Client, ch chan struct{}), event twitch.EventSubscription, suffixes ...string) {
	assertEventOccured(t, func(ch chan struct{}) {
		client := newClientWithWelcome(t, event, getTestEventData(event, suffixes...))
		register(client, ch)
		go connect(t, client)
	})
}

func TestNotification(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnNotification(func(message twitch.NotificationMessage) { close(ch) })
	}, twitch.SubStreamOnline)
}

func TestUnkownSubscription(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnError(func(err error) { close(ch) })
	}, "unknown")
}

func TestEventChannelUpdate(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelUpdate(func(event twitch.EventChannelUpdate) { close(ch) })
	}, twitch.SubChannelUpdate)
}

func TestEventChannelFollow(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelFollow(func(event twitch.EventChannelFollow) { close(ch) })
	}, twitch.SubChannelFollow)
}

func TestEventChannelSubscribe(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelSubscribe(func(event twitch.EventChannelSubscribe) { close(ch) })
	}, twitch.SubChannelSubscribe)
}

func TestEventChannelSubscriptionEnd(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelSubscriptionEnd(func(event twitch.EventChannelSubscriptionEnd) { close(ch) })
	}, twitch.SubChannelSubscriptionEnd)
}

func TestEventChannelSubscriptionGift(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelSubscriptionGift(func(event twitch.EventChannelSubscriptionGift) { close(ch) })
	}, twitch.SubChannelSubscriptionGift)
}

func TestEventChannelSubscriptionGiftAnon(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelSubscriptionGift(func(event twitch.EventChannelSubscriptionGift) { close(ch) })
	}, twitch.SubChannelSubscriptionGift, "anon")
}

func TestEventChannelSubscriptionMessage(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelSubscriptionMessage(func(event twitch.EventChannelSubscriptionMessage) { close(ch) })
	}, twitch.SubChannelSubscriptionMessage)
}

func TestEventChannelSubscriptionMessageNoStreak(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelSubscriptionMessage(func(event twitch.EventChannelSubscriptionMessage) { close(ch) })
	}, twitch.SubChannelSubscriptionMessage, "nostreak")
}

func TestEventChannelCheer(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelCheer(func(event twitch.EventChannelCheer) { close(ch) })
	}, twitch.SubChannelCheer)
}

func TestEventChannelCheerAnon(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelCheer(func(event twitch.EventChannelCheer) { close(ch) })
	}, twitch.SubChannelCheer, "anon")
}

func TestEventChannelRaid(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelRaid(func(event twitch.EventChannelRaid) { close(ch) })
	}, twitch.SubChannelRaid)
}

func TestEventChannelBan(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelBan(func(event twitch.EventChannelBan) { close(ch) })
	}, twitch.SubChannelBan)
}

func TestEventChannelUnban(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelUnban(func(event twitch.EventChannelUnban) { close(ch) })
	}, twitch.SubChannelUnban)
}

func TestEventChannelModeratorAdd(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelModeratorAdd(func(event twitch.EventChannelModeratorAdd) { close(ch) })
	}, twitch.SubChannelModeratorAdd)
}

func TestEventChannelModeratorRemove(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelModeratorRemove(func(event twitch.EventChannelModeratorRemove) { close(ch) })
	}, twitch.SubChannelModeratorRemove)
}

func TestEventChannelChannelPointsCustomRewardAdd(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelChannelPointsCustomRewardAdd(func(event twitch.EventChannelChannelPointsCustomRewardAdd) { close(ch) })
	}, twitch.SubChannelChannelPointsCustomRewardAdd)
}

func TestEventChannelChannelPointsCustomRewardUpdate(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelChannelPointsCustomRewardUpdate(func(event twitch.EventChannelChannelPointsCustomRewardUpdate) { close(ch) })
	}, twitch.SubChannelChannelPointsCustomRewardUpdate)
}

func TestEventChannelChannelPointsCustomRewardRemove(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelChannelPointsCustomRewardRemove(func(event twitch.EventChannelChannelPointsCustomRewardRemove) { close(ch) })
	}, twitch.SubChannelChannelPointsCustomRewardRemove)
}

func TestEventChannelChannelPointsCustomRewardRedemptionAdd(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelChannelPointsCustomRewardRedemptionAdd(func(event twitch.EventChannelChannelPointsCustomRewardRedemptionAdd) { close(ch) })
	}, twitch.SubChannelChannelPointsCustomRewardRedemptionAdd)
}

func TestEventChannelChannelPointsCustomRewardRedemptionUpdate(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelChannelPointsCustomRewardRedemptionUpdate(func(event twitch.EventChannelChannelPointsCustomRewardRedemptionUpdate) { close(ch) })
	}, twitch.SubChannelChannelPointsCustomRewardRedemptionUpdate)
}

func TestEventChannelPollBegin(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelPollBegin(func(event twitch.EventChannelPollBegin) { close(ch) })
	}, twitch.SubChannelPollBegin)
}

func TestEventChannelPollProgress(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelPollProgress(func(event twitch.EventChannelPollProgress) { close(ch) })
	}, twitch.SubChannelPollProgress)
}

func TestEventChannelPollEnd(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelPollEnd(func(event twitch.EventChannelPollEnd) { close(ch) })
	}, twitch.SubChannelPollEnd)
}

func TestEventChannelPredictionBegin(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelPredictionBegin(func(event twitch.EventChannelPredictionBegin) { close(ch) })
	}, twitch.SubChannelPredictionBegin)
}

func TestEventChannelPredictionProgress(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelPredictionProgress(func(event twitch.EventChannelPredictionProgress) { close(ch) })
	}, twitch.SubChannelPredictionProgress)
}

func TestEventChannelPredictionLock(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelPredictionLock(func(event twitch.EventChannelPredictionLock) { close(ch) })
	}, twitch.SubChannelPredictionLock)
}

func TestEventChannelPredictionEnd(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelPredictionEnd(func(event twitch.EventChannelPredictionEnd) { close(ch) })
	}, twitch.SubChannelPredictionEnd)
}

func TestEventDropEntitlementGrant(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventDropEntitlementGrant(func(event []twitch.EventDropEntitlementGrant) { close(ch) })
	}, twitch.SubDropEntitlementGrant)
}

func TestEventExtensionBitsTransactionCreate(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventExtensionBitsTransactionCreate(func(event twitch.EventExtensionBitsTransactionCreate) { close(ch) })
	}, twitch.SubExtensionBitsTransactionCreate)
}

func TestEventChannelGoalBegin(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelGoalBegin(func(event twitch.EventChannelGoalBegin) { close(ch) })
	}, twitch.SubChannelGoalBegin)
}

func TestEventChannelGoalProgress(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelGoalProgress(func(event twitch.EventChannelGoalProgress) { close(ch) })
	}, twitch.SubChannelGoalProgress)
}

func TestEventChannelGoalEnd(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelGoalEnd(func(event twitch.EventChannelGoalEnd) { close(ch) })
	}, twitch.SubChannelGoalEnd)
}

func TestEventChannelHypeTrainBegin(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelHypeTrainBegin(func(event twitch.EventChannelHypeTrainBegin) { close(ch) })
	}, twitch.SubChannelHypeTrainBegin)
}

func TestEventChannelHypeTrainProgress(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelHypeTrainProgress(func(event twitch.EventChannelHypeTrainProgress) { close(ch) })
	}, twitch.SubChannelHypeTrainProgress)
}

func TestEventChannelHypeTrainEnd(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventChannelHypeTrainEnd(func(event twitch.EventChannelHypeTrainEnd) { close(ch) })
	}, twitch.SubChannelHypeTrainEnd)
}

func TestEventStreamOnline(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventStreamOnline(func(event twitch.EventStreamOnline) { close(ch) })
	}, twitch.SubStreamOnline)
}

func TestEventStreamOffline(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventStreamOffline(func(event twitch.EventStreamOffline) { close(ch) })
	}, twitch.SubStreamOffline)
}

func TestEventUserAuthorizationGrant(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventUserAuthorizationGrant(func(event twitch.EventUserAuthorizationGrant) { close(ch) })
	}, twitch.SubUserAuthorizationGrant)
}

func TestEventUserAuthorizationRevoke(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventUserAuthorizationRevoke(func(event twitch.EventUserAuthorizationRevoke) { close(ch) })
	}, twitch.SubUserAuthorizationRevoke)
}

func TestEventUserAuthorizationRevokeNoUser(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventUserAuthorizationRevoke(func(event twitch.EventUserAuthorizationRevoke) { close(ch) })
	}, twitch.SubUserAuthorizationRevoke, "nouser")
}

func TestEventUserUpdate(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventUserUpdate(func(event twitch.EventUserUpdate) { close(ch) })
	}, twitch.SubUserUpdate)
}

func TestEventUserUpdateNoEmail(t *testing.T) {
	t.Parallel()

	assertSpecificEventOccured(t, func(client *twitch.Client, ch chan struct{}) {
		client.OnEventUserUpdate(func(event twitch.EventUserUpdate) { close(ch) })
	}, twitch.SubUserUpdate, "noemail")
}
