package twitch

import (
	"math"
	"time"
)

type User struct {
	UserID    string `json:"user_id"`
	UserLogin string `json:"user_login"`
	UserName  string `json:"user_name"`
}

type Broadcaster struct {
	BroadcasterUserId    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
}

type Moderator struct {
	ModeratorUserId    string `json:"moderator_user_id"`
	ModeratorUserLogin string `json:"moderator_user_login"`
	ModeratorUserName  string `json:"moderator_user_name"`
}

type EventChannelUpdate struct {
	Broadcaster

	Title                       string   `json:"title"`
	Language                    string   `json:"language"`
	CategoryID                  string   `json:"category_id"`
	CategoryName                string   `json:"category_name"`
	ContentClassificationLabels []string `json:"content_classification_labels"`
}

type EventChannelFollow struct {
	User
	Broadcaster

	FollowedAt time.Time `json:"followed_at"`
}

type EventChannelSubscribe struct {
	User
	Broadcaster

	Tier   string `json:"tier"`
	IsGift bool   `json:"is_gift"`
}

type EventChannelSubscriptionEnd struct {
	User
	Broadcaster

	Tier   string `json:"tier"`
	IsGift bool   `json:"is_gift"`
}

type EventChannelSubscriptionGift struct {
	User
	Broadcaster

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
	User
	Broadcaster

	Tier             string  `json:"tier"`
	Message          Message `json:"message"`
	CumulativeMonths int     `json:"cumulative_months"`
	StreakMonths     int     `json:"streak_months"`
	DurationMonths   int     `json:"duration_months"`
}

type EventChannelCheer struct {
	User
	Broadcaster

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
	User
	Broadcaster
	Moderator

	Reason      string `json:"reason"`
	BannedAt    string `json:"banned_at"`
	EndsAt      string `json:"ends_at"`
	IsPermanent bool   `json:"is_permanent"`
}

type EventChannelUnban struct {
	User
	Broadcaster
	Moderator
}

type EventChannelModeratorAdd struct {
	Broadcaster
	User
}

type EventChannelModeratorRemove struct {
	Broadcaster
	User
}

type EventChannelVIPAdd struct {
	Broadcaster
	User
}

type EventChannelVIPRemove struct {
	Broadcaster
	User
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
	Broadcaster

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

type CustomChannelPointReward struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Cost   int    `json:"cost"`
	Prompt string `json:"prompt"`
}

type EventChannelChannelPointsCustomRewardRedemptionAdd struct {
	Broadcaster
	User

	ID         string                   `json:"id"`
	UserInput  string                   `json:"user_input"`
	Status     string                   `json:"status"`
	Reward     CustomChannelPointReward `json:"reward"`
	RedeemedAt time.Time                `json:"redeemed_at"`
}

type EventChannelChannelPointsCustomRewardRedemptionUpdate EventChannelChannelPointsCustomRewardRedemptionAdd

type AutomaticChannelPointRewardUnlockedEmote struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type AutomaticChannelPointReward struct {
	Type          string                                    `json:"type"`
	Cost          int                                       `json:"cost"`
	UnlockedEmote *AutomaticChannelPointRewardUnlockedEmote `json:"unlocked_emote"`
}

type EventChannelChannelPointsAutomaticRewardRedemptionAdd struct {
	Broadcaster
	User

	ID         string                      `json:"id"`
	Reward     AutomaticChannelPointReward `json:"reward"`
	Message    Message                     `json:"message"`
	UserInput  string                      `json:"user_input"`
	RedeemedAt time.Time                   `json:"redeemed_at"`
}

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
	Broadcaster

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
	User

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
	Broadcaster

	ID        string              `json:"id"`
	Title     string              `json:"title"`
	Outcomes  []PredictionOutcome `json:"outcomes"`
	StartedAt time.Time           `json:"started_at"`
	LocksAt   time.Time           `json:"locks_at"`
}

type EventChannelPredictionProgress EventChannelPredictionBegin

type EventChannelPredictionLock EventChannelPredictionBegin

type EventChannelPredictionEnd struct {
	Broadcaster

	ID               string              `json:"id"`
	Title            string              `json:"title"`
	WinningOutcomeID string              `json:"winning_outcome_id"`
	Outcomes         []PredictionOutcome `json:"outcomes"`
	Status           string              `json:"status"`
	StartedAt        time.Time           `json:"started_at"`
	EndedAt          time.Time           `json:"ended_at"`
}

type DropEntitlement struct {
	User

	OrganizationId string    `json:"organization_id"`
	CategoryId     string    `json:"category_id"`
	CategoryName   string    `json:"category_name"`
	CampaignId     string    `json:"campaign_id"`
	EntitlementId  string    `json:"entitlement_id"`
	BenefitId      string    `json:"benefit_id"`
	CreatedAt      time.Time `json:"created_at"`
}

type EventDropEntitlementGrant struct {
	ID   string          `json:"id"`
	Data DropEntitlement `json:"data"`
}

type ExtensionProduct struct {
	Name          string `json:"name"`
	Bits          int    `json:"bits"`
	SKU           string `json:"sku"`
	InDevelopment bool   `json:"in_development"`
}

type EventExtensionBitsTransactionCreate struct {
	Broadcaster
	User

	ID                string           `json:"id"`
	ExtensionClientID string           `json:"extension_client_id"`
	Product           ExtensionProduct `json:"product"`
}

type EventChannelGoalBegin struct {
	Broadcaster

	ID                 string    `json:"id"`
	CharityName        string    `json:"charity_name"`
	CharityDescription string    `json:"charity_description"`
	CharityLogo        string    `json:"charity_logo"`
	CharityWebsite     string    `json:"charity_website"`
	CurrentAmount      int       `json:"current_amount"`
	TargetAmount       int       `json:"target_amount"`
	StoppedAt          time.Time `json:"stopped_at"`
}

type EventChannelGoalProgress EventChannelGoalBegin

type EventChannelGoalEnd EventChannelGoalBegin

type HypeTrainContribution struct {
	User

	Type  string `json:"type"`
	Total int    `json:"total"`
}

type EventChannelHypeTrainBegin struct {
	Broadcaster

	Id               string                  `json:"id"`
	Total            int                     `json:"total"`
	Progress         int                     `json:"progress"`
	Goal             int                     `json:"goal"`
	TopContributions []HypeTrainContribution `json:"top_contributions"`
	LastContribution HypeTrainContribution   `json:"last_contribution"`
	Level            int                     `json:"level"`
	StartedAt        time.Time               `json:"started_at"`
	ExpiresAt        time.Time               `json:"expires_at"`
}

type EventChannelHypeTrainProgress struct {
	EventChannelHypeTrainBegin

	Level int `json:"level"`
}

type EventChannelHypeTrainEnd struct {
	Broadcaster

	Id               string                  `json:"id"`
	Level            int                     `json:"level"`
	Total            int                     `json:"total"`
	TopContributions []HypeTrainContribution `json:"top_contributions"`
	StartedAt        time.Time               `json:"started_at"`
	ExpiresAt        time.Time               `json:"expires_at"`
	CooldownEndsAt   time.Time               `json:"cooldown_ends_at"`
}

type EventStreamOnline struct {
	Broadcaster

	Id        string    `json:"id"`
	Type      string    `json:"type"`
	StartedAt time.Time `json:"started_at"`
}

type EventStreamOffline Broadcaster

type EventUserAuthorizationGrant struct {
	User

	ClientID string `json:"client_id"`
}

type EventUserAuthorizationRevoke EventUserAuthorizationGrant

type EventUserUpdate struct {
	User

	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Description   string `json:"description"`
}

type GoalAmount struct {
	Value         int    `json:"value"`
	DecimalPlaces int    `json:"decimal_places"`
	Currency      string `json:"currency"`
}

func (a GoalAmount) Amount() float64 {
	return float64(a.Value) / math.Pow10(a.DecimalPlaces)
}

type BaseCharity struct {
	Broadcaster
	User

	CharityName        string `json:"charity_name"`
	CharityDescription string `json:"charity_description"`
	CharityLogo        string `json:"charity_logo"`
	CharityWebsite     string `json:"charity_website"`
}

type EventChannelCharityCampaignDonate struct {
	BaseCharity

	Amount GoalAmount `json:"amount"`
}

type EventChannelCharityCampaignProgress struct {
	BaseCharity

	CurrentAmount GoalAmount `json:"current_amount"`
	TargetAmount  GoalAmount `json:"target_amount"`
}

type EventChannelCharityCampaignStart struct {
	EventChannelCharityCampaignProgress

	StartedAt time.Time `json:"started_at"`
}

type EventChannelCharityCampaignStop struct {
	EventChannelCharityCampaignProgress

	StoppedAt time.Time `json:"stopped_at"`
}

type EventChannelShieldModeBegin struct {
	Broadcaster
	Moderator

	StartedAt time.Time `json:"started_at"`
	StoppedAt time.Time `json:"stopped_at"`
}

type EventChannelShieldModeEnd EventChannelShieldModeBegin

type EventChannelShoutoutCreate struct {
	Broadcaster
	Moderator

	ToBroadcasterUserId    string    `json:"to_broadcaster_user_id"`
	ToBroadcasterUserLogin string    `json:"to_broadcaster_user_login"`
	ToBroadcasterUserName  string    `json:"to_broadcaster_user_name"`
	StartedAt              time.Time `json:"started_at"`
	ViewerCount            int       `json:"viewer_count"`
	CooldownEndsAt         time.Time `json:"cooldown_ends_at"`
	TargetCooldownEndsAt   time.Time `json:"target_cooldown_ends_at"`
}

type EventChannelShoutoutReceive struct {
	Broadcaster
	Moderator

	FromBroadcasterUserId    string    `json:"from_broadcaster_user_id"`
	FromBroadcasterUserLogin string    `json:"from_broadcaster_user_login"`
	FromBroadcasterUserName  string    `json:"from_broadcaster_user_name"`
	ViewerCount              int       `json:"viewer_count"`
	StartedAt                time.Time `json:"started_at"`
}

type EventChannelAdBreakBegin struct {
	Broadcaster

	DurationSeconds    int       `json:"duration_seconds"`
	StartedAt          time.Time `json:"started_at"`
	IsAutomatic        bool      `json:"is_automatic"`
	RequesterUserId    string    `json:"requester_user_id"`
	RequesterUserLogin string    `json:"requester_user_login"`
	RequesterUserName  string    `json:"requester_user_name"`
}

type EventChannelWarningAcknowledge struct {
	Broadcaster
	User
}

type EventChannelWarningSend struct {
	Broadcaster
	Moderator
	User

	Reason         string   `json:"reason"`
	ChatRulesCited []string `json:"chat_rules_cited"`
}

type EventChannelUnbanRequestCreate struct {
	Broadcaster
	User

	Id        string    `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type EventChannelUnbanRequestResolve struct {
	Broadcaster
	Moderator
	User

	Id             string `json:"id"`
	ResolutionText string `json:"resolution_text"`
	Status         string `json:"status"`
}
