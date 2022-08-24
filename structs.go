package utopiago

import "github.com/ctengiz/evtwebsocket"

// Query is a filter for API requests
type Query struct {
	Method string                 `json:"method"`
	Token  string                 `json:"token"`
	Params map[string]interface{} `json:"params"`
}

// UtopiaClient lets you connect to Utopia Client
type UtopiaClient struct {
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Token    string `json:"token"`
	Port     int    `json:"port"`
	WsPort   int    `json:"wsport"`
}

// WsEvent - websocket event from Utopia Client
type WsEvent struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
}

type SetWsStateTask struct {
	Enabled       bool   `json:"enabled"`
	Port          int    `json:"port"`
	EnableSSL     bool   `json:"enablessl"`
	Notifications string `json:"notifications"` // example: "contact, wallet" example2: "all"
}

type WsEventsCallback func(ws WsEvent)

type WsErrorCallback func(err error)

type WsSubscribeTask struct {
	// required
	OnConnected func()           // required
	Callback    WsEventsCallback // required
	ErrCallback WsErrorCallback  // required
	Port        int

	// optional
	DisablePing bool
}

// ContactData - user contact data
type ContactData struct {
	AuthStatus int    `json:"authorizationStatus"`
	AvatarHash string `json:"avatarMd5"`
	Group      string `json:"group"`
	PubkeyHash string `json:"hashedPk"`
	IsFriend   bool   `json:"isFriend"`
	Nick       string `json:"nick"`
	Pubkey     string `json:"pk"`
	Status     int    `json:"status"` // 65536 - offline, 4096 - online, 4097 - away, 4099 - do not disturb, 32768 - invisible
}

// ChannelContactData - channel contact data
type ChannelContactData struct {
	PubkeyHash  string `json:"hashedPk"`
	LastSeen    string `json:"lastSeen"`
	IsLocal     bool   `json:"local"`
	IsModerator bool   `json:"moderator"`
	Nick        string `json:"nick"`
	Pubkey      string `json:"pk"`
}

// InstantMessage - contact message
type InstantMessage struct {
	ID               int         `json:"id"`
	DateTime         string      `json:"dateTime"`
	File             interface{} `json:"file"`
	MessageType      int         `json:"messageType"`
	Nick             string      `json:"nick"`             // message author nick
	Pubkey           string      `json:"pk"`               // can be empty
	ReadDateTime     *string     `json:"readDateTime"`     // can be nil when message is unread
	ReceivedDateTime string      `json:"receivedDateTime"` // when message delivered
	Text             string      `json:"text"`             // message text
}

// WsChannelMessage - channel message data
type WsChannelMessage struct {
	ID          int64  `json:"id"`
	ChannelName string `json:"channel"`
	ChannelID   string `json:"channelid"`
	DateTime    string `json:"dateTime"`
	PubkeyHash  string `json:"hashedPk"`
	IsIncoming  bool   `json:"isIncoming"`
	MessageType int    `json:"messageType"`
	Nick        string `json:"nick"`    // message author nick
	Pubkey      string `json:"pk"`      // can be empty
	Text        string `json:"text"`    // message text
	TopicID     string `json:"topicId"` // for reply
}

// ChannelMessage - channel message data
type ChannelMessage struct {
	ID          int64  `json:"id"`
	DateTime    string `json:"dateTime"`
	PubkeyHash  string `json:"hashedPk"`
	IsIncoming  bool   `json:"isIncoming"`
	MessageType int    `json:"messageType"`
	Nick        string `json:"nick"`    // message author nick
	Pubkey      string `json:"pk"`      // can be empty
	Text        string `json:"text"`    // message text
	TopicID     string `json:"topicId"` // for reply
}

type wsHandler struct {
	WsURL string
	Conn  evtwebsocket.Conn
	Task  WsSubscribeTask
}
