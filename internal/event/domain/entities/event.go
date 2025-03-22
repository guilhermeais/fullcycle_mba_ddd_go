package entities

import (
	"ingressos/internal/common"
	"ingressos/internal/event/domain/events"
	"time"
)

type EventId = common.UUID

type Event struct {
	common.AggregateRoot
	id                 EventId
	name               string
	description        string
	date               time.Time
	isPublished        bool
	totalSpots         int
	totalSpotsReserved int
	partnerId          PartnerId
}

type CreateEventCommand struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	IsPublished bool      `json:"isPublished"`
	TotalSpots  int       `json:"totalSpots"`
	PartnerId   string    `json:"partnerId"`
}

func CreateEvent(command CreateEventCommand) (*Event, error) {
	uuid, err := common.CreateUUID()
	if err != nil {
		return nil, err
	}

	partnerId, err := common.RestoreUUID(command.PartnerId)
	if err != nil {
		return nil, err
	}

	event := &Event{
		id:            uuid,
		name:          command.Name,
		description:   command.Description,
		date:          command.Date,
		isPublished:   command.IsPublished,
		totalSpots:    command.TotalSpots,
		partnerId:     partnerId,
		AggregateRoot: common.NewAggregateRoot(uuid),
	}

	event.AddDomainEvent(events.EventCreatedEvent{
		Id:                 string(event.id),
		Name:               event.name,
		Description:        event.description,
		Date:               event.date,
		IsPublished:        event.isPublished,
		TotalSpots:         event.totalSpots,
		TotalSpotsReserved: event.totalSpotsReserved,
		PartnerId:          string(event.partnerId),
	})

	return event, nil
}
