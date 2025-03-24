package events

import "ingressos/internal/common"

const CustomerCreatedEventName = common.EventName("CustomerCreatedEvent")

type CustomerCreatedEvent struct {
	ID                         common.UUID
	Name, Email, CPF, Birthday string
}

func (c CustomerCreatedEvent) GetEventName() common.EventName {
	return CustomerCreatedEventName
}
