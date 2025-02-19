package entities

import (
	common "ingressos/internal/common"
	"time"
)

type customerId = common.UUID

type customer struct {
	id       customerId
	cpf      common.CPF
	name     string
	birthday time.Time
}

func (c customer) GetCPF() common.CPF {
	return c.cpf
}

func (c customer) GetID() customerId {
	return c.id
}

func (c customer) IsEqual(other *customer) bool {
	return c.id.IsEqual(other.id)
}

type CreateCustomerCommand struct {
	CPF      string
	Name     string
	Birthday time.Time
}

func CreateCustomer(c CreateCustomerCommand) (*customer, error) {
	cpf, err := common.CreateCPF(c.CPF)
	if err != nil {
		return nil, err
	}

	uuid, err := common.CreateUUID()
	if err != nil {
		return nil, err
	}

	return &customer{id: uuid, cpf: cpf, name: c.Name, birthday: c.Birthday}, nil
}

type RestoreCustomerCommand struct {
	Id       string
	CPF      string
	Name     string
	Birthday time.Time
}

func RestoreCustomer(c RestoreCustomerCommand) (*customer, error) {
	cpf, err := common.CreateCPF(c.CPF)
	if err != nil {
		return nil, err
	}

	uuid, err := common.RestoreUUID(c.Id)
	if err != nil {
		return nil, err
	}
	return &customer{id: uuid, cpf: cpf, name: c.Name, birthday: c.Birthday}, nil
}
