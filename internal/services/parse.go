package services

import (
	"encoding/json"
	"feishuBot/internal/messages"
	"fmt"
)

type EventMapping[T any] map[string]T

var (
	eventMapping EventMapping[messages.ReceiveMsgV2]
)

func ParseEvent[T any](event *messages.LarkEvent) (T, error) {
	_, ok := eventMapping[event.Header.EventType]
	if !ok {
		return *new(T), fmt.Errorf("unknown event type: %s", event.Header.EventType)
	}

	var result T
	err := json.Unmarshal(event.Event, &result)

	return result, err
}

func (e EventMapping[T]) Register(eventType string, obj T) {
	e[eventType] = obj
}

func init() {
	eventMapping = make(EventMapping[messages.ReceiveMsgV2])
	eventMapping[messages.EventTypeReceiveMsg] = messages.ReceiveMsgV2{}
}
