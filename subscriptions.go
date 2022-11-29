package twitch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const twitchEventSubUrl = "https://api.twitch.tv/helix/eventsub/subscriptions"

type EventSubscription string

var (
	SubChannelUpdate    EventSubscription = "channel.update"
	SubChannelFollow    EventSubscription = "channel.follow"
	SubChannelSubscribe EventSubscription = "channel.subscribe"

	SubChannelSubscriptionEnd     EventSubscription = "channel.subscription.end"
	SubChannelSubscriptionGift    EventSubscription = "channel.subscription.gift"
	SubChannelSubscriptionMessage EventSubscription = "channel.subscription.message"

	SubChannelCheer EventSubscription = "channel.cheer"
	SubChannelRaid  EventSubscription = "channel.raid"
	SubChannelBan   EventSubscription = "channel.ban"
	SubChannelUnban EventSubscription = "channel.unban"

	SubChannelModeratorAdd    EventSubscription = "channel.moderator.add"
	SubChannelModeratorRemove EventSubscription = "channel.moderator.remove"

	SubChannelChannelPointsCustomRewardAdd              EventSubscription = "channel.channel_points_custom_reward.add"
	SubChannelChannelPointsCustomRewardUpdate           EventSubscription = "channel.channel_points_custom_reward.update"
	SubChannelChannelPointsCustomRewardRemove           EventSubscription = "channel.channel_points_custom_reward.remove"
	SubChannelChannelPointsCustomRewardRedemptionAdd    EventSubscription = "channel.channel_points_custom_reward_redemption.add"
	SubChannelChannelPointsCustomRewardRedemptionUpdate EventSubscription = "channel.channel_points_custom_reward_redemption.update"

	SubChannelPollBegin    EventSubscription = "channel.poll.begin"
	SubChannelPollProgress EventSubscription = "channel.poll.progress"
	SubChannelPollEnd      EventSubscription = "channel.poll.end"

	SubChannelPredictionBegin    EventSubscription = "channel.prediction.begin"
	SubChannelPredictionProgress EventSubscription = "channel.prediction.progress"
	SubChannelPredictionLock     EventSubscription = "channel.prediction.lock"
	SubChannelPredictionEnd      EventSubscription = "channel.prediction.end"

	SubDropEntitlementGrant           EventSubscription = "drop.entitlement.grant"
	SubExtensionBitsTransactionCreate EventSubscription = "extension.bits_transaction.create"

	SubChannelGoalBegin    EventSubscription = "channel.goal.begin"
	SubChannelGoalProgress EventSubscription = "channel.goal.progress"
	SubChannelGoalEnd      EventSubscription = "channel.goal.end"

	SubChannelHypeTrainBegin    EventSubscription = "channel.hype_train.begin"
	SubChannelHypeTrainProgress EventSubscription = "channel.hype_train.progress"
	SubChannelHypeTrainEnd      EventSubscription = "channel.hype_train.end"

	SubStreamOnline  EventSubscription = "stream.online"
	SubStreamOffline EventSubscription = "stream.offline"

	SubUserAuthorizationGrant  EventSubscription = "user.authorization.grant"
	SubUserAuthorizationRevoke EventSubscription = "user.authorization.revoke"
	SubUserUpdate              EventSubscription = "user.update"

	subMetadata = map[EventSubscription]subscriptionMetadata{
		SubChannelUpdate: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelUpdate](),
		},
		SubChannelFollow: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelFollow](),
		},
		SubChannelSubscribe: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelSubscribe](),
		},
		SubChannelSubscriptionEnd: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelSubscriptionEnd](),
		},
		SubChannelSubscriptionGift: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelSubscriptionGift](),
		},
		SubChannelSubscriptionMessage: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelSubscriptionMessage](),
		},
		SubChannelCheer: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelCheer](),
		},
		SubChannelRaid: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelRaid](),
		},
		SubChannelBan: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelBan](),
		},
		SubChannelUnban: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelUnban](),
		},
		SubChannelModeratorAdd: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelModeratorAdd](),
		},
		SubChannelModeratorRemove: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelModeratorRemove](),
		},
		SubChannelChannelPointsCustomRewardAdd: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelChannelPointsCustomRewardAdd](),
		},
		SubChannelChannelPointsCustomRewardUpdate: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelChannelPointsCustomRewardUpdate](),
		},
		SubChannelChannelPointsCustomRewardRemove: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelChannelPointsCustomRewardRemove](),
		},
		SubChannelChannelPointsCustomRewardRedemptionAdd: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelChannelPointsCustomRewardRedemptionAdd](),
		},
		SubChannelChannelPointsCustomRewardRedemptionUpdate: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelChannelPointsCustomRewardRedemptionUpdate](),
		},
		SubChannelPollBegin: {
			Version: "1",
		},
		SubChannelPollProgress: {
			Version: "1",
		},
		SubChannelPollEnd: {
			Version: "1",
		},
		SubChannelPredictionBegin: {
			Version: "1",
		},
		SubChannelPredictionProgress: {
			Version: "1",
		},
		SubChannelPredictionLock: {
			Version: "1",
		},
		SubChannelPredictionEnd: {
			Version: "1",
		},
		SubDropEntitlementGrant: {
			Version: "1",
		},
		SubExtensionBitsTransactionCreate: {
			Version: "1",
		},
		SubChannelGoalBegin: {
			Version: "1",
		},
		SubChannelGoalProgress: {
			Version: "1",
		},
		SubChannelGoalEnd: {
			Version: "1",
		},
		SubChannelHypeTrainBegin: {
			Version: "1",
		},
		SubChannelHypeTrainProgress: {
			Version: "1",
		},
		SubChannelHypeTrainEnd: {
			Version: "1",
		},
		SubStreamOnline: {
			Version: "1",
		},
		SubStreamOffline: {
			Version: "1",
		},
		SubUserAuthorizationGrant: {
			Version: "1",
		},
		SubUserAuthorizationRevoke: {
			Version: "1",
		},
		SubUserUpdate: {
			Version: "1",
		},
	}
)

type subscriptionMetadata struct {
	Version  string
	EventGen func() interface{}
}

type SubscribeRequest struct {
	SessionID   string
	ClientID    string
	AccessToken string
	Event       EventSubscription
	Condition   map[string]string
}

type SubscribeResponse struct {
	Data         []payloadSubscription `json:"data"`
	Total        int                   `json:"total"`
	TotalCost    int                   `json:"total_cost"`
	MaxTotalCost int                   `json:"max_total_cost"`
}

func SubscribeEvent(request SubscribeRequest) (SubscribeResponse, error) {
	metadata := subMetadata[request.Event]

	b, err := json.Marshal(subscriptionRequest{
		Type:      request.Event,
		Version:   metadata.Version,
		Condition: request.Condition,
		Transport: subscriptionTransport{
			Method:    "websocket",
			SessionID: request.SessionID,
		},
	})
	if err != nil {
		return SubscribeResponse{}, fmt.Errorf("could not convert request to json: %w", err)
	}
	buf := bytes.NewBuffer(b)

	req, err := http.NewRequest(http.MethodPost, twitchEventSubUrl, buf)
	if err != nil {
		return SubscribeResponse{}, fmt.Errorf("could not create new request: %w", err)
	}

	req.Header.Set("Client-Id", request.ClientID)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", request.AccessToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return SubscribeResponse{}, fmt.Errorf("could not subscribe to event: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != 202 {
		return SubscribeResponse{}, fmt.Errorf("could not subscribe to event: %s: %s", resp.Status, string(body))
	}

	var subscription SubscribeResponse
	err = json.Unmarshal(body, &subscription)
	if err != nil {
		return SubscribeResponse{}, fmt.Errorf("could not unmarshal subscription response: %w", err)
	}

	return subscription, nil
}
