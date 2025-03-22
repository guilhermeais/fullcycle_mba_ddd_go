package events

import (
	"ingressos/internal/common"
	"time"
)

const EventCreatedEventName = common.EventName("EventCreatedEvent")

type EventCreatedEvent struct {
	Id                 string
	Name, Description  string
	Date               time.Time
	IsPublished        bool
	TotalSpots         int
	TotalSpotsReserved int
	PartnerId          string
}

func (c EventCreatedEvent) GetEventName() common.EventName {
	return CustomerCreatedEventName
}
