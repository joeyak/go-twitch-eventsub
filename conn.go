package twitch

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"nhooyr.io/websocket"
)

const (
	twitchWebsocketUrl = "wss://eventsub-beta.wss.twitch.tv/ws"
)

var (
	ErrConnClosed   = fmt.Errorf("connection closed")
	ErrNilOnWelcome = fmt.Errorf("OnWelcome function was not set")

	messageTypeMap = map[string]func() interface{}{
		"session_welcome":       zeroPtrGen[WelcomeMessage](),
		"session_keepalive":     zeroPtrGen[KeepAliveMessage](),
		"notification":          zeroPtrGen[NotificationMessage](),
		"session_reconnect":     zeroPtrGen[ReconnectMessage](),
		"authorization_revoked": zeroPtrGen[RevokeMessage](),
	}
)

func zeroPtrGen[T any]() func() interface{} {
	return func() interface{} {
		return new(T)
	}
}

func callFunc[T any](f func(T), v T) {
	if f != nil {
		f(v)
	}
}

type Client struct {
	Address string
	ws      *websocket.Conn
	closed  bool
	ctx     context.Context

	// Responses
	onError        func(err error)
	onWelcome      func(message WelcomeMessage)
	onKeepAlive    func(message KeepAliveMessage)
	onNotification func(message NotificationMessage)
	onReconnect    func(message ReconnectMessage)
	onRevoke       func(message RevokeMessage)

	// Events
	onEventChannelUpdate              func(event EventChannelUpdate)
	onEventChannelFollow              func(event EventChannelFollow)
	onEventChannelSubscribe           func(event EventChannelSubscribe)
	onEventChannelSubscriptionEnd     func(event EventChannelSubscriptionEnd)
	onEventChannelSubscriptionGift    func(event EventChannelSubscriptionGift)
	onEventChannelSubscriptionMessage func(event EventChannelSubscriptionMessage)
	onEventChannelBan                 func(event EventChannelBan)
	onEventChannelUnban               func(event EventChannelUnban)
}

func NewClient() *Client {
	return NewClientWithUrl(twitchWebsocketUrl)
}

func NewClientWithUrl(url string) *Client {
	return &Client{
		Address: url,
		closed:  true,
		onError: func(err error) { fmt.Printf("ERROR: %v\n", err) },
	}
}

func (c *Client) Connect() error {
	return c.ConnectWithContext(context.Background())
}

func (c *Client) ConnectWithContext(ctx context.Context) error {
	if c.onWelcome == nil {
		return ErrNilOnWelcome
	}

	c.ctx = ctx
	err := c.dial()
	if err != nil {
		return err
	}
	defer func() { c.ws = nil }()

	for {
		_, data, err := c.ws.Read(ctx)
		if err != nil {
			var closeError websocket.CloseError
			if c.closed && errors.As(err, &closeError) {
				return nil
			}
			return fmt.Errorf("could not read message: %w", err)
		}

		err = c.handleMessage(data)
		if err != nil {
			c.onError(err)
		}
	}
}

func (c *Client) Close() error {
	c.closed = true
	if c.ws == nil {
		return nil
	}
	return c.ws.Close(websocket.StatusNormalClosure, "Stopping Connection")
}

func (c *Client) IsClosed() bool {
	return c.closed
}

func (c *Client) OnError(callback func(err error)) {
	c.onError = callback
}

func (c *Client) OnWelcome(callback func(message WelcomeMessage)) {
	c.onWelcome = callback
}

func (c *Client) OnKeepAlive(callback func(message KeepAliveMessage)) {
	c.onKeepAlive = callback
}

func (c *Client) OnNotification(callback func(message NotificationMessage)) {
	c.onNotification = callback
}

func (c *Client) OnReconnect(callback func(message ReconnectMessage)) {
	c.onReconnect = callback
}

func (c *Client) OnRevoke(callback func(message RevokeMessage)) {
	c.onRevoke = callback
}

func (c *Client) OnEventChannelUpdate(callback func(event EventChannelUpdate)) {
	c.onEventChannelUpdate = callback
}

func (c *Client) OnEventChannelFollow(callback func(event EventChannelFollow)) {
	c.onEventChannelFollow = callback
}

func (c *Client) OnEventChannelSubscribe(callback func(event EventChannelSubscribe)) {
	c.onEventChannelSubscribe = callback
}

func (c *Client) OnEventChannelSubscriptionEnd(callback func(event EventChannelSubscriptionEnd)) {
	c.onEventChannelSubscriptionEnd = callback
}

func (c *Client) OnEventChannelSubscriptionGift(callback func(event EventChannelSubscriptionGift)) {
	c.onEventChannelSubscriptionGift = callback
}

func (c *Client) OnEventChannelSubscriptionMessage(callback func(event EventChannelSubscriptionMessage)) {
	c.onEventChannelSubscriptionMessage = callback
}

func (c *Client) OnEventChannelBan(callback func(event EventChannelBan)) {
	c.onEventChannelBan = callback
}

func (c *Client) OnEventChannelUnban(callback func(event EventChannelUnban)) {
	c.onEventChannelUnban = callback
}

func (c *Client) handleMessage(data []byte) error {
	var baseMessage messageBase
	err := json.Unmarshal(data, &baseMessage)
	if err != nil {
		return fmt.Errorf("could not unmarshal basemessage to get message type: %w", err)
	}

	messageType := baseMessage.Metadata.MessageType
	genMessage, ok := messageTypeMap[messageType]
	if !ok {
		return fmt.Errorf("unkown message type %s: %s", messageType, string(data))
	}

	message := genMessage()
	err = json.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("could not unmarshal message into %s: %w", messageType, err)
	}

	switch msg := message.(type) {
	case *WelcomeMessage:
		c.onWelcome(*msg)
	case *KeepAliveMessage:
		callFunc(c.onKeepAlive, *msg)
	case *NotificationMessage:
		callFunc(c.onNotification, *msg)
		err = c.handleNotification(*msg)
		if err != nil {
			return fmt.Errorf("could not handle notification: %w", err)
		}
	case *ReconnectMessage:
		err = c.handleReconnect(*msg)
		if err != nil {
			return fmt.Errorf("could not reconnect: %w", err)
		}
	case *RevokeMessage:
		callFunc(c.onRevoke, *msg)
	default:
		return fmt.Errorf("unhandled %T message: %v", msg, msg)
	}

	return nil
}

func (c *Client) handleReconnect(message ReconnectMessage) error {
	c.Address = message.Payload.Session.ReconnectUrl
	err := c.dial()
	if err != nil {
		return fmt.Errorf("could not reconnect: %w", err)
	}

	callFunc(c.onReconnect, message)

	return nil
}

func (c *Client) handleNotification(message NotificationMessage) error {
	data, err := message.Payload.Event.MarshalJSON()
	if err != nil {
		return fmt.Errorf("could not get event json: %w", err)
	}

	subType := message.Payload.Subscription.Type
	metadata, ok := subMetadata[subType]
	if !ok {
		return fmt.Errorf("unkown subscription type %s", subType)
	}

	var newEvent interface{}
	if metadata.EventGen != nil {
		newEvent = metadata.EventGen()
		err = json.Unmarshal(data, newEvent)
		if err != nil {
			return fmt.Errorf("could not unmarshal %s json: %w", subType, err)
		}
	}

	switch event := newEvent.(type) {
	case *EventChannelUpdate:
		callFunc(c.onEventChannelUpdate, *event)
	case *EventChannelFollow:
		callFunc(c.onEventChannelFollow, *event)
	case *EventChannelSubscribe:
		callFunc(c.onEventChannelSubscribe, *event)
	case *EventChannelSubscriptionEnd:
		callFunc(c.onEventChannelSubscriptionEnd, *event)
	case *EventChannelSubscriptionGift:
		callFunc(c.onEventChannelSubscriptionGift, *event)
	case *EventChannelSubscriptionMessage:
		callFunc(c.onEventChannelSubscriptionMessage, *event)
	case *EventChannelBan:
		callFunc(c.onEventChannelBan, *event)
	case *EventChannelUnban:
		callFunc(c.onEventChannelUnban, *event)
	default:
		c.onError(fmt.Errorf("unkown event type %s", subType))
	}

	return nil
}

func (c *Client) dial() error {
	ws, _, err := websocket.Dial(c.ctx, c.Address, nil)
	if err != nil {
		return fmt.Errorf("could not dial twitch: %w", err)
	}

	if c.ws != nil && !c.closed {
		err := c.Close()
		if err != nil {
			return fmt.Errorf("could not close existing connection: %w", err)
		}
	}
	c.ws = ws
	c.closed = false

	return nil
}
