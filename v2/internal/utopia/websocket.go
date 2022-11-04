package utopia

import (
	"encoding/json"
	"errors"

	"github.com/Sagleft/utopialib-go/v2/pkg/structs"
	"github.com/Sagleft/utopialib-go/v2/pkg/websocket"
)

// WsSubscribe - connect to websocket & receive messages.
// NOTE: it's blocking method
func (c *UtopiaClient) WsSubscribe(task websocket.WsSubscribeTask) error {
	return websocket.NewWsHandler(c.getWsURL(), task).Connect()
}

// ParseWsChannelMessage - get the event data converted to ChannelMessage.
// actual only for `newPrivateChannelMessage` and `newChannelMessage` events
func ParseWsChannelMessage(e *websocket.WsEvent) (structs.WsChannelMessage, error) {
	result := structs.WsChannelMessage{}
	eventBytes, err := json.Marshal(e.Data)
	if err != nil {
		return result, errors.New("failed to encode channel message: " + err.Error())
	}

	err = json.Unmarshal(eventBytes, &result)
	if err != nil {
		return result, errors.New("failed to decode event data as channel message: " + err.Error())
	}
	return result, nil
}

// ParseWsInstantMessage - get the event data converted to InstantMessage.
// actual only for `newInstantMessage` event
func ParseWsInstantMessage(e *websocket.WsEvent) (structs.InstantMessage, error) {
	result := structs.InstantMessage{}
	eventBytes, err := json.Marshal(e.Data)
	if err != nil {
		return result, errors.New("failed to encode contact message: " + err.Error())
	}

	err = json.Unmarshal(eventBytes, &result)
	if err != nil {
		return result, errors.New("failed to decode event data as contact message: " + err.Error())
	}
	return result, nil
}
