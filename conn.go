package eventsub

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
		"session_welcome":       genDefault[WelcomeMessage](),
		"session_keepalive":     genDefault[KeepAliveMessage](),
		"notification":          genDefault[NotificationMessage](),
		"session_reconnect":     genDefault[ReconnectMessage](),
		"authorization_revoked": genDefault[RevokeMessage](),
	}
)

func genDefault[T any]() func() interface{} {
	return func() interface{} {
		return new(T)
	}
}

type Conn struct {
	Address string
	ws      *websocket.Conn
	closed  bool
	ctx     context.Context

	onError        func(err error)
	onWelcome      func(message WelcomeMessage)
	onKeepAlive    func(message KeepAliveMessage)
	onNotification func(message NotificationMessage)
	onReconnect    func(message ReconnectMessage)
	onRevoke       func(message RevokeMessage)
}

func NewConn() *Conn {
	return NewConnWithUrl(twitchWebsocketUrl)
}

func NewConnWithUrl(url string) *Conn {
	return &Conn{
		Address: url,
		closed:  true,
		onError: func(err error) { fmt.Printf("ERROR: %v\n", err) },
	}
}

func (c *Conn) Connect() error {
	return c.ConnectWithContext(context.Background())
}

func (c *Conn) ConnectWithContext(ctx context.Context) error {
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

func (c *Conn) Close() error {
	c.closed = true
	if c.ws == nil {
		return nil
	}
	return c.ws.Close(websocket.StatusNormalClosure, "Stopping Connection")
}

func (c *Conn) OnError(callback func(err error)) {
	c.onError = callback
}

func (c *Conn) OnWelcome(callback func(message WelcomeMessage)) {
	c.onWelcome = callback
}

func (c *Conn) OnKeepAlive(callback func(message KeepAliveMessage)) {
	c.onKeepAlive = callback
}

func (c *Conn) OnNotification(callback func(message NotificationMessage)) {
	c.onNotification = callback
}

func (c *Conn) OnReconnect(callback func(message ReconnectMessage)) {
	c.onReconnect = callback
}

func (c *Conn) OnRevoke(callback func(message RevokeMessage)) {
	c.onRevoke = callback
}

func (c *Conn) handleMessage(data []byte) error {
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
		if c.onKeepAlive != nil {
			c.onKeepAlive(*msg)
		}
	case *NotificationMessage:
		if c.onNotification != nil {
			c.onNotification(*msg)
		}
	case *ReconnectMessage:
		err = c.handleReconnect(*msg)
		if err != nil {
			return fmt.Errorf("could not reconnect: %w", err)
		}
	case *RevokeMessage:
		if c.onRevoke != nil {
			c.onRevoke(*msg)
		}
	default:
		return fmt.Errorf("unhandled %T message: %v", msg, msg)
	}

	return nil
}

func (c *Conn) handleReconnect(message ReconnectMessage) error {
	c.Address = message.Payload.Session.ReconnectUrl
	err := c.dial()
	if err != nil {
		return fmt.Errorf("could not reconnect: %w", err)
	}

	if c.onReconnect != nil {
		c.onReconnect(message)
	}

	return nil
}

func (c *Conn) dial() error {
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
