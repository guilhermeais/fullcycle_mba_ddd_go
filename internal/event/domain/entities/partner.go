package entities

import (
	common "ingressos/internal/common"
	"time"
)

type PartnerId = common.UUID

type Partner struct {
	common.AggregateRoot
	id   PartnerId
	name string
}

func (c Partner) IsEqual(other *Partner) bool {
	return c.id.IsEqual(other.id)
}

type InitEventCommand struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	TotalSpots  int       `json:"totalSpots"`
}

func (p Partner) InitEvent(c InitEventCommand) (*Event, error) {
	event, err := CreateEvent(CreateEventCommand{
		Name:        c.Name,
		Description: c.Description,
		Date:        c.Date,
		IsPublished: false,
		TotalSpots:  c.TotalSpots,
		PartnerId:   string(p.id),
	})

	return event, err
}

type CreatePartnerCommand struct {
	Name string
}

func CreatePartner(c CreatePartnerCommand) (*Partner, error) {
	uuid, err := common.CreateUUID()
	if err != nil {
		return nil, err
	}

	return &Partner{id: uuid, name: c.Name, AggregateRoot: common.NewAggregateRoot(uuid)}, nil
}

type RestorePartnerCommand struct {
	Id   string
	Name string
}

func RestorePartner(c RestorePartnerCommand) (*Partner, error) {
	uuid, err := common.RestoreUUID(c.Id)
	if err != nil {
		return nil, err
	}
	return &Partner{id: uuid, name: c.Name, AggregateRoot: common.NewAggregateRoot(uuid)}, nil
}
