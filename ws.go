package utopiago

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/Sagleft/evtwebsocket"
)

// GetString - get string field from ws event.
// throw error when not found or is not convertable to this type
func (ws *WsEvent) GetString(field string) (string, error) {
	errHandler := func(err error) (string, error) {
		return "", err
	}

	valRaw, isFound := ws.Data[field]
	if !isFound {
		return errHandler(errors.New("field `" + field + "` not found"))
	}

	val, isConvertable := valRaw.(string)
	if !isConvertable {
		return errHandler(errors.New("field `" + field + "` type is `" + reflect.ValueOf(valRaw).String() + "` not a string"))
	}

	return val, nil
}

// GetBool - get bool field from ws event.
// throw error when not found or is not convertable to this type
func (ws *WsEvent) GetBool(field string) (bool, error) {
	errHandler := func(err error) (bool, error) {
		return false, err
	}

	valRaw, isFound := ws.Data[field]
	if !isFound {
		return errHandler(errors.New("field `" + field + "` not found"))
	}

	val, isConvertable := valRaw.(bool)
	if !isConvertable {
		return errHandler(errors.New("field `" + field + "` type is `" + reflect.ValueOf(valRaw).String() + "` not a bool"))
	}

	return val, nil
}

// GetInt - get int64 field from ws event.
// throw error when not found or is not convertable to this type
func (ws *WsEvent) GetInt(field string) (int64, error) {
	errHandler := func(err error) (int64, error) {
		return 0, err
	}

	valRaw, isFound := ws.Data[field]
	if !isFound {
		return errHandler(errors.New("field `" + field + "` not found"))
	}

	val, isConvertable := valRaw.(int64)
	if !isConvertable {
		return errHandler(errors.New("field `" + field + "` type is `" + reflect.ValueOf(valRaw).String() + "` not a int64"))
	}

	return val, nil
}

// GetFloat - get int64 field from ws event.
// throw error when not found or is not convertable to this type
func (ws *WsEvent) GetFloat(field string) (float64, error) {
	errHandler := func(err error) (float64, error) {
		return 0, err
	}

	valRaw, isFound := ws.Data[field]
	if !isFound {
		return errHandler(errors.New("field `" + field + "` not found"))
	}

	val, isConvertable := valRaw.(float64)
	if !isConvertable {
		return errHandler(errors.New("field `" + field + "` type is `" + reflect.ValueOf(valRaw).String() + "` not a float64"))
	}

	return val, nil
}

// GetChannelMessage - get the event data converted to ChannelMessage.
// actual only for `newPrivateChannelMessage` and `newChannelMessage` events
func (ws *WsEvent) GetChannelMessage() (WsChannelMessage, error) {
	result := WsChannelMessage{}
	eventBytes, err := json.Marshal(ws.Data)
	if err != nil {
		return result, errors.New("failed to encode channel message: " + err.Error())
	}

	err = json.Unmarshal(eventBytes, &result)
	if err != nil {
		return result, errors.New("failed to decode event data as channel message: " + err.Error())
	}
	return result, nil
}

// GetInstantMessage - get the event data converted to InstantMessage.
// actual only for `newInstantMessage` event
func (ws *WsEvent) GetInstantMessage() (InstantMessage, error) {
	result := InstantMessage{}
	eventBytes, err := json.Marshal(ws.Data)
	if err != nil {
		return result, errors.New("failed to encode contact message: " + err.Error())
	}

	err = json.Unmarshal(eventBytes, &result)
	if err != nil {
		return result, errors.New("failed to decode event data as contact message: " + err.Error())
	}
	return result, nil
}

func newWsEvent(jsonRaw []byte) (WsEvent, error) {
	event := WsEvent{}
	err := json.Unmarshal(jsonRaw, &event)
	if err != nil {
		return event, errors.New("failed to decode event json: " + err.Error())
	}
	return event, nil
}

func (c *UtopiaClient) newWsHandler(task WsSubscribeTask) *wsHandler {
	h := wsHandler{
		WsURL: c.getWsURL(),
		Conn:  evtwebsocket.Conn{},
		Task:  task,
	}

	h.Conn.OnConnected = h.onConnected
	h.Conn.OnMessage = h.onMessage
	h.Conn.OnError = h.onError

	if !task.DisablePing {
		h.Conn.PingIntervalSecs = 5     // ping interval in seconds
		h.Conn.PingMsg = []byte("PING") // ping message to send
	}
	return &h
}

// Fires when the connection is established
func (h *wsHandler) onConnected(w *evtwebsocket.Conn) {
	h.Task.OnConnected()
}

// Fires when a new message arrives from the server
func (h *wsHandler) onMessage(msg []byte, w *evtwebsocket.Conn) {
	event, err := newWsEvent(msg)
	if err == nil {
		go h.Task.Callback(event)
	} else {
		go h.Task.ErrCallback(err)
	}
}

// Fires when an error occurs and connection is closed
func (h *wsHandler) onError(err error) {
	h.Task.ErrCallback(err)
}

// NOTE: it's blocking method
func (h *wsHandler) connect() error {
	// open connection
	return h.Conn.Dial(h.WsURL, "")
}

// WsSubscribe - connect to websocket & recive messages.
// NOTE: it's blocking method
func (c *UtopiaClient) WsSubscribe(task WsSubscribeTask) error {
	return c.newWsHandler(task).connect()
}
