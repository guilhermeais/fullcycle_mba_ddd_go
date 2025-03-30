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
	sections           []*EventSection
}

func (e Event) GetID() EventId {
	return e.id
}

func (e Event) GetName() string {
	return e.name
}

func (e Event) GetDescription() string {
	return e.description
}

func (e Event) GetDate() time.Time {
	return e.date
}

func (e Event) IsPublished() bool {
	return e.isPublished
}

func (e Event) GetTotalSpots() int {
	return e.totalSpots
}

func (e Event) GetTotalSpotsReserved() int {
	return e.totalSpotsReserved
}

func (e Event) GetPartnerId() PartnerId {
	return e.partnerId
}

func (e *Event) Publish() {
	e.isPublished = true
	var sectionIds = make([]EventSectionId, 0, len(e.sections))
	for _, section := range e.sections {
		section.Publish()
		sectionIds = append(sectionIds, section.id)
	}
	e.AddDomainEvent(events.EventPublishedEvent{
		EventId:  e.id,
		Sections: sectionIds,
	})
}

func (e *Event) AddSection(c CreateEventSectionCommand) error {
	section, err := CreateEventSection(c)

	if err != nil {
		return err
	}

	e.totalSpots += c.TotalSpots
	e.sections = append(e.sections, section)
	return nil
}

func (e *Event) GetSections() []EventSection {
	sections := make([]EventSection, len(e.sections))
	for i, section := range e.sections {
		sections[i] = *section
	}
	return sections
}

type CreateEventCommand struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	IsPublished bool      `json:"isPublished"`
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
		partnerId:     partnerId,
		AggregateRoot: common.NewAggregateRoot(uuid),
	}

	event.AddDomainEvent(events.EventCreatedEvent{
		Id:                 event.id,
		Name:               event.name,
		Description:        event.description,
		Date:               event.date,
		IsPublished:        event.isPublished,
		TotalSpots:         event.totalSpots,
		TotalSpotsReserved: event.totalSpotsReserved,
		PartnerId:          event.partnerId,
	})

	return event, nil
}

type RestoreEventCommand struct {
	Id                 string    `json:"id"`
	Name               string    `json:"name"`
	Description        string    `json:"description"`
	Date               time.Time `json:"date"`
	IsPublished        bool      `json:"isPublished"`
	TotalSpots         int       `json:"totalSpots"`
	PartnerId          string    `json:"partnerId"`
	TotalSpotsReserved int       `json:"totalSpotsReserved"`
	Sections           []*EventSection
}

func RestoreEvent(command RestoreEventCommand) (*Event, error) {
	uuid, err := common.RestoreUUID(command.Id)
	if err != nil {
		return nil, err
	}

	partnerId, err := common.RestoreUUID(command.PartnerId)
	if err != nil {
		return nil, err
	}

	event := &Event{
		id:                 uuid,
		name:               command.Name,
		description:        command.Description,
		date:               command.Date,
		isPublished:        command.IsPublished,
		partnerId:          partnerId,
		AggregateRoot:      common.NewAggregateRoot(uuid),
		totalSpots:         command.TotalSpots,
		totalSpotsReserved: command.TotalSpotsReserved,
		sections:           command.Sections,
	}

	return event, nil
}
