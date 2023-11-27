package bitvavo

import "github.com/orsinium-labs/enum"

const zerof = float64(0)

type WsEvent enum.Member[string]

var (
	wsEventSubscribed   = WsEvent{"subscribed"}
	wsEventUnsubscribed = WsEvent{"unsubscribed"}
	wsEventCandles      = WsEvent{"candle"}
	wsEventTicker       = WsEvent{"ticker"}
	wsEventTicker24h    = WsEvent{"ticker24h"}
	wsEventTrades       = WsEvent{"trade"}
	wsEventBook         = WsEvent{"book"}
	wsEventAuth         = WsEvent{"authenticate"}
	wsEventOrder        = WsEvent{"order"}
	wsEventFill         = WsEvent{"fill"}
)

type Action enum.Member[string]

var (
	actionSubscribe    = Action{"subscribe"}
	actionUnsubscribe  = Action{"unsubscribe"}
	actionAuthenticate = Action{"authenticate"}
)

type ChannelName enum.Member[string]

var (
	channelNameCandles   = ChannelName{"candles"}
	channelNameTicker    = ChannelName{"ticker"}
	channelNameTicker24h = ChannelName{"ticker24h"}
	channelNameTrades    = ChannelName{"trades"}
	channelNameBook      = ChannelName{"book"}
	channelNameAccount   = ChannelName{"account"}
)

type AuthEvent struct {
	// Describes the returned event over the socket.
	Event string `json:"event"`

	// Whether the user is authenticated.
	Authenticated bool `json:"authenticated"`
}

type BaseEvent struct {
	Event string `json:"event"`
}

type WebSocketErr struct {
	Action  string `json:"action"`
	Code    int    `json:"errorCode"`
	Message string `json:"error"`
}

type WebSocketMessage struct {
	Action   string    `json:"action"`
	Channels []Channel `json:"channels,omitempty"`

	// Api Key.
	Key string `json:"key,omitempty"`
	// SHA256 HMAC hex digest of timestamp + method + url + body.
	Signature string `json:"signature,omitempty"`
	// The current timestamp in milliseconds since 1 Jan 1970.
	Timestamp int64 `json:"timestamp,omitempty"`
	// The window that allows execution of your request in milliseconds since 1 Jan 1970. The default value is 10000 (10s) and maximum value is 60000 (60s).
	Window uint64 `json:"window,omitempty"`
}

type Channel struct {
	Name      string   `json:"name"`
	Intervals []string `json:"interval,omitempty"`
	Markets   []string `json:"markets,omitempty"`
}
