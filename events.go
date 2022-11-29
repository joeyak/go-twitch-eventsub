package twitch

import (
	"math"
	"time"
)

type user struct {
	UserID    string `json:"user_id"`
	UserLogin string `json:"user_login"`
	UserName  string `json:"user_name"`
}

type broadcaster struct {
	BroadcasterUserId    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
}

type moderator struct {
	ModeratorUserId    string `json:"moderator_user_id"`
	ModeratorUserLogin string `json:"moderator_user_login"`
	ModeratorUserName  string `json:"moderator_user_name"`
}

type EventChannelUpdate struct {
	broadcaster

	Title        string `json:"title"`
	Language     string `json:"language"`
	CategoryID   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	IsMature     bool   `json:"is_mature"`
}

type EventChannelFollow struct {
	user
	broadcaster

	FollowedAt time.Time `json:"followed_at"`
}

type EventChannelSubscribe struct {
	user
	broadcaster

	Tier   string `json:"tier"`
	IsGift bool   `json:"is_gift"`
}

type EventChannelSubscriptionEnd struct {
	user
	broadcaster

	Tier   string `json:"tier"`
	IsGift bool   `json:"is_gift"`
}

type EventChannelSubscriptionGift struct {
	user
	broadcaster

	Total           int    `json:"total"`
	Tier            string `json:"tier"`
	CumulativeTotal int    `json:"cumulative_total"`
	IsAnonymous     bool   `json:"is_anonymous"`
}

type Emote struct {
	ID    string `json:"id"`
	Begin int    `json:"begin"`
	End   int    `json:"end"`
}

type Message struct {
	Text   string  `json:"text"`
	Emotes []Emote `json:"emotes"`
}

type EventChannelSubscriptionMessage struct {
	user
	broadcaster

	Tier             string  `json:"tier"`
	Message          Message `json:"message"`
	CumulativeMonths int     `json:"cumulative_months"`
	StreakMonths     int     `json:"streak_months"`
	DurationMonths   int     `json:"duration_months"`
}

type EventChannelCheer struct {
	user
	broadcaster

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
	user
	broadcaster
	moderator

	Reason      string `json:"reason"`
	BannedAt    string `json:"banned_at"`
	EndsAt      string `json:"ends_at"`
	IsPermanent bool   `json:"is_permanent"`
}

type EventChannelUnban struct {
	user
	broadcaster
	moderator
}

type EventChannelModeratorAdd struct {
	broadcaster
	user
}

type EventChannelModeratorRemove struct {
	broadcaster
	user
}

type MaxChannelPointsPerStream struct {
	IsEnabled bool `json:"is_enabled"`
	Value     int  `json:"value"`
}

type Image struct {
	Url1x string `json:"url_1x"`
	Url2x string `json:"url_2x"`
	Url4x string `json:"url_4x"`
}

type GlobalCooldown struct {
	IsEnabled bool `json:"is_enabled"`
	Seconds   int  `json:"seconds"`
}

type EventChannelChannelPointsCustomRewardAdd struct {
	broadcaster

	ID                                string                    `json:"id"`
	IsEnabled                         bool                      `json:"is_enabled"`
	IsPaused                          bool                      `json:"is_paused"`
	IsInStock                         bool                      `json:"is_in_stock"`
	Title                             string                    `json:"title"`
	Cost                              int                       `json:"cost"`
	Prompt                            string                    `json:"prompt"`
	IsUserInputRequired               bool                      `json:"is_user_input_required"`
	ShouldRedemptionsSkipRequestQueue bool                      `json:"should_redemptions_skip_request_queue"`
	MaxPerStream                      MaxChannelPointsPerStream `json:"max_per_stream"`
	MaxPerUserPerStream               MaxChannelPointsPerStream `json:"max_per_user_per_stream"`
	BackgroundColor                   string                    `json:"background_color"`
	Image                             Image                     `json:"image"`
	DefaultImage                      Image                     `json:"default_image"`
	GlobalCooldown                    GlobalCooldown            `json:"global_cooldown"`
	CooldownExpiresAt                 time.Time                 `json:"cooldown_expires_at"`
	RedemptionsRedeemedCurrentStream  int                       `json:"redemptions_redeemed_current_stream"`
}

type EventChannelChannelPointsCustomRewardUpdate EventChannelChannelPointsCustomRewardAdd

type EventChannelChannelPointsCustomRewardRemove EventChannelChannelPointsCustomRewardAdd

type ChannelPointReward struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Cost   int    `json:"cost"`
	Prompt string `json:"prompt"`
}

type EventChannelChannelPointsCustomRewardRedemptionAdd struct {
	broadcaster
	user

	ID         string             `json:"id"`
	UserInput  string             `json:"user_input"`
	Status     string             `json:"status"`
	Reward     ChannelPointReward `json:"reward"`
	RedeemedAt time.Time          `json:"redeemed_at"`
}

type EventChannelChannelPointsCustomRewardRedemptionUpdate EventChannelChannelPointsCustomRewardRedemptionAdd

type PollChoice struct {
	ID                string `json:"id"`
	Title             string `json:"title"`
	BitsVotes         int    `json:"bits_votes"`
	ChannelPointVotes int    `json:"channel_points_votes"`
	Votes             int    `json:"votes"`
}

type PollVoting struct {
	IsEnabled     bool `json:"is_enabled"`
	AmountPerVote int  `json:"amount_per_vote"`
}

type EventChannelPollBegin struct {
	broadcaster

	ID                  string       `json:"id"`
	Title               string       `json:"title"`
	Choices             []PollChoice `json:"choices"`
	BitsVoting          PollVoting   `json:"bits_voting"`
	ChannelPointsVoting PollVoting   `json:"channel_points_voting"`
	StartedAt           time.Time    `json:"started_at"`
	EndsAt              time.Time    `json:"ends_at"`
}

type EventChannelPollProgress EventChannelPollBegin

type EventChannelPollEnd struct {
	EventChannelPollBegin

	Status string `json:"status"`
}

type TopPredictor struct {
	user

	ChannelPointsWon  int `json:"channel_points_won"`
	ChannelPointsUsed int `json:"channel_points_used"`
}

type PredictionOutcome struct {
	ID            string         `json:"id"`
	Title         string         `json:"title"`
	Color         string         `json:"color"`
	Users         int            `json:"users"`
	ChannelPoints int            `json:"channel_points"`
	TopPredictors []TopPredictor `json:"top_predictors"`
}

type EventChannelPredictionBegin struct {
	broadcaster

	ID        string              `json:"id"`
	Title     string              `json:"title"`
	Outcomes  []PredictionOutcome `json:"outcomes"`
	StartedAt time.Time           `json:"started_at"`
	LocksAt   time.Time           `json:"locks_at"`
}

type EventChannelPredictionProgress EventChannelPredictionBegin

type EventChannelPredictionLock EventChannelPredictionBegin

type EventChannelPredictionEnd struct {
	broadcaster

	ID               string              `json:"id"`
	Title            string              `json:"title"`
	WinningOutcomeID int                 `json:"winning_outcome_id"`
	Outcomes         []PredictionOutcome `json:"outcomes"`
	Status           string              `json:"status"`
	StartedAt        time.Time           `json:"started_at"`
	EndedAt          time.Time           `json:"ended_at"`
}

type DropEntitlement struct {
	user

	OrganizationId string    `json:"organization_id"`
	CategoryId     string    `json:"category_id"`
	CategoryName   string    `json:"category_name"`
	CampaignId     string    `json:"campaign_id"`
	EntitlementId  string    `json:"entitlement_id"`
	BenefitId      string    `json:"benefit_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type EventDropEntitlementGrant struct {
	ID   string            `json:"id"`
	Data []DropEntitlement `json:"data"`
}

type ExtensionProduct struct {
	Name          string `json:"name"`
	Bits          int    `json:"bits"`
	SKU           string `json:"sku"`
	InDevelopment bool   `json:"in_development"`
}

type EventExtensionBitsTransactionCreate struct {
	broadcaster
	user

	ID                string           `json:"id"`
	ExtensionClientID string           `json:"extension_client_id"`
	Product           ExtensionProduct `json:"product"`
}

type GoalAmount struct {
	Value         int    `json:"value"`
	DecimalPlaces int    `json:"decimal_places"`
	Currency      string `json:"currency"`
}

func (a GoalAmount) Amount() float64 {
	return float64(a.Value) / math.Pow10(a.DecimalPlaces)
}

type EventChannelGoalBegin struct {
	broadcaster

	ID                 string     `json:"id"`
	CharityName        string     `json:"charity_name"`
	CharityDescription string     `json:"charity_description"`
	CharityLogo        string     `json:"charity_logo"`
	CharityWebsite     string     `json:"charity_website"`
	CurrentAmount      GoalAmount `json:"current_amount"`
	TargetAmount       GoalAmount `json:"target_amount"`
	StoppedAt          time.Time  `json:"stopped_at"`
}

type EventChannelGoalProgress EventChannelGoalBegin

type EventChannelGoalEnd EventChannelGoalBegin

type HypeTrainContribution struct {
	user

	Type  string `json:"type"`
	Total int    `json:"total"`
}

type EventChannelHypeTrainBegin struct {
	broadcaster

	Id               string                `json:"id"`
	Total            int                   `json:"total"`
	Progress         int                   `json:"progress"`
	Goal             int                   `json:"goal"`
	TopContributions HypeTrainContribution `json:"top_contributions"`
	LastContribution HypeTrainContribution `json:"last_contribution"`
	Level            int                   `json:"level"`
	StartedAt        time.Time             `json:"started_at"`
	ExpiresAt        time.Time             `json:"expires_at"`
}

type EventChannelHypeTrainProgress struct {
	EventChannelHypeTrainBegin

	Level int `json:"level"`
}

type EventChannelHypeTrainEnd struct {
	broadcaster

	Id               string                `json:"id"`
	Level            int                   `json:"level"`
	Total            int                   `json:"total"`
	TopContributions HypeTrainContribution `json:"top_contributions"`
	StartedAt        time.Time             `json:"started_at"`
	ExpiresAt        time.Time             `json:"expires_at"`
	CooldownEndsAt   time.Time             `json:"cooldown_ends_at"`
}

type EventStreamOnline struct {
	broadcaster

	Id        string    `json:"id"`
	Type      string    `json:"type"`
	StartedAt time.Time `json:"started_at"`
}

type EventStreamOffline broadcaster

type EventUserAuthorizationGrant struct {
	user

	ClientID string `json:"client_id"`
}

type EventUserAuthorizationRevoke EventUserAuthorizationGrant

type EventUserUpdate struct {
	user

	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Description   string `json:"description"`
}
