package events

import (
	"ingressos/internal/common"
)

const EventPublishedEventName = common.EventName("EventPublishedEvent")

type EventPublishedEvent struct {
	EventId  common.UUID
	Sections []common.UUID
}

func (c EventPublishedEvent) GetEventName() common.EventName {
	return EventPublishedEventName
}
