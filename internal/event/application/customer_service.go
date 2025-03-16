package application

import (
	"context"
	"fmt"
	"ingressos/internal/common"
	"ingressos/internal/event/domain/entities"
	"time"
)

type DomainEventManager interface {
	Dispatch(domainEvent []any) error
}
type CustomerService struct {
	repository    *entities.CustomerRepository
	clockProvider common.Clock
}

func NewCustomerService(repository *entities.CustomerRepository, clock common.Clock) CustomerService {
	return CustomerService{repository, clock}
}

type RegisterCustomerCommand = entities.CreateCustomerCommand
type RegisterCustomerResult struct {
	CustomerId string `json:"customerId"`
}

func (cs *CustomerService) Register(ctx context.Context, command RegisterCustomerCommand) (RegisterCustomerResult, error) {
	c, err := entities.CreateCustomer(command, cs.clockProvider)
	if err != nil {
		return RegisterCustomerResult{}, err
	}

	exists, err := cs.repository.ExistsByCPF(ctx, c.GetCPF())
	if err != nil {
		return RegisterCustomerResult{}, err
	}

	if exists {
		return RegisterCustomerResult{}, fmt.Errorf("%w: cliente com o CPF %q já existe", common.ErrConflict, c.GetCPF())
	}

	err = cs.repository.Save(ctx, c)
	if err != nil {
		return RegisterCustomerResult{}, err
	}

	return RegisterCustomerResult{CustomerId: string(c.GetID())}, nil
}

type UpdateCustomerCommand struct {
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
}

func (cs *CustomerService) Update(ctx context.Context, id entities.CustomerId, command UpdateCustomerCommand) error {
	c, err := cs.repository.GetById(ctx, id)
	if err != nil {
		return err
	}

	if command.Name != "" {
		c.UpdateName(command.Name)
	}

	if command.Birthdate != "" {
		parsedBirthdate, err := time.Parse(common.BirthdateLayout, command.Birthdate)

		if err != nil {
			return ErrInvalidDate
		}
		c.UpdateBirthdate(parsedBirthdate)
	}

	err = cs.repository.Save(ctx, c)
	if err != nil {
		return err
	}

	return err
}

func MakeErrCPFInUse(cpf common.CPF) error {
	return fmt.Errorf("%w: cliente com o CPF %q já existe", common.ErrConflict, cpf)
}

var ErrInvalidDate = fmt.Errorf("%w: data com fortmado inválido. Exemplo de formato válido (%q)", common.ErrValidation, common.BirthdateLayout)
