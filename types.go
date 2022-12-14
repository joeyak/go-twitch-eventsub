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

type payloadSession struct {
	ID                      string    `json:"id"`
	Status                  string    `json:"status"`
	ConnectedAt             time.Time `json:"connected_at"`
	KeepaliveTimeoutSeconds int       `json:"keepalive_timeout_seconds"`
	ReconnectUrl            string    `json:"reconnect_url"`
}

type subscriptionTransport struct {
	Method    string `json:"method"`
	SessionID string `json:"session_id"`
}

type subscriptionRequest struct {
	Type      EventSubscription     `json:"type"`
	Version   string                `json:"version"`
	Condition map[string]string     `json:"condition"`
	Transport subscriptionTransport `json:"transport"`
}

type payloadSubscription struct {
	subscriptionRequest

	ID       string    `json:"id"`
	Status   string    `json:"status"`
	Cost     int       `json:"cost"`
	CreateAt time.Time `json:"created_at"`
}

type messageBase struct {
	Metadata MessageMetadata `json:"metadata"`
}

type WelcomeMessage struct {
	messageBase
	Payload struct {
		Session payloadSession `json:"session"`
	} `json:"payload"`
}

type KeepAliveMessage struct {
	messageBase
	Payload struct{} `json:"payload"`
}

type NotificationMessage struct {
	messageBase
	Payload struct {
		Subscription payloadSubscription `json:"subscription"`
		Event        *json.RawMessage    `json:"event"`
	} `json:"payload"`
}

type ReconnectMessage struct {
	messageBase
	Payload struct {
		Session payloadSession `json:"session"`
	} `json:"payload"`
}

type RevokeMessage struct {
	messageBase
	Payload struct {
		Subscription payloadSubscription `json:"subscription"`
	}
}
