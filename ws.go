package utopiago

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"

	"github.com/rgamba/evtwebsocket"
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

type WsEventsCallback func(ws WsEvent)

type WsErrorCallback func(err error)

type WsSubscribeTask struct {
	Port        int
	Callback    WsEventsCallback // required
	ErrCallback WsErrorCallback  // required
}

/*func newWsEvent(jsonRaw string) (WsEvent, error) {
	event := WsEvent{}
	err := json.Unmarshal([]byte(jsonRaw), &event)
	if err != nil {
		return event, errors.New("failed to decode event json: " + err.Error())
	}
	return event, nil
}*/

// WsSubscribe - connect to websocket & recive messages.
// NOTE: it's blocking method
/*func (c *UtopiaClient) WsSubscribe(task WsSubscribeTask) error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// create ws
	socket := gowebsocket.New("ws://" + c.getBaseURLWithoutProtocol())

	// setup callbacks
	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		event, err := newWsEvent(message)
		if err != nil {
			task.ErrCallback(err)
		} else {
			task.Callback(event)
		}
	}

	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		task.ErrCallback(err)
	}

	// connect
	socket.Connect()

	// wait for close
	<-interrupt
	socket.Close()
	return nil
}*/

// WsSubscribe - connect to websocket & recive messages.
// NOTE: it's blocking method
func (c *UtopiaClient) WsSubscribe(task WsSubscribeTask) error {
	conn := evtwebsocket.Conn{
		// Fires when the connection is established
		OnConnected: func(w *evtwebsocket.Conn) {
			fmt.Println("Connected!")
		},
		// Fires when a new message arrives from the server
		OnMessage: func(msg []byte, w *evtwebsocket.Conn) {
			fmt.Printf("New message: %s\n", msg)
		},
		// Fires when an error occurs and connection is closed
		OnError: func(err error) {
			fmt.Printf("Error: %s\n", err.Error())
			os.Exit(1)
		},
		// Ping interval in secs (optional)
		PingIntervalSecs: 5,
		// Ping message to send (optional)
		PingMsg: []byte("PING"),
	}

	err := conn.Dial(c.getBaseURL()+":"+strconv.Itoa(task.Port), "")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
