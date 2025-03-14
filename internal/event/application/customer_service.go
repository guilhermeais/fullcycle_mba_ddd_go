package application

import (
	"context"
	"fmt"
	"ingressos/internal/common"
	"ingressos/internal/event/domain/entities"
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
