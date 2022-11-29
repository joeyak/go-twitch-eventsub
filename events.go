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

type eventEmote struct {
	ID    string `json:"id"`
	Begin int    `json:"begin"`
	End   int    `json:"end"`
}

type eventMessage struct {
	Text   string       `json:"text"`
	Emotes []eventEmote `json:"emotes"`
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
