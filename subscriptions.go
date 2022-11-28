package eventsub

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

	versionMap = map[EventSubscription]string{
		SubChannelUpdate:                                    "1",
		SubChannelFollow:                                    "1",
		SubChannelSubscribe:                                 "1",
		SubChannelSubscriptionEnd:                           "1",
		SubChannelSubscriptionGift:                          "1",
		SubChannelSubscriptionMessage:                       "1",
		SubChannelCheer:                                     "1",
		SubChannelRaid:                                      "1",
		SubChannelBan:                                       "1",
		SubChannelUnban:                                     "1",
		SubChannelModeratorAdd:                              "1",
		SubChannelModeratorRemove:                           "1",
		SubChannelChannelPointsCustomRewardAdd:              "1",
		SubChannelChannelPointsCustomRewardUpdate:           "1",
		SubChannelChannelPointsCustomRewardRemove:           "1",
		SubChannelChannelPointsCustomRewardRedemptionAdd:    "1",
		SubChannelChannelPointsCustomRewardRedemptionUpdate: "1",
		SubChannelPollBegin:                                 "1",
		SubChannelPollProgress:                              "1",
		SubChannelPollEnd:                                   "1",
		SubChannelPredictionBegin:                           "1",
		SubChannelPredictionProgress:                        "1",
		SubChannelPredictionLock:                            "1",
		SubChannelPredictionEnd:                             "1",
		SubDropEntitlementGrant:                             "1",
		SubExtensionBitsTransactionCreate:                   "1",
		SubChannelGoalBegin:                                 "1",
		SubChannelGoalProgress:                              "1",
		SubChannelGoalEnd:                                   "1",
		SubChannelHypeTrainBegin:                            "1",
		SubChannelHypeTrainProgress:                         "1",
		SubChannelHypeTrainEnd:                              "1",
		SubStreamOnline:                                     "1",
		SubStreamOffline:                                    "1",
		SubUserAuthorizationGrant:                           "1",
		SubUserAuthorizationRevoke:                          "1",
		SubUserUpdate:                                       "1",
	}
)

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
	version := versionMap[request.Event]

	b, err := json.Marshal(subscriptionRequest{
		Type:      request.Event,
		Version:   version,
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
