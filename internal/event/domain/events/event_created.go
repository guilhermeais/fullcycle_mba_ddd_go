package events

import (
	"ingressos/internal/common"
	"time"
)

const EventCreatedEventName = common.EventName("EventCreatedEvent")

type EventCreatedEvent struct {
	Id                 common.UUID
	Name, Description  string
	Date               time.Time
	IsPublished        bool
	TotalSpots         int
	TotalSpotsReserved int
	PartnerId          common.UUID
}

func (c EventCreatedEvent) GetEventName() common.EventName {
	return CustomerCreatedEventName
}
