package utopia

import (
	"encoding/json"
	"errors"
)

// WsSubscribe - connect to websocket & receive messages.
// NOTE: it's blocking method
func (c *UtopiaClient) WsSubscribe(task WsSubscribeTask) error {
	return c.newWsHandler(task).connect()
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
