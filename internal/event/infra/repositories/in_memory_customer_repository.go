package repositories

import (
	"context"
	"errors"
	"ingressos/internal/common"
	"ingressos/internal/event/domain/entities"
)

type InMemoryCustomerRepository struct {
	customers map[string]*entities.Customer
}

func NewInMemoryCustomerRepository() *InMemoryCustomerRepository {
	return &InMemoryCustomerRepository{
		customers: make(map[string]*entities.Customer),
	}
}

func (m *InMemoryCustomerRepository) Save(ctx context.Context, c *entities.Customer) error {
	if c == nil {
		return errors.New("customer is nil")
	}

	m.customers[string(c.GetID())] = c
	return nil
}

func (m *InMemoryCustomerRepository) GetById(ctx context.Context, id entities.CustomerId) (*entities.Customer, error) {
	customer, exists := m.customers[string(id)]
	if !exists {
		return nil, errors.New("customer not found")
	}
	return customer, nil
}

func (m *InMemoryCustomerRepository) ExistsByCPF(ctx context.Context, cpf common.CPF) (bool, error) {
	for _, c := range m.customers {
		if c.GetCPF().IsEqual(cpf) {
			return true, nil
		}
	}
	return false, nil
}
