package websocket

import (
	"encoding/json"
	"errors"

	"github.com/Sagleft/evtwebsocket"
)

type wsHandler struct {
	url  string
	conn evtwebsocket.Conn
	task WsSubscribeTask
}

func NewWsHandler(URL string, task WsSubscribeTask) Handler {
	h := &wsHandler{
		url:  URL,
		conn: evtwebsocket.Conn{},
		task: task,
	}

	h.conn.OnConnected = h.onConnected
	h.conn.OnMessage = h.onMessage
	h.conn.OnError = h.onError

	if !task.DisablePing {
		h.conn.PingIntervalSecs = 5     // ping interval in seconds
		h.conn.PingMsg = []byte("PING") // ping message to send
	}
	return h
}

func (h *wsHandler) Close() error {
	return h.conn.Close()
}

func (h *wsHandler) Connect() error {
	return h.conn.Dial(h.url, "")
}

func newEvent(jsonRaw []byte) (WsEvent, error) {
	event := WsEvent{}
	err := json.Unmarshal(jsonRaw, &event)
	if err != nil {
		return event, errors.New("failed to decode event json: " + err.Error())
	}
	return event, nil
}

// Fires when the connection is established
func (h *wsHandler) onConnected(w *evtwebsocket.Conn) {
	h.task.OnConnected()
}

// Fires when a new message arrives from the server
func (h *wsHandler) onMessage(msg []byte, w *evtwebsocket.Conn) {
	event, err := newEvent(msg)
	if err == nil {
		go h.task.Callback(event)
	} else {
		go h.task.ErrCallback(err)
	}
}

// Fires when an error occurs and connection is closed
func (h *wsHandler) onError(err error) {
	h.task.ErrCallback(err)
}
