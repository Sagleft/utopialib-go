package utopiago

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
