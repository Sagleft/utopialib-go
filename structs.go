package utopiago

import (
	"github.com/ctengiz/evtwebsocket"
)

// Query is a filter for API requests
type Query struct {
	Method string                 `json:"method"`
	Token  string                 `json:"token"`
	Params map[string]interface{} `json:"params"`
}

// UtopiaClient lets you connect to Utopia Client
type UtopiaClient struct {
	// protected
	logCallback LogCallback

	// required
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Token    string `json:"token"`
	Port     int    `json:"port"`

	// optional
	RequestTimeoutSeconds int `json:"timeout"`
	WsPort                int `json:"wsport"`
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

	// optional
	DisablePing bool
}

// OwnContactData - own account data
type OwnContactData struct {
	AvatarHash string `json:"avatarMd5"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	PubkeyHash string `json:"hashedPk"`
	IsFriend   bool   `json:"isFriend"`
	Mood       string `json:"moodMessage"`
	Nick       string `json:"nick"`
	Pubkey     string `json:"pk"`
	Status     int    `json:"status"`
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

type ChannelData struct {
	HideInCommonList bool   `json:"HideInCommonList"` // example: false
	CreatedOn        string `json:"created"`          // 2022-09-09T05:47:52.972Z
	ModifiedOn       string `json:"modified"`         // 2022-09-09T05:47:52.973Z
	Description      string `json:"description"`
	GeoTag           string `json:"geotag"`
	HasgTags         string `json:"hashtags"`
	Languages        string `json:"languages"`
	Owner            string `json:"owner"`
	ReadOnly         bool   `json:"readonly"`
	ReadOnlyPrivacy  bool   `json:"readonly_privacy"`
	Title            string `json:"title"`
	Type             string `json:"type"` // public
}

type SearchChannelData struct {
	AvatarID    string `json:"avatarId"`    // example: defAvatar_F10383EA72AC6263C21F356CD8D2E2A2
	ChannelID   string `json:"channelid"`   // F10383EA72AC6263C21F356CD8D2E2A2
	CreatedOn   string `json:"created"`     // 2022-01-28T16:11:39.144Z
	Description string `json:"description"` // can be empty
	IsJoined    bool   `json:"isjoined"`    // false
	IsPrivate   bool   `json:"isprivate"`   // true
	EditedOn    string `json:"modified"`    // 2022-01-28T16:11:39.145Z
	Name        string `json:"name"`        // Monica
	OwnerPubkey string `json:"owner"`       // 1B742E8D8DAE682ADD2568BE25B23F35BA7A8BFC1D5D3BCA0EE219A754A48201
}
