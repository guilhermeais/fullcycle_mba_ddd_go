package events

import "ingressos/internal/common"

const CustomerUpdatedEventName = common.EventName("CustomerUpdatedEvent")

type CustomerUpdatedEvent struct {
	ID, Name, Birthday string
}

func (c CustomerUpdatedEvent) GetEventName() common.EventName {
	return CustomerUpdatedEventName
}
