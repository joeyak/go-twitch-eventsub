package twitch

import "time"

type eventUser struct {
	UserID    string `json:"user_id"`
	UserLogin string `json:"user_login"`
	UserName  string `json:"user_name"`
}

type eventBroadcasterUser struct {
	BroadcasterUserId    string `json:"broadcaster_user_id"`
	BroadcasterUserLogin string `json:"broadcaster_user_login"`
	BroadcasterUserName  string `json:"broadcaster_user_name"`
}

type eventModeratorUser struct {
	ModeratorUserId    string `json:"moderator_user_id"`
	ModeratorUserLogin string `json:"moderator_user_login"`
	ModeratorUserName  string `json:"moderator_user_name"`
}

type EventChannelUpdate struct {
	eventBroadcasterUser

	Title        string `json:"title"`
	Language     string `json:"language"`
	CategoryID   string `json:"category_id"`
	CategoryName string `json:"category_name"`
	IsMature     bool   `json:"is_mature"`
}

type EventChannelFollow struct {
	eventUser
	eventBroadcasterUser

	FollowedAt time.Time `json:"followed_at"`
}

type EventChannelSubscribe struct {
	eventUser
	eventBroadcasterUser

	Tier   string `json:"tier"`
	IsGift bool   `json:"is_gift"`
}

type EventChannelSubscriptionEnd struct {
	eventUser
	eventBroadcasterUser

	Tier   string `json:"tier"`
	IsGift bool   `json:"is_gift"`
}

type EventChannelSubscriptionGift struct {
	eventUser
	eventBroadcasterUser

	Total           int    `json:"total"`
	Tier            string `json:"tier"`
	CumulativeTotal int    `json:"cumulative_total"`
	IsAnonymous     bool   `json:"is_anonymous"`
}

type EventChannelSubscriptionMessage struct {
	eventUser
	eventBroadcasterUser

	Tier             string   `json:"tier"`
	Message          struct{} `json:"message"`
	CumulativeMonths int      `json:"cumulative_months"`
	StreakMonths     int      `json:"streak_months"`
	DurationMonths   int      `json:"duration_months"`
}

type EventChannelBan struct {
	eventUser
	eventBroadcasterUser
	eventModeratorUser

	Reason      string `json:"reason"`
	BannedAt    string `json:"banned_at"`
	EndsAt      string `json:"ends_at"`
	IsPermanent bool   `json:"is_permanent"`
}

type EventChannelUnban struct {
	eventUser
	eventBroadcasterUser
	eventModeratorUser
}
