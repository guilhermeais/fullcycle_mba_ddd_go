package entities

import (
	common "ingressos/internal/common"
)

type partnerId = common.UUID

type partner struct {
	id   partnerId
	name string
}

func (c partner) IsEqual(other *partner) bool {
	return c.id.IsEqual(other.id)
}

type CreatePartnerCommand struct {
	Name string
}

func CreatePartner(c CreatePartnerCommand) (*partner, error) {
	uuid, err := common.CreateUUID()
	if err != nil {
		return nil, err
	}

	return &partner{id: uuid, name: c.Name}, nil
}

type RestorePartnerCommand struct {
	Id   string
	Name string
}

func RestorePartner(c RestorePartnerCommand) (*partner, error) {
	uuid, err := common.RestoreUUID(c.Id)
	if err != nil {
		return nil, err
	}
	return &partner{id: uuid, name: c.Name}, nil
}
