package entities

import "ingressos/internal/common"

type EventSpotId = common.UUID

type EventSpot struct {
	id          EventSpotId
	location    string
	isReserved  bool
	isPublished bool
}

func (es *EventSpot) Publish() {
	es.isPublished = true
}

func (es *EventSpot) IsPublished() bool {
	return es.isPublished
}

func (es *EventSpot) GetID() EventSpotId {
	return es.id
}

type CreateEventSpotCommand struct {
	Location                string
	IsReserved, IsPublished bool
}

func CreateEventSpot(c CreateEventSpotCommand) (*EventSpot, error) {
	uuid, err := common.CreateUUID()
	if err != nil {
		return nil, err
	}

	return &EventSpot{
		id:          uuid,
		location:    c.Location,
		isReserved:  c.IsReserved,
		isPublished: c.IsPublished,
	}, nil
}
