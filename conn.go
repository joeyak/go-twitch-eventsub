package twitch

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"nhooyr.io/websocket"
)

const (
	twitchWebsocketUrl = "wss://eventsub.wss.twitch.tv/ws"
)

var (
	ErrConnClosed   = fmt.Errorf("connection closed")
	ErrNilOnWelcome = fmt.Errorf("OnWelcome function was not set")

	messageTypeMap = map[string]func() any{
		"session_welcome":   zeroPtrGen[WelcomeMessage](),
		"session_keepalive": zeroPtrGen[KeepAliveMessage](),
		"notification":      zeroPtrGen[NotificationMessage](),
		"session_reconnect": zeroPtrGen[ReconnectMessage](),
		"revocation":        zeroPtrGen[RevokeMessage](),
	}
)

func zeroPtrGen[T any]() func() any {
	return func() any {
		return new(T)
	}
}

func callFunc[T any](f func(T), v T) {
	if f != nil {
		go f(v)
	}
}

type Client struct {
	Address   string
	ws        *websocket.Conn
	connected bool
	ctx       context.Context

	reconnecting bool
	reconnected  chan struct{}

	// Responses
	onError        func(err error)
	onWelcome      func(message WelcomeMessage)
	onKeepAlive    func(message KeepAliveMessage)
	onNotification func(message NotificationMessage)
	onReconnect    func(message ReconnectMessage)
	onRevoke       func(message RevokeMessage)

	// Events
	onRawEvent                                              func(event string, metadata MessageMetadata, subscription PayloadSubscription)
	onEventChannelUpdate                                    func(event EventChannelUpdate)
	onEventChannelFollow                                    func(event EventChannelFollow)
	onEventChannelSubscribe                                 func(event EventChannelSubscribe)
	onEventChannelSubscriptionEnd                           func(event EventChannelSubscriptionEnd)
	onEventChannelSubscriptionGift                          func(event EventChannelSubscriptionGift)
	onEventChannelSubscriptionMessage                       func(event EventChannelSubscriptionMessage)
	onEventChannelCheer                                     func(event EventChannelCheer)
	onEventChannelRaid                                      func(event EventChannelRaid)
	onEventChannelBan                                       func(event EventChannelBan)
	onEventChannelUnban                                     func(event EventChannelUnban)
	onEventChannelModeratorAdd                              func(event EventChannelModeratorAdd)
	onEventChannelModeratorRemove                           func(event EventChannelModeratorRemove)
	onEventChannelChannelPointsCustomRewardAdd              func(event EventChannelChannelPointsCustomRewardAdd)
	onEventChannelChannelPointsCustomRewardUpdate           func(event EventChannelChannelPointsCustomRewardUpdate)
	onEventChannelChannelPointsCustomRewardRemove           func(event EventChannelChannelPointsCustomRewardRemove)
	onEventChannelChannelPointsCustomRewardRedemptionAdd    func(event EventChannelChannelPointsCustomRewardRedemptionAdd)
	onEventChannelChannelPointsCustomRewardRedemptionUpdate func(event EventChannelChannelPointsCustomRewardRedemptionUpdate)
	onEventChannelPollBegin                                 func(event EventChannelPollBegin)
	onEventChannelPollProgress                              func(event EventChannelPollProgress)
	onEventChannelPollEnd                                   func(event EventChannelPollEnd)
	onEventChannelPredictionBegin                           func(event EventChannelPredictionBegin)
	onEventChannelPredictionProgress                        func(event EventChannelPredictionProgress)
	onEventChannelPredictionLock                            func(event EventChannelPredictionLock)
	onEventChannelPredictionEnd                             func(event EventChannelPredictionEnd)
	onEventDropEntitlementGrant                             func(event []EventDropEntitlementGrant)
	onEventExtensionBitsTransactionCreate                   func(event EventExtensionBitsTransactionCreate)
	onEventChannelGoalBegin                                 func(event EventChannelGoalBegin)
	onEventChannelGoalProgress                              func(event EventChannelGoalProgress)
	onEventChannelGoalEnd                                   func(event EventChannelGoalEnd)
	onEventChannelHypeTrainBegin                            func(event EventChannelHypeTrainBegin)
	onEventChannelHypeTrainProgress                         func(event EventChannelHypeTrainProgress)
	onEventChannelHypeTrainEnd                              func(event EventChannelHypeTrainEnd)
	onEventStreamOnline                                     func(event EventStreamOnline)
	onEventStreamOffline                                    func(event EventStreamOffline)
	onEventUserAuthorizationGrant                           func(event EventUserAuthorizationGrant)
	onEventUserAuthorizationRevoke                          func(event EventUserAuthorizationRevoke)
	onEventUserUpdate                                       func(event EventUserUpdate)
	onEventChannelCharityCampaignDonate                     func(event EventChannelCharityCampaignDonate)
	onEventChannelCharityCampaignProgress                   func(event EventChannelCharityCampaignProgress)
	onEventChannelCharityCampaignStart                      func(event EventChannelCharityCampaignStart)
	onEventChannelCharityCampaignStop                       func(event EventChannelCharityCampaignStop)
	onEventChannelShieldModeBegin                           func(event EventChannelShieldModeBegin)
	onEventChannelShieldModeEnd                             func(event EventChannelShieldModeEnd)
	onEventChannelShoutoutCreate                            func(event EventChannelShoutoutCreate)
	onEventChannelShoutoutReceive                           func(event EventChannelShoutoutReceive)
	onEventAutomodMessageHold                               func(event EventAutomodMessageHold)
	onEventAutomodMessageUpdate                             func(event EventAutomodMessageUpdate)
	onEventAutomodSettingsUpdate                            func(event EventAutomodSettingsUpdate)
	onEventAutomodTermsUpdate                               func(event EventAutomodTermsUpdate)
	onEventChannelChatUserMessageHold                       func(event EventChannelChatUserMessageHold)
	onEventChannelChatUserMessageUpdate                     func(event EventChannelChatUserMessageUpdate)
	onEventChannelChatClear                                 func(event EventChannelChatClear)
	onEventChannelChatClearUserMessages                     func(event EventChannelChatClearUserMessages)
	onEventChannelChatMessage                               func(event EventChannelChatMessage)
	onEventChannelChatMessageDelete                         func(event EventChannelChatMessageDelete)
	onEventChannelChatNotification                          func(event EventChannelChatNotification)
	onEventChannelChatSettingsUpdate                        func(event EventChannelChatSettingsUpdate)
	onEventChannelSuspiciousUserMessage                     func(event EventChannelSuspiciousUserMessage)
	onEventChannelSharedChatBegin                           func(event EventChannelSharedChatBegin)
	onEventChannelSharedChatUpdate                          func(event EventChannelSharedChatUpdate)
	onEventChannelSharedChatEnd                             func(event EventChannelSharedChatEnd)
	onEventUserWhisperMessage                               func(event EventUserWhisperMessage)
}

func NewClient() *Client {
	return NewClientWithUrl(twitchWebsocketUrl)
}

func NewClientWithUrl(url string) *Client {
	return &Client{
		Address:     url,
		reconnected: make(chan struct{}),
		onError:     func(err error) { fmt.Printf("ERROR: %v\n", err) },
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
	ws, err := c.dial()
	if err != nil {
		return err
	}
	c.ws = ws
	c.connected = true

	for {
		_, data, err := c.ws.Read(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return nil
			}

			if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
				if c.reconnecting {
					c.reconnecting = false
					<-c.reconnected
					continue
				}
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
	defer func() { c.ws = nil }()
	if !c.connected {
		return nil
	}
	c.connected = false

	err := c.ws.Close(websocket.StatusNormalClosure, "Stopping Connection")

	var closeError websocket.CloseError
	if err != nil && !errors.As(err, &closeError) {
		return fmt.Errorf("could not close websocket connection: %w", err)
	}
	return nil
}

func (c *Client) handleMessage(data []byte) error {
	metadata, err := parseBaseMessage(data)
	if err != nil {
		return err
	}

	messageType := metadata.MessageType
	genMessage, ok := messageTypeMap[messageType]
	if !ok {
		return fmt.Errorf("unknown message type %s: %s", messageType, string(data))
	}

	message := genMessage()
	err = json.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("could not unmarshal message into %s: %w", messageType, err)
	}

	switch msg := message.(type) {
	case *WelcomeMessage:
		callFunc(c.onWelcome, *msg)
	case *KeepAliveMessage:
		callFunc(c.onKeepAlive, *msg)
	case *NotificationMessage:
		callFunc(c.onNotification, *msg)

		err = c.handleNotification(*msg)
		if err != nil {
			return fmt.Errorf("could not handle notification: %w", err)
		}
	case *ReconnectMessage:
		callFunc(c.onReconnect, *msg)

		err = c.reconnect(*msg)
		if err != nil {
			return fmt.Errorf("could not handle reconnect: %w", err)
		}
	case *RevokeMessage:
		callFunc(c.onRevoke, *msg)
	default:
		return fmt.Errorf("unhandled %T message: %v", msg, msg)
	}

	return nil
}

func (c *Client) reconnect(message ReconnectMessage) error {
	c.Address = message.Payload.Session.ReconnectUrl
	ws, err := c.dial()
	if err != nil {
		return fmt.Errorf("could not dial to reconnect")
	}

	go func() {
		_, data, err := ws.Read(c.ctx)
		if err != nil {
			c.onError(fmt.Errorf("reconnect failed: could not read reconnect websocket for welcome: %w", err))
		}

		metadata, err := parseBaseMessage(data)
		if err != nil {
			c.onError(fmt.Errorf("reconnect failed: could parse base message: %w", err))
		}

		if metadata.MessageType != "session_welcome" {
			c.onError(fmt.Errorf("reconnect failed: did not get a session_welcome message first: got message %s", metadata.MessageType))
			return
		}

		c.reconnecting = true
		c.ws.Close(websocket.StatusNormalClosure, "Stopping Connection")
		c.ws = ws
		c.reconnected <- struct{}{}
	}()

	return nil
}

func (c *Client) handleNotification(message NotificationMessage) error {
	data, err := message.Payload.Event.MarshalJSON()
	if err != nil {
		return fmt.Errorf("could not get event json: %w", err)
	}

	subscription := message.Payload.Subscription
	metadata, ok := subMetadata[subscription.Type]
	if !ok {
		return fmt.Errorf("unknown subscription type %s", subscription.Type)
	}

	if c.onRawEvent != nil {
		c.onRawEvent(string(data), message.Metadata, subscription)
	}

	var newEvent any
	if metadata.EventGen != nil {
		newEvent = metadata.EventGen()
		err = json.Unmarshal(data, newEvent)
		if err != nil {
			return fmt.Errorf("could not unmarshal %s into %T: %w", subscription.Type, newEvent, err)
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
	case *EventChannelCheer:
		callFunc(c.onEventChannelCheer, *event)
	case *EventChannelRaid:
		callFunc(c.onEventChannelRaid, *event)
	case *EventChannelBan:
		callFunc(c.onEventChannelBan, *event)
	case *EventChannelUnban:
		callFunc(c.onEventChannelUnban, *event)
	case *EventChannelModeratorAdd:
		callFunc(c.onEventChannelModeratorAdd, *event)
	case *EventChannelModeratorRemove:
		callFunc(c.onEventChannelModeratorRemove, *event)
	case *EventChannelChannelPointsCustomRewardAdd:
		callFunc(c.onEventChannelChannelPointsCustomRewardAdd, *event)
	case *EventChannelChannelPointsCustomRewardUpdate:
		callFunc(c.onEventChannelChannelPointsCustomRewardUpdate, *event)
	case *EventChannelChannelPointsCustomRewardRemove:
		callFunc(c.onEventChannelChannelPointsCustomRewardRemove, *event)
	case *EventChannelChannelPointsCustomRewardRedemptionAdd:
		callFunc(c.onEventChannelChannelPointsCustomRewardRedemptionAdd, *event)
	case *EventChannelChannelPointsCustomRewardRedemptionUpdate:
		callFunc(c.onEventChannelChannelPointsCustomRewardRedemptionUpdate, *event)
	case *EventChannelPollBegin:
		callFunc(c.onEventChannelPollBegin, *event)
	case *EventChannelPollProgress:
		callFunc(c.onEventChannelPollProgress, *event)
	case *EventChannelPollEnd:
		callFunc(c.onEventChannelPollEnd, *event)
	case *EventChannelPredictionBegin:
		callFunc(c.onEventChannelPredictionBegin, *event)
	case *EventChannelPredictionProgress:
		callFunc(c.onEventChannelPredictionProgress, *event)
	case *EventChannelPredictionLock:
		callFunc(c.onEventChannelPredictionLock, *event)
	case *EventChannelPredictionEnd:
		callFunc(c.onEventChannelPredictionEnd, *event)
	case *[]EventDropEntitlementGrant:
		callFunc(c.onEventDropEntitlementGrant, *event)
	case *EventExtensionBitsTransactionCreate:
		callFunc(c.onEventExtensionBitsTransactionCreate, *event)
	case *EventChannelGoalBegin:
		callFunc(c.onEventChannelGoalBegin, *event)
	case *EventChannelGoalProgress:
		callFunc(c.onEventChannelGoalProgress, *event)
	case *EventChannelGoalEnd:
		callFunc(c.onEventChannelGoalEnd, *event)
	case *EventChannelHypeTrainBegin:
		callFunc(c.onEventChannelHypeTrainBegin, *event)
	case *EventChannelHypeTrainProgress:
		callFunc(c.onEventChannelHypeTrainProgress, *event)
	case *EventChannelHypeTrainEnd:
		callFunc(c.onEventChannelHypeTrainEnd, *event)
	case *EventStreamOnline:
		callFunc(c.onEventStreamOnline, *event)
	case *EventStreamOffline:
		callFunc(c.onEventStreamOffline, *event)
	case *EventUserAuthorizationGrant:
		callFunc(c.onEventUserAuthorizationGrant, *event)
	case *EventUserAuthorizationRevoke:
		callFunc(c.onEventUserAuthorizationRevoke, *event)
	case *EventUserUpdate:
		callFunc(c.onEventUserUpdate, *event)
	case *EventChannelCharityCampaignDonate:
		callFunc(c.onEventChannelCharityCampaignDonate, *event)
	case *EventChannelCharityCampaignProgress:
		callFunc(c.onEventChannelCharityCampaignProgress, *event)
	case *EventChannelCharityCampaignStart:
		callFunc(c.onEventChannelCharityCampaignStart, *event)
	case *EventChannelCharityCampaignStop:
		callFunc(c.onEventChannelCharityCampaignStop, *event)
	case *EventChannelShieldModeBegin:
		callFunc(c.onEventChannelShieldModeBegin, *event)
	case *EventChannelShieldModeEnd:
		callFunc(c.onEventChannelShieldModeEnd, *event)
	case *EventChannelShoutoutCreate:
		callFunc(c.onEventChannelShoutoutCreate, *event)
	case *EventChannelShoutoutReceive:
		callFunc(c.onEventChannelShoutoutReceive, *event)
	case *EventAutomodMessageHold:
		callFunc(c.onEventAutomodMessageHold, *event)
	case *EventAutomodMessageUpdate:
		callFunc(c.onEventAutomodMessageUpdate, *event)
	case *EventAutomodSettingsUpdate:
		callFunc(c.onEventAutomodSettingsUpdate, *event)
	case *EventAutomodTermsUpdate:
		callFunc(c.onEventAutomodTermsUpdate, *event)
	case *EventChannelChatUserMessageHold:
		callFunc(c.onEventChannelChatUserMessageHold, *event)
	case *EventChannelChatUserMessageUpdate:
		callFunc(c.onEventChannelChatUserMessageUpdate, *event)
	case *EventChannelChatClear:
		callFunc(c.onEventChannelChatClear, *event)
	case *EventChannelChatClearUserMessages:
		callFunc(c.onEventChannelChatClearUserMessages, *event)
	case *EventChannelChatMessage:
		callFunc(c.onEventChannelChatMessage, *event)
	case *EventChannelChatMessageDelete:
		callFunc(c.onEventChannelChatMessageDelete, *event)
	case *EventChannelChatNotification:
		callFunc(c.onEventChannelChatNotification, *event)
	case *EventChannelChatSettingsUpdate:
		callFunc(c.onEventChannelChatSettingsUpdate, *event)
	case *EventChannelSuspiciousUserMessage:
		callFunc(c.onEventChannelSuspiciousUserMessage, *event)
	case *EventChannelSharedChatBegin:
		callFunc(c.onEventChannelSharedChatBegin, *event)
	case *EventChannelSharedChatUpdate:
		callFunc(c.onEventChannelSharedChatUpdate, *event)
	case *EventChannelSharedChatEnd:
		callFunc(c.onEventChannelSharedChatEnd, *event)
	case *EventUserWhisperMessage:
		callFunc(c.onEventUserWhisperMessage, *event)
	default:
		c.onError(fmt.Errorf("unknown event type %s", subscription.Type))
	}

	return nil
}

func (c *Client) dial() (*websocket.Conn, error) {
	ws, _, err := websocket.Dial(c.ctx, c.Address, nil)
	if err != nil {
		return nil, fmt.Errorf("could not dial %s: %w", c.Address, err)
	}
	return ws, nil
}

func parseBaseMessage(data []byte) (MessageMetadata, error) {
	type BaseMessage struct {
		Metadata MessageMetadata `json:"metadata"`
	}

	var baseMessage BaseMessage
	err := json.Unmarshal(data, &baseMessage)
	if err != nil {
		return MessageMetadata{}, fmt.Errorf("could not unmarshal basemessage to get message type: %w", err)
	}

	return baseMessage.Metadata, nil
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

func (c *Client) OnRawEvent(callback func(event string, metadata MessageMetadata, subscription PayloadSubscription)) {
	c.onRawEvent = callback
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

func (c *Client) OnEventChannelCheer(callback func(event EventChannelCheer)) {
	c.onEventChannelCheer = callback
}

func (c *Client) OnEventChannelRaid(callback func(event EventChannelRaid)) {
	c.onEventChannelRaid = callback
}

func (c *Client) OnEventChannelBan(callback func(event EventChannelBan)) {
	c.onEventChannelBan = callback
}

func (c *Client) OnEventChannelUnban(callback func(event EventChannelUnban)) {
	c.onEventChannelUnban = callback
}

func (c *Client) OnEventChannelModeratorAdd(callback func(event EventChannelModeratorAdd)) {
	c.onEventChannelModeratorAdd = callback
}

func (c *Client) OnEventChannelModeratorRemove(callback func(event EventChannelModeratorRemove)) {
	c.onEventChannelModeratorRemove = callback
}

func (c *Client) OnEventChannelChannelPointsCustomRewardAdd(callback func(event EventChannelChannelPointsCustomRewardAdd)) {
	c.onEventChannelChannelPointsCustomRewardAdd = callback
}

func (c *Client) OnEventChannelChannelPointsCustomRewardUpdate(callback func(event EventChannelChannelPointsCustomRewardUpdate)) {
	c.onEventChannelChannelPointsCustomRewardUpdate = callback
}

func (c *Client) OnEventChannelChannelPointsCustomRewardRemove(callback func(event EventChannelChannelPointsCustomRewardRemove)) {
	c.onEventChannelChannelPointsCustomRewardRemove = callback
}

func (c *Client) OnEventChannelChannelPointsCustomRewardRedemptionAdd(callback func(event EventChannelChannelPointsCustomRewardRedemptionAdd)) {
	c.onEventChannelChannelPointsCustomRewardRedemptionAdd = callback
}

func (c *Client) OnEventChannelChannelPointsCustomRewardRedemptionUpdate(callback func(event EventChannelChannelPointsCustomRewardRedemptionUpdate)) {
	c.onEventChannelChannelPointsCustomRewardRedemptionUpdate = callback
}

func (c *Client) OnEventChannelPollBegin(callback func(event EventChannelPollBegin)) {
	c.onEventChannelPollBegin = callback
}

func (c *Client) OnEventChannelPollProgress(callback func(event EventChannelPollProgress)) {
	c.onEventChannelPollProgress = callback
}

func (c *Client) OnEventChannelPollEnd(callback func(event EventChannelPollEnd)) {
	c.onEventChannelPollEnd = callback
}

func (c *Client) OnEventChannelPredictionBegin(callback func(event EventChannelPredictionBegin)) {
	c.onEventChannelPredictionBegin = callback
}

func (c *Client) OnEventChannelPredictionProgress(callback func(event EventChannelPredictionProgress)) {
	c.onEventChannelPredictionProgress = callback
}

func (c *Client) OnEventChannelPredictionLock(callback func(event EventChannelPredictionLock)) {
	c.onEventChannelPredictionLock = callback
}

func (c *Client) OnEventChannelPredictionEnd(callback func(event EventChannelPredictionEnd)) {
	c.onEventChannelPredictionEnd = callback
}

func (c *Client) OnEventDropEntitlementGrant(callback func(event []EventDropEntitlementGrant)) {
	c.onEventDropEntitlementGrant = callback
}

func (c *Client) OnEventExtensionBitsTransactionCreate(callback func(event EventExtensionBitsTransactionCreate)) {
	c.onEventExtensionBitsTransactionCreate = callback
}

func (c *Client) OnEventChannelGoalBegin(callback func(event EventChannelGoalBegin)) {
	c.onEventChannelGoalBegin = callback
}

func (c *Client) OnEventChannelGoalProgress(callback func(event EventChannelGoalProgress)) {
	c.onEventChannelGoalProgress = callback
}

func (c *Client) OnEventChannelGoalEnd(callback func(event EventChannelGoalEnd)) {
	c.onEventChannelGoalEnd = callback
}

func (c *Client) OnEventChannelHypeTrainBegin(callback func(event EventChannelHypeTrainBegin)) {
	c.onEventChannelHypeTrainBegin = callback
}

func (c *Client) OnEventChannelHypeTrainProgress(callback func(event EventChannelHypeTrainProgress)) {
	c.onEventChannelHypeTrainProgress = callback
}

func (c *Client) OnEventChannelHypeTrainEnd(callback func(event EventChannelHypeTrainEnd)) {
	c.onEventChannelHypeTrainEnd = callback
}

func (c *Client) OnEventStreamOnline(callback func(event EventStreamOnline)) {
	c.onEventStreamOnline = callback
}

func (c *Client) OnEventStreamOffline(callback func(event EventStreamOffline)) {
	c.onEventStreamOffline = callback
}

func (c *Client) OnEventUserAuthorizationGrant(callback func(event EventUserAuthorizationGrant)) {
	c.onEventUserAuthorizationGrant = callback
}

func (c *Client) OnEventUserAuthorizationRevoke(callback func(event EventUserAuthorizationRevoke)) {
	c.onEventUserAuthorizationRevoke = callback
}

func (c *Client) OnEventUserUpdate(callback func(event EventUserUpdate)) {
	c.onEventUserUpdate = callback
}

func (c *Client) OnEventChannelCharityCampaignDonate(callback func(event EventChannelCharityCampaignDonate)) {
	c.onEventChannelCharityCampaignDonate = callback
}

func (c *Client) OnEventChannelCharityCampaignProgress(callback func(event EventChannelCharityCampaignProgress)) {
	c.onEventChannelCharityCampaignProgress = callback
}

func (c *Client) OnEventChannelCharityCampaignStart(callback func(event EventChannelCharityCampaignStart)) {
	c.onEventChannelCharityCampaignStart = callback
}

func (c *Client) OnEventChannelCharityCampaignStop(callback func(event EventChannelCharityCampaignStop)) {
	c.onEventChannelCharityCampaignStop = callback
}

func (c *Client) OnEventChannelShieldModeBegin(callback func(event EventChannelShieldModeBegin)) {
	c.onEventChannelShieldModeBegin = callback
}

func (c *Client) OnEventChannelShieldModeEnd(callback func(event EventChannelShieldModeEnd)) {
	c.onEventChannelShieldModeEnd = callback
}

func (c *Client) OnEventChannelShoutoutCreate(callback func(event EventChannelShoutoutCreate)) {
	c.onEventChannelShoutoutCreate = callback
}

func (c *Client) OnEventChannelShoutoutReceive(callback func(event EventChannelShoutoutReceive)) {
	c.onEventChannelShoutoutReceive = callback
}

func (c *Client) OnEventAutomodMessageHold(callback func(event EventAutomodMessageHold)) {
	c.onEventAutomodMessageHold = callback
}

func (c *Client) OnEventAutomodMessageUpdate(callback func(event EventAutomodMessageUpdate)) {
	c.onEventAutomodMessageUpdate = callback
}

func (c *Client) OnEventAutomodSettingsUpdate(callback func(event EventAutomodSettingsUpdate)) {
	c.onEventAutomodSettingsUpdate = callback
}

func (c *Client) OnEventAutomodTermsUpdate(callback func(event EventAutomodTermsUpdate)) {
	c.onEventAutomodTermsUpdate = callback
}

func (c *Client) OnEventChannelChatUserMessageHold(callback func(event EventChannelChatUserMessageHold)) {
	c.onEventChannelChatUserMessageHold = callback
}

func (c *Client) OnEventChannelChatUserMessageUpdate(callback func(event EventChannelChatUserMessageUpdate)) {
	c.onEventChannelChatUserMessageUpdate = callback
}

func (c *Client) OnEventChannelChatClear(callback func(event EventChannelChatClear)) {
	c.onEventChannelChatClear = callback
}

func (c *Client) OnEventChannelChatClearUserMessages(callback func(event EventChannelChatClearUserMessages)) {
	c.onEventChannelChatClearUserMessages = callback
}

func (c *Client) OnEventChannelChatMessage(callback func(event EventChannelChatMessage)) {
	c.onEventChannelChatMessage = callback
}

func (c *Client) OnEventChannelChatMessageDelete(callback func(event EventChannelChatMessageDelete)) {
	c.onEventChannelChatMessageDelete = callback
}

func (c *Client) OnEventChannelChatNotification(callback func(event EventChannelChatNotification)) {
	c.onEventChannelChatNotification = callback
}

func (c *Client) OnEventChannelChatSettingsUpdate(callback func(event EventChannelChatSettingsUpdate)) {
	c.onEventChannelChatSettingsUpdate = callback
}

func (c *Client) OnEventChannelSuspiciousUserMessage(callback func(event EventChannelSuspiciousUserMessage)) {
	c.onEventChannelSuspiciousUserMessage = callback
}

func (c *Client) OnEventChannelSharedChatBegin(callback func(event EventChannelSharedChatBegin)) {
	c.onEventChannelSharedChatBegin = callback
}

func (c *Client) OnEventChannelSharedChatUpdate(callback func(event EventChannelSharedChatUpdate)) {
	c.onEventChannelSharedChatUpdate = callback
}

func (c *Client) OnEventChannelSharedChatEnd(callback func(event EventChannelSharedChatEnd)) {
	c.onEventChannelSharedChatEnd = callback
}

func (c *Client) OnEventUserWhisperMessage(callback func(event EventUserWhisperMessage)) {
	c.onEventUserWhisperMessage = callback
}
