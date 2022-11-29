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

type eventPollChoices struct {
	ID                string `json:"id"`
	Title             string `json:"title"`
	BitsVotes         int    `json:"bits_votes"`
	ChannelPointVotes int    `json:"channel_points_votes"`
	Votes             int    `json:"votes"`
}

type eventPollVoting struct {
	IsEnabled     bool `json:"is_enabled"`
	AmountPerVote int  `json:"amount_per_vote"`
}

type EventChannelPollBegin struct {
	eventBroadcaster

	ID                  string             `json:"id"`
	Title               string             `json:"title"`
	Choices             []eventPollChoices `json:"choices"`
	BitsVoting          eventPollVoting    `json:"bits_voting"`
	ChannelPointsVoting eventPollVoting    `json:"channel_points_voting"`
	StartedAt           time.Time          `json:"started_at"`
	EndsAt              time.Time          `json:"ends_at"`
}

type EventChannelPollProgress EventChannelPollBegin

type EventChannelPollEnd struct {
	EventChannelPollBegin

	Status string `json:"status"`
}

type eventTopPredictors struct {
	eventUser

	ChannelPointsWon  int `json:"channel_points_won"`
	ChannelPointsUsed int `json:"channel_points_used"`
}

type eventPredictionOutcome struct {
	ID            string               `json:"id"`
	Title         string               `json:"title"`
	Color         string               `json:"color"`
	Users         int                  `json:"users"`
	ChannelPoints int                  `json:"channel_points"`
	TopPredictors []eventTopPredictors `json:"top_predictors"`
}

type EventChannelPredictionBegin struct {
	eventBroadcaster

	ID        string                   `json:"id"`
	Title     string                   `json:"title"`
	Outcomes  []eventPredictionOutcome `json:"outcomes"`
	StartedAt time.Time                `json:"started_at"`
	LocksAt   time.Time                `json:"locks_at"`
}

type EventChannelPredictionProgress EventChannelPredictionBegin

type EventChannelPredictionLock EventChannelPredictionBegin

type EventChannelPredictionEnd struct {
	eventBroadcaster

	ID               string                   `json:"id"`
	Title            string                   `json:"title"`
	WinningOutcomeID int                      `json:"winning_outcome_id"`
	Outcomes         []eventPredictionOutcome `json:"outcomes"`
	Status           string                   `json:"status"`
	StartedAt        time.Time                `json:"started_at"`
	EndedAt          time.Time                `json:"ended_at"`
}
