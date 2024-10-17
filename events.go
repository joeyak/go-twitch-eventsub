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

type Target struct {
	TargetUserId    string `json:"target_user_id"`
	TargetUserLogin string `json:"target_user_login"`
	TargetUserName  string `json:"target_user_name"`
}

type SourceBroadcaster struct {
	SourceBroadcasterUserId    string `json:"source_broadcaster_user_id"`
	SourceBroadcasterUserLogin string `json:"source_broadcaster_user_login"`
	SourceBroadcasterUserName  string `json:"source_broadcaster_user_name"`
}

type Chatter struct {
	ChatterUserId    string `json:"chatter_user_id"`
	ChatterUserLogin string `json:"chatter_user_login"`
	ChatterUserName  string `json:"chatter_user_name"`
}

type HostBroadcaster struct {
	HostBroadcasterUserId    string `json:"host_broadcaster_user_id"`
	HostBroadcasterUserLogin string `json:"host_broadcaster_user_login"`
	HostBroadcasterUserName  string `json:"host_broadcaster_user_name"`
}

type Ban struct {
	User
	Reason *string `json:"reason,omitempty"`
}

type Timeout struct {
	Ban
	ExpiresAt time.Time `json:"expires_at"`
}

type Raid struct {
	User
	ViewerCount int `json:"viewer_count"`
}

type DeletedMessage struct {
	User
	MessageId   string `json:"message_id"`
	MessageBody string `json:"message_body"`
}

type AutomodTerms struct {
	Action      string   `json:"action"`
	List        string   `json:"list"`
	Terms       []string `json:"terms"`
	FromAutomod bool     `json:"from_automod"`
}

type UnbanRequest struct {
	User
	IsApproved       bool   `json:"is_approved"`
	ModeratorMessage string `json:"moderator_message"`
}

type Followers struct {
	FollowDurationMinutes int `json:"follow_duration_minutes"`
}

type SlowMode struct {
	WaitTimeSeconds int `json:"wait_time_seconds"`
}

type Warning struct {
	User
	Reason         string   `json:"reason"`
	ChatRulesCited []string `json:"chat_rules_cited"`
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

	LowTrustStatus string `json:"low_trust_status"`
}

type EventChannelSharedChatBegin struct {
	Broadcaster
	HostBroadcaster

	SessionId    string        `json:"session_id"`
	Participants []Broadcaster `json:"participants"`
}

type EventChannelSharedChatUpdate EventChannelSharedChatBegin

type EventChannelSharedChatEnd struct {
	Broadcaster
	HostBroadcaster

	SessionId string `json:"session_id"`
}

type UserWhisper struct {
	Text string `json:"text"`
}

type EventUserWhisperMessage struct {
	FromUserId    string      `json:"from_user_id"`
	FromUserLogin string      `json:"from_user_login"`
	FromUserName  string      `json:"from_user_name"`
	ToUserId      string      `json:"to_user_id"`
	ToUserLogin   string      `json:"to_user_login"`
	ToUserName    string      `json:"to_user_name"`
	WhisperId     string      `json:"whisper_id"`
	Whisper       UserWhisper `json:"whisper"`
}

type EventChannelModerate struct {
	Broadcaster
	SourceBroadcaster
	Moderator

	Action              string          `json:"action"`
	Followers           *Followers      `json:"followers,omitempty"`
	Slow                *SlowMode       `json:"slow,omitempty"`
	Vip                 *User           `json:"vip,omitempty"`
	Unvip               *User           `json:"unvip,omitempty"`
	Mod                 *User           `json:"mod,omitempty"`
	Unmod               *User           `json:"unmod,omitempty"`
	Ban                 *Ban            `json:"ban,omitempty"`
	Unban               *User           `json:"unban,omitempty"`
	Timeout             *Timeout        `json:"timeout,omitempty"`
	Untimeout           *User           `json:"untimeout,omitempty"`
	Raid                *Raid           `json:"raid,omitempty"`
	Unraid              *User           `json:"unraid,omitempty"`
	Delete              *DeletedMessage `json:"delete,omitempty"`
	AutomodTerms        *AutomodTerms   `json:"automod_terms,omitempty"`
	UnbanRequest        *UnbanRequest   `json:"unban_request,omitempty"`
	Warn                *Warning        `json:"warn,omitempty"`
	SharedChatBan       *Ban            `json:"shared_chat_ban,omitempty"`
	SharedChatUnban     *User           `json:"shared_chat_unban,omitempty"`
	SharedChatTimeout   *Timeout        `json:"shared_chat_timeout,omitempty"`
	SharedChatuntimeout *User           `json:"shared_chat_untimeout,omitempty"`
	SharedChatDelete    *DeletedMessage `json:"shared_chat_delete,omitempty"`
}

type ChatMessageFragmentCheermote struct {
	Prefix string `json:"prefix"`
	Bits   int    `json:"bits"`
	Tier   int    `json:"tier"`
}

type ChatMessageFragmentEmote struct {
	Id         string   `json:"id"`
	EmoteSetId string   `json:"emote_set_id"`
	OwnerId    string   `json:"owner_id"`
	Format     []string `json:"format"`
}

type ChatMessageFragmentMention User

type ChatMessageFragment struct {
	Type      string                        `json:"type"`
	Text      string                        `json:"text"`
	Cheermote *ChatMessageFragmentCheermote `json:"cheermote,omitempty"`
	Emote     *ChatMessageFragmentEmote     `json:"emote,omitempty"`
	Mention   *ChatMessageFragmentMention   `json:"mention,omitempty"`
}

type ChatMessage struct {
	Text      string                `json:"text"`
	Fragments []ChatMessageFragment `json:"fragments"`
}

type EventAutomodMessageHold struct {
	Broadcaster
	User

	MessageId string      `json:"message_id"`
	Message   ChatMessage `json:"message"`
	Level     int         `json:"level"`
	Category  string      `json:"category"`
	HeldAt    time.Time   `json:"held_at"`
}

type EventAutomodMessageUpdate struct {
	Broadcaster
	User
	Moderator

	MessageId string      `json:"message_id"`
	Message   ChatMessage `json:"message"`
	Level     int         `json:"level"`
	Category  string      `json:"category"`
	Status    string      `json:"status"`
	HeldAt    time.Time   `json:"held_at"`
}

type EventAutomodSettingsUpdate struct {
	Broadcaster
	Moderator

	OverallLevel            *int `json:"overall_level"`
	Disability              int  `json:"disability"`
	Aggression              int  `json:"aggression"`
	SexualitySexOrGender    int  `json:"sexuality_sex_or_gender"`
	Misogyny                int  `json:"misogyny"`
	Bullying                int  `json:"bullying"`
	Swearing                int  `json:"swearing"`
	RaceEthnicityOrReligion int  `json:"race_ethnicity_or_religion"`
	SexBasedTerms           int  `json:"sex_based_terms"`
}

type EventAutomodTermsUpdate struct {
	Broadcaster
	Moderator

	Action      string   `json:"action"`
	FromAutomod bool     `json:"from_automod"`
	Terms       []string `json:"terms"`
}

type EventChannelChatUserMessageHold struct {
	Broadcaster
	User

	MessageId string      `json:"message_id"`
	Message   ChatMessage `json:"message"`
}

type EventChannelChatUserMessageUpdate struct {
	Broadcaster
	User

	Status    string      `json:"status"`
	MessageId string      `json:"message_id"`
	Message   ChatMessage `json:"message"`
}

type EventChannelChatClear Broadcaster

type EventChannelChatClearUserMessages struct {
	Broadcaster
	Target
}

type ChatMessageUserBadge struct {
	SetId string `json:"set_id"`
	Id    string `json:"id"`
	Info  string `json:"info"`
}

type ChatMessageCheer struct {
	Bits int `json:"bits"`
}

type ChatMessageReply struct {
	ParentMessageId   string `json:"parent_message_id"`
	ParentMessageBody string `json:"parent_message_body"`
	ParentUserId      string `json:"parent_user_id"`
	ParentUserName    string `json:"parent_user_name"`
	ParentUserLogin   string `json:"parent_user_login"`
	ThreadMessageId   string `json:"thread_message_id"`
	ThreadUserId      string `json:"thread_user_id"`
	ThreadUserName    string `json:"thread_user_name"`
	ThreadUserLogin   string `json:"thread_user_login"`
}

type EventChannelChatMessage struct {
	Broadcaster
	SourceBroadcaster
	Chatter

	MessageId                   string                  `json:"message_id"`
	SourceMessageId             string                  `json:"source_message_id"`
	Message                     ChatMessage             `json:"message"`
	Color                       string                  `json:"color"`
	Badges                      []ChatMessageUserBadge  `json:"badges"`
	SourceBadges                *[]ChatMessageUserBadge `json:"source_badges,omitempty"`
	MessageType                 string                  `json:"message_type"`
	Cheer                       *ChatMessageCheer       `json:"cheer,omitempty"`
	Reply                       *ChatMessageReply       `json:"reply,omitempty"`
	ChannelPointsCustomRewardId string                  `json:"channel_points_custom_reward_id"`
	ChannelPointsAnimationId    string                  `json:"channel_points_animation_id"`
}

type EventChannelChatMessageDelete struct {
	Broadcaster
	Target

	MessageId string `json:"message_id"`
}

type ChatNotificationSub struct {
	SubTier        string `json:"sub_tier"`
	IsPrime        bool   `json:"is_prime"`
	DurationMonths int    `json:"duration_months"`
}

type ChatNotificationResub struct {
	CumulativeMonths  int    `json:"cumulative_months"`
	DurationMonths    int    `json:"duration_months"`
	StreakMonths      int    `json:"streak_months"`
	SubTier           string `json:"sub_tier"`
	IsPrime           bool   `json:"is_prime"`
	IsGift            bool   `json:"is_gift"`
	GifterIsAnonymous bool   `json:"gifter_is_anonymous"`
	GifterUserId      string `json:"gifter_user_id"`
	GifterUserName    string `json:"gifter_user_name"`
	GifterUserLogin   string `json:"gifter_user_login"`
}

type ChatNotificationSubGift struct {
	DurationMonths     int    `json:"duration_months"`
	CumulativeTotal    int    `json:"cumulative_total"`
	RecipientUserId    string `json:"recipient_user_id"`
	RecipientUserName  string `json:"recipient_user_name"`
	RecipientUserLogin string `json:"recipient_user_login"`
	SubTier            string `json:"sub_tier"`
	CommunityGiftId    string `json:"community_gift_id"`
}

type ChatNotificationCommunitySubGift struct {
	Id              string `json:"id"`
	Total           int    `json:"total"`
	SubTier         string `json:"sub_tier"`
	CumulativeTotal int    `json:"cumulative_total"`
}

type ChatNotificationGiftPaidUpgrade struct {
	GifterIsAnonymous bool   `json:"gifter_is_anonymous"`
	GifterUserId      string `json:"gifter_user_id"`
	GifterUserName    string `json:"gifter_user_name"`
}

type ChatNotificationPrimePaidUpgrade struct {
	SubTier string `json:"sub_tier"`
}

type ChatNotificationPayItForward struct {
	GifterIsAnonymous bool   `json:"gifter_is_anonymous"`
	GifterUserId      string `json:"gifter_user_id"`
	GifterUserName    string `json:"gifter_user_name"`
	GifterUserLogin   string `json:"gifter_user_login"`
}

type ChatNotificationRaid struct {
	User

	ViewerCount     string `json:"viewer_count"`
	ProfileImageUrl string `json:"profile_image_url"`
}

type ChatNotificationUnraid struct{}

type ChatNotificationAnnouncement struct {
	Color string `json:"color"`
}

type ChatNotificationBitsBadgeTier struct {
	Tier int `json:"tier"`
}

type ChatNotificationCharityDonationAmount struct {
	Value        int    `json:"value"`
	DecimalPlace int    `json:"decimal_place"`
	Currency     string `json:"currency"`
}

type ChatNotificationCharityDonation struct {
	CharityName string                                `json:"charity_name"`
	Amount      ChatNotificationCharityDonationAmount `json:"amount"`
}

type EventChannelChatNotification struct {
	Broadcaster
	SourceBroadcaster
	Chatter

	ChatterIsAnonymous bool                    `json:"chatter_is_anonymous"`
	Color              string                  `json:"color"`
	Badges             []ChatMessageUserBadge  `json:"badges"`
	SourceBadges       *[]ChatMessageUserBadge `json:"source_badges"`
	SystemMessage      string                  `json:"system_message"`
	MessageId          string                  `json:"message_id"`
	SourceMessageId    string                  `json:"source_message_id"`
	Message            ChatMessage             `json:"message"`

	NoticeType       string                            `json:"notice_type"`
	Sub              *ChatNotificationSub              `json:"sub,omitempty"`
	Resub            *ChatNotificationResub            `json:"resub,omitempty"`
	SubGift          *ChatNotificationSubGift          `json:"sub_gift,omitempty"`
	CommunitySubGift *ChatNotificationCommunitySubGift `json:"community_sub_gift,omitempty"`
	GiftPaidUpgrade  *ChatNotificationGiftPaidUpgrade  `json:"gift_paid_upgrade,omitempty"`
	PrimePaidUpgrade *ChatNotificationPrimePaidUpgrade `json:"prime_paid_upgrade,omitempty"`
	PayItForward     *ChatNotificationPayItForward     `json:"pay_it_forward,omitempty"`
	Raid             *ChatNotificationRaid             `json:"raid,omitempty"`
	Unraid           *ChatNotificationUnraid           `json:"unraid,omitempty"`
	Announcement     *ChatNotificationAnnouncement     `json:"announcement,omitempty"`
	BitsBadgeTier    *ChatNotificationBitsBadgeTier    `json:"bits_badge_tier,omitempty"`
	CharityDonation  *ChatNotificationCharityDonation  `json:"charity_donation,omitempty"`

	SharedChatSub              *ChatNotificationSub              `json:"shared_chat_sub,omitempty"`
	SharedChatResub            *ChatNotificationResub            `json:"shared_chat_resub,omitempty"`
	SharedChatSubGift          *ChatNotificationSubGift          `json:"shared_chat_sub_gift,omitempty"`
	SharedChatCommunitySubGift *ChatNotificationCommunitySubGift `json:"shared_chat_community_sub_gift,omitempty"`
	SharedChatGiftPaidUpgrade  *ChatNotificationGiftPaidUpgrade  `json:"shared_chat_gift_paid_upgrade,omitempty"`
	SharedChatPrimePaidUpgrade *ChatNotificationPrimePaidUpgrade `json:"shared_chat_prime_paid_upgrade,omitempty"`
	SharedChatPayItForward     *ChatNotificationPayItForward     `json:"shared_chat_pay_it_forward,omitempty"`
	SharedChatRaid             *ChatNotificationRaid             `json:"shared_chat_raid,omitempty"`
	SharedChatAnnouncement     *ChatNotificationAnnouncement     `json:"shared_chat_announcement,omitempty"`
}

type EventChannelChatSettingsUpdate struct {
	Broadcaster

	EmoteMode                   bool `json:"emote_mode"`
	FollowerMode                bool `json:"follower_mode"`
	FollowerModeDurationMinutes int  `json:"follower_mode_duration_minutes"`
	SlowMode                    bool `json:"slow_mode"`
	SlowModeWaitTimeSeconds     int  `json:"slow_mode_wait_time_seconds"`
	SubscriberMode              bool `json:"subscriber_mode"`
	UniqueChatMode              bool `json:"unique_chat_mode"`
}

type SuspiciousUserChatMessage struct {
	ChatMessage

	MessageId string `json:"message_id"`
}

type EventChannelSuspiciousUserMessage struct {
	Broadcaster
	User

	LowTrustStatus       string                    `json:"low_trust_status"`
	SharedBanChannelIds  []string                  `json:"shared_ban_channel_ids"`
	Types                []string                  `json:"types"`
	BanEvasionEvaluation string                    `json:"ban_evasion_evaluation"`
	Message              SuspiciousUserChatMessage `json:"message"`
}

type EventChannelSuspiciousUserUpdate struct {
	Broadcaster
	User
	Moderator

	Id             string `json:"id"`
	ResolutionText string `json:"resolution_text"`
	Status         string `json:"status"`
}
