package twitch

import "time"

type eventUser struct {
	UserID    string `json:"user_id"`
	UserLogin string `json:"user_login"`
	UserName  string `json:"user_name"`
}

type eventBroadcaster struct {
	BroadcasterUserId    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
}

type eventModerator struct {
	ModeratorUserId    string `json:"moderator_user_id"`
	ModeratorUserLogin string `json:"moderator_user_login"`
	ModeratorUserName  string `json:"moderator_user_name"`
}

type EventChannelUpdate struct {
	eventBroadcaster

	Title        string `json:"title"`
	Language     string `json:"language"`
	CategoryID   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	IsMature     bool   `json:"is_mature"`
}

type EventChannelFollow struct {
	eventUser
	eventBroadcaster

	FollowedAt time.Time `json:"followed_at"`
}

type EventChannelSubscribe struct {
	eventUser
	eventBroadcaster

	Tier   string `json:"tier"`
	IsGift bool   `json:"is_gift"`
}

type EventChannelSubscriptionEnd struct {
	eventUser
	eventBroadcaster

	Tier   string `json:"tier"`
	IsGift bool   `json:"is_gift"`
}

type EventChannelSubscriptionGift struct {
	eventUser
	eventBroadcaster

	Total           int    `json:"total"`
	Tier            string `json:"tier"`
	CumulativeTotal int    `json:"cumulative_total"`
	IsAnonymous     bool   `json:"is_anonymous"`
}

type eventEmote struct {
	ID    string `json:"id"`
	Begin int    `json:"begin"`
	End   int    `json:"end"`
}

type eventMessage struct {
	Text   string       `json:"text"`
	Emotes []eventEmote `json:"emotes"`
}

type EventChannelSubscriptionMessage struct {
	eventUser
	eventBroadcaster

	Tier             string       `json:"tier"`
	Message          eventMessage `json:"message"`
	CumulativeMonths int          `json:"cumulative_months"`
	StreakMonths     int          `json:"streak_months"`
	DurationMonths   int          `json:"duration_months"`
}

type EventChannelCheer struct {
	eventUser
	eventBroadcaster

	Message     string `json:"message"`
	Bits        int    `json:"bits"`
	IsAnonymous bool   `json:"is_anonymous"`
}

type EventChannelRaid struct {
	FromBroadcasterUserId    string `json:"from_broadcaster_user_id"`
	FromBroadcasterUserLogin string `json:"from_broadcaster_user_login"`
	FromBroadcasterUserName  string `json:"from_broadcaster_user_name"`
	ToBroadcasterUserId      string `json:"to_broadcaster_user_id"`
	ToBroadcasterUserLogin   string `json:"to_broadcaster_user_login"`
	ToBroadcasterUserName    string `json:"to_broadcaster_user_name"`
	Viewers                  int    `json:"viewers"`
}

type EventChannelBan struct {
	eventUser
	eventBroadcaster
	eventModerator

	Reason      string `json:"reason"`
	BannedAt    string `json:"banned_at"`
	EndsAt      string `json:"ends_at"`
	IsPermanent bool   `json:"is_permanent"`
}

type EventChannelUnban struct {
	eventUser
	eventBroadcaster
	eventModerator
}

type EventChannelModeratorAdd struct {
	eventBroadcaster
	eventUser
}

type EventChannelModeratorRemove struct {
	eventBroadcaster
	eventUser
}

type eventCpMaxPerStream struct {
	IsEnabled bool `json:"is_enabled"`
	Value     int  `json:"value"`
}

type eventImage struct {
	Url1x string `json:"url_1x"`
	Url2x string `json:"url_2x"`
	Url4x string `json:"url_4x"`
}

type eventGlobalCooldown struct {
	IsEnabled bool `json:"is_enabled"`
	Seconds   int  `json:"seconds"`
}

type EventChannelChannelPointsCustomRewardAdd struct {
	eventBroadcaster

	ID                                string              `json:"id"`
	IsEnabled                         bool                `json:"is_enabled"`
	IsPaused                          bool                `json:"is_paused"`
	IsInStock                         bool                `json:"is_in_stock"`
	Title                             string              `json:"title"`
	Cost                              int                 `json:"cost"`
	Prompt                            string              `json:"prompt"`
	IsUserInputRequired               bool                `json:"is_user_input_required"`
	ShouldRedemptionsSkipRequestQueue bool                `json:"should_redemptions_skip_request_queue"`
	MaxPerStream                      eventCpMaxPerStream `json:"max_per_stream"`
	MaxPerUserPerStream               eventCpMaxPerStream `json:"max_per_user_per_stream"`
	BackgroundColor                   string              `json:"background_color"`
	Image                             eventImage          `json:"image"`
	DefaultImage                      eventImage          `json:"default_image"`
	GlobalCooldown                    eventGlobalCooldown `json:"global_cooldown"`
	CooldownExpiresAt                 time.Time           `json:"cooldown_expires_at"`
	RedemptionsRedeemedCurrentStream  int                 `json:"redemptions_redeemed_current_stream"`
}

type EventChannelChannelPointsCustomRewardUpdate EventChannelChannelPointsCustomRewardAdd

type EventChannelChannelPointsCustomRewardRemove EventChannelChannelPointsCustomRewardAdd

type eventChannelPointReward struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Cost   int    `json:"cost"`
	Prompt string `json:"prompt"`
}

type EventChannelChannelPointsCustomRewardRedemptionAdd struct {
	eventBroadcaster
	eventUser

	ID         string                  `json:"id"`
	UserInput  string                  `json:"user_input"`
	Status     string                  `json:"status"`
	Reward     eventChannelPointReward `json:"reward"`
	RedeemedAt time.Time               `json:"redeemed_at"`
}

type EventChannelChannelPointsCustomRewardRedemptionUpdate EventChannelChannelPointsCustomRewardRedemptionAdd
