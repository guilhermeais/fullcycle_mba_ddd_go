package entities

import (
	"context"
	"ingressos/internal/common"
)

type customerRepository interface {
	Save(ctx context.Context, c *Customer) error
	GetById(ctx context.Context, id CustomerId) (*Customer, error)
	ExistsByCPF(ctx context.Context, cpf common.CPF) (bool, error)
}

type CustomerRepository struct {
	customerRepoImpl customerRepository
	uow              *common.UnitOfWork
}

func NewCustomerRepository(impl customerRepository, uow *common.UnitOfWork) *CustomerRepository {
	return &CustomerRepository{
		customerRepoImpl: impl,
		uow:              uow,
	}
}

func (cr *CustomerRepository) Save(ctx context.Context, c *Customer) error {
	err := cr.customerRepoImpl.Save(ctx, c)
	if err != nil {
		return err
	}

	cr.uow.RegisterAggregate(c.AggregateRoot)

	return nil
}

func (cr *CustomerRepository) GetById(ctx context.Context, id CustomerId) (*Customer, error) {
	return cr.customerRepoImpl.GetById(ctx, id)
}

func (cr *CustomerRepository) ExistsByCPF(ctx context.Context, cpf common.CPF) (bool, error) {
	return cr.customerRepoImpl.ExistsByCPF(ctx, cpf)
}
