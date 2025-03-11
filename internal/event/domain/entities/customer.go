package entities

import (
	common "ingressos/internal/common"
	"time"
)

type customerId = common.UUID

type customer struct {
	common.AggregateRoot
	id       customerId
	cpf      common.CPF
	name     string
	birthday common.Birthday
}

func (c customer) GetCPF() common.CPF {
	return c.cpf
}

func (c customer) GetID() customerId {
	return c.id
}

func (c customer) GetBirtday() common.Birthday {
	return c.birthday
}

func (c customer) IsEqual(other *customer) bool {
	return c.id.IsEqual(other.id)
}

type CreateCustomerCommand struct {
	CPF      string
	Name     string
	Birthday time.Time
}

func CreateCustomer(c CreateCustomerCommand, clock common.Clock) (*customer, error) {
	cpf, err := common.CreateCPF(c.CPF)
	if err != nil {
		return nil, err
	}

	uuid, err := common.CreateUUID()
	if err != nil {
		return nil, err
	}

	birthday, err := common.CreateBirthday(c.Birthday, clock)
	if err != nil {
		return nil, err
	}

	return &customer{id: uuid, cpf: cpf, name: c.Name, birthday: birthday}, nil
}

type RestoreCustomerCommand struct {
	Id       string
	CPF      string
	Name     string
	Birthday time.Time
}

func RestoreCustomer(c RestoreCustomerCommand, clock common.Clock) (*customer, error) {
	cpf, err := common.CreateCPF(c.CPF)
	if err != nil {
		return nil, err
	}

	uuid, err := common.RestoreUUID(c.Id)
	if err != nil {
		return nil, err
	}

	birthday, err := common.CreateBirthday(c.Birthday, clock)
	if err != nil {
		return nil, err
	}

	return &customer{id: uuid, cpf: cpf, name: c.Name, birthday: birthday}, nil
}
