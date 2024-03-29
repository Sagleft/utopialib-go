package websocket

// WsEvent - websocket event
type WsEvent struct {
	Type string                 `json:"type"`
	Data map[string]interface{} `json:"data"`
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
