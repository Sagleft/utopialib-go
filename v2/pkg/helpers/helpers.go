package helpers

import (
	"encoding/json"
	"errors"

	"github.com/Sagleft/utopialib-go/v2/pkg/structs"
	"github.com/Sagleft/utopialib-go/v2/pkg/websocket"
)

// GetChannelMessageFromEvent - get the event data converted to ChannelMessage.
// actual only for `newPrivateChannelMessage` and `newChannelMessage` events
func GetChannelMessageFromEvent(ws websocket.WsEvent) (structs.WsChannelMessage, error) {
	result := structs.WsChannelMessage{}
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

// GetInstantMessageFromEvent - get the event data converted to InstantMessage.
// actual only for `newInstantMessage` event
func GetInstantMessageFromEvent(ws websocket.WsEvent) (structs.InstantMessage, error) {
	result := structs.InstantMessage{}
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
