package twitch

import (
	"encoding/json"
	"time"
)

type MessageMetadata struct {
	MessageID        string    `json:"message_id"`
	MessageType      string    `json:"message_type"`
	MessageTimestamp time.Time `json:"message_timestamp"`
}

type PayloadSession struct {
	ID                      string    `json:"id"`
	Status                  string    `json:"status"`
	ConnectedAt             time.Time `json:"connected_at"`
	KeepaliveTimeoutSeconds int       `json:"keepalive_timeout_seconds"`
	ReconnectUrl            string    `json:"reconnect_url"`
}

type SubscriptionTransport struct {
	Method    string `json:"method"`
	SessionID string `json:"session_id"`
}

type SubscriptionRequest struct {
	Type      EventSubscription     `json:"type"`
	Version   string                `json:"version"`
	Condition map[string]string     `json:"condition"`
	Transport SubscriptionTransport `json:"transport"`
}

type PayloadSubscription struct {
	SubscriptionRequest

	ID       string    `json:"id"`
	Status   string    `json:"status"`
	Cost     int       `json:"cost"`
	CreateAt time.Time `json:"created_at"`
}

type WelcomeMessage struct {
	Metadata MessageMetadata `json:"metadata"`
	Payload  struct {
		Session PayloadSession `json:"session"`
	} `json:"payload"`
}

type KeepAliveMessage struct {
	Metadata MessageMetadata `json:"metadata"`
	Payload  struct{}        `json:"payload"`
}

type NotificationMessage struct {
	Metadata MessageMetadata `json:"metadata"`
	Payload  struct {
		Subscription PayloadSubscription `json:"subscription"`
		Event        *json.RawMessage    `json:"event"`
	} `json:"payload"`
}

type ReconnectMessage struct {
	Metadata MessageMetadata `json:"metadata"`
	Payload  struct {
		Session PayloadSession `json:"session"`
	} `json:"payload"`
}

type RevokeMessage struct {
	Metadata MessageMetadata `json:"metadata"`
	Payload  struct {
		Subscription PayloadSubscription `json:"subscription"`
	}
}
