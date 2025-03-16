package entities

import (
	common "ingressos/internal/common"
	"ingressos/internal/event/domain/events"
	"time"
)

type CustomerId = common.UUID

type Customer struct {
	common.AggregateRoot
	id            CustomerId
	cpf           common.CPF
	name          string
	email         common.Email
	birthday      common.Birthday
	clockProvider common.Clock
}

func (c Customer) GetCPF() common.CPF {
	return c.cpf
}

func (c Customer) GetID() CustomerId {
	return c.id
}

func (c Customer) GetBirtday() common.Birthday {
	return c.birthday
}

func (c *Customer) UpdateBirthdate(newBirthdate time.Time) error {
	parsedBirthday, err := common.CreateBirthday(newBirthdate, c.clockProvider)
	if err != nil {
		return err
	}
	c.birthday = parsedBirthday
	c.AddDomainEvent(events.CustomerUpdatedEvent{
		ID:       string(c.id),
		Name:     c.name,
		Birthday: c.birthday.Format(common.BirthdateLayout)},
	)
	return nil
}

func (c Customer) GetEmail() common.Email {
	return c.email
}

func (c Customer) GetName() string {
	return c.name
}

func (c *Customer) UpdateName(newName string) {
	c.name = newName
	c.AddDomainEvent(events.CustomerUpdatedEvent{
		ID:       string(c.id),
		Name:     c.name,
		Birthday: c.birthday.Format(common.BirthdateLayout)},
	)
}

func (c Customer) IsEqual(other *Customer) bool {
	return c.id.IsEqual(other.id)
}

type CreateCustomerCommand struct {
	CPF      string    `json:"cpf"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Birthday time.Time `json:"birthday"`
}

func CreateCustomer(c CreateCustomerCommand, clock common.Clock) (*Customer, error) {
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

	email, err := common.CreateEmail(c.Email)
	if err != nil {
		return nil, err
	}

	customer := &Customer{id: uuid, cpf: cpf, name: c.Name, birthday: birthday, email: email, AggregateRoot: common.NewAggregateRoot(uuid), clockProvider: clock}
	customer.AddDomainEvent(events.CustomerCreatedEvent{
		ID:       string(customer.id),
		Name:     customer.name,
		Email:    string(customer.email),
		CPF:      string(customer.cpf),
		Birthday: customer.birthday.Format(common.BirthdateLayout)},
	)
	return customer, nil
}

type RestoreCustomerCommand struct {
	Id       string
	CPF      string
	Name     string
	Birthday time.Time
	Email    string
}

func RestoreCustomer(c RestoreCustomerCommand, clock common.Clock) (*Customer, error) {
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

	email, err := common.CreateEmail(c.Email)
	if err != nil {
		return nil, err
	}

	return &Customer{id: uuid, cpf: cpf, name: c.Name, birthday: birthday, email: email}, nil
}
