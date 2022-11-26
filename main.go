package eventsub

import (
	"context"
	"fmt"

	"nhooyr.io/websocket"
)

const (
	TwitchWebsocketUrl = "wss://eventsub-beta.wss.twitch.tv/ws"

	MtSessionWelcome = "session_welcome"
)

type Conn struct {
	ws  *websocket.Conn
	url string
}

func NewConn() Conn {
	return NewConnWithUrl(TwitchWebsocketUrl)
}

func NewConnWithUrl(url string) Conn {
	return Conn{
		url: url,
	}
}

func (c *Conn) Dial() error {
	return c.DialWithContext(context.Background())
}

func (c *Conn) DialWithContext(ctx context.Context) error {
	ws, _, err := websocket.Dial(ctx, c.url, nil)
	if err != nil {
		return fmt.Errorf("could not dial")
	}
	c.ws = ws

	return nil
}

func (c *Conn) Close() error {
	return c.ws.Close(websocket.StatusNormalClosure, "Stopping Connection")
}
