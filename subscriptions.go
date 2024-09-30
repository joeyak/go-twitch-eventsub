package twitch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const twitchEventSubUrl = "https://api.twitch.tv/helix/eventsub/subscriptions"

type EventSubscription string

var (
	SubChannelUpdate EventSubscription = "channel.update"
	SubChannelFollow EventSubscription = "channel.follow"

	SubChannelSubscribe           EventSubscription = "channel.subscribe"
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
	SubChannelChannelPointsAutomaticRewardRedemptionAdd EventSubscription = "channel.channel_points_automatic_reward_redemption.add"

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

	SubChannelCharityCampaignDonate   EventSubscription = "channel.charity_campaign.donate"
	SubChannelCharityCampaignStart    EventSubscription = "channel.charity_campaign.start"
	SubChannelCharityCampaignProgress EventSubscription = "channel.charity_campaign.progress"
	SubChannelCharityCampaignStop     EventSubscription = "channel.charity_campaign.stop"

	SubChannelShieldModeBegin EventSubscription = "channel.shield_mode.begin"
	SubChannelShieldModeEnd   EventSubscription = "channel.shield_mode.end"

	SubChannelShoutoutCreate  EventSubscription = "channel.shoutout.create"
	SubChannelShoutoutReceive EventSubscription = "channel.shoutout.receive"

	SubChannelAdBreakBegin EventSubscription = "channel.ad_break.begin"

	subMetadata = map[EventSubscription]subscriptionMetadata{
		SubChannelUpdate: {
			Version:  "2",
			EventGen: zeroPtrGen[EventChannelUpdate](),
		},
		SubChannelFollow: {
			Version:  "2",
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
		SubChannelChannelPointsAutomaticRewardRedemptionAdd: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelChannelPointsAutomaticRewardRedemptionAdd](),
		},
		SubChannelPollBegin: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelPollBegin](),
		},
		SubChannelPollProgress: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelPollProgress](),
		},
		SubChannelPollEnd: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelPollEnd](),
		},
		SubChannelPredictionBegin: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelPredictionBegin](),
		},
		SubChannelPredictionProgress: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelPredictionProgress](),
		},
		SubChannelPredictionLock: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelPredictionLock](),
		},
		SubChannelPredictionEnd: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelPredictionEnd](),
		},
		SubDropEntitlementGrant: {
			Version:  "1",
			EventGen: zeroPtrGen[[]EventDropEntitlementGrant](), //func() any { return &[]EventDropEntitlementGrant{} },
		},
		SubExtensionBitsTransactionCreate: {
			Version:  "1",
			EventGen: zeroPtrGen[EventExtensionBitsTransactionCreate](),
		},
		SubChannelGoalBegin: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelGoalBegin](),
		},
		SubChannelGoalProgress: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelGoalProgress](),
		},
		SubChannelGoalEnd: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelGoalEnd](),
		},
		SubChannelHypeTrainBegin: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelHypeTrainBegin](),
		},
		SubChannelHypeTrainProgress: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelHypeTrainProgress](),
		},
		SubChannelHypeTrainEnd: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelHypeTrainEnd](),
		},
		SubStreamOnline: {
			Version:  "1",
			EventGen: zeroPtrGen[EventStreamOnline](),
		},
		SubStreamOffline: {
			Version:  "1",
			EventGen: zeroPtrGen[EventStreamOffline](),
		},
		SubUserAuthorizationGrant: {
			Version:  "1",
			EventGen: zeroPtrGen[EventUserAuthorizationGrant](),
		},
		SubUserAuthorizationRevoke: {
			Version:  "1",
			EventGen: zeroPtrGen[EventUserAuthorizationRevoke](),
		},
		SubUserUpdate: {
			Version:  "1",
			EventGen: zeroPtrGen[EventUserUpdate](),
		},
		SubChannelCharityCampaignDonate: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelCharityCampaignDonate](),
		},
		SubChannelCharityCampaignStart: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelCharityCampaignStart](),
		},
		SubChannelCharityCampaignProgress: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelCharityCampaignProgress](),
		},
		SubChannelCharityCampaignStop: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelCharityCampaignStop](),
		},
		SubChannelShieldModeBegin: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelShieldModeBegin](),
		},
		SubChannelShieldModeEnd: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelShieldModeEnd](),
		},
		SubChannelShoutoutCreate: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelShoutoutCreate](),
		},
		SubChannelShoutoutReceive: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelShoutoutReceive](),
		},
		SubChannelAdBreakBegin: {
			Version:  "1",
			EventGen: zeroPtrGen[EventChannelAdBreakBegin](),
		},
	}
)

type subscriptionMetadata struct {
	Version  string
	EventGen func() interface{}
}

type SubscribeRequest struct {
	SessionID       string
	ClientID        string
	AccessToken     string
	VersionOverride string

	Event     EventSubscription
	Condition map[string]string
}

type SubscribeResponse struct {
	Data         []PayloadSubscription `json:"data"`
	Total        int                   `json:"total"`
	TotalCost    int                   `json:"total_cost"`
	MaxTotalCost int                   `json:"max_total_cost"`
}

func SubscribeEvent(request SubscribeRequest) (SubscribeResponse, error) {
	return SubscribeEventUrlWithContext(context.Background(), request, twitchEventSubUrl)
}

func SubscribeEventUrl(request SubscribeRequest, url string) (SubscribeResponse, error) {
	return SubscribeEventUrlWithContext(context.Background(), request, url)
}

func SubscribeEventWithContext(ctx context.Context, request SubscribeRequest) (SubscribeResponse, error) {
	return SubscribeEventUrlWithContext(ctx, request, twitchEventSubUrl)
}

func SubscribeEventUrlWithContext(ctx context.Context, request SubscribeRequest, url string) (SubscribeResponse, error) {
	version := subMetadata[request.Event].Version
	if request.VersionOverride != "" {
		version = request.VersionOverride
	}

	b, err := json.Marshal(SubscriptionRequest{
		Type:      request.Event,
		Version:   version,
		Condition: request.Condition,
		Transport: SubscriptionTransport{
			Method:    "websocket",
			SessionID: request.SessionID,
		},
	})
	if err != nil {
		return SubscribeResponse{}, fmt.Errorf("could not convert request to json: %w", err)
	}
	buf := bytes.NewBuffer(b)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, buf)
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
