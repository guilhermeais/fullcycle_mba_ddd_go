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

func (m *InMemoryCustomerRepository) Save(c *entities.Customer, _ context.Context) error {
	if c == nil {
		return errors.New("customer is nil")
	}

	m.customers[string(c.GetID())] = c
	return nil
}

func (m *InMemoryCustomerRepository) GetById(id entities.CustomerId, _ context.Context) (*entities.Customer, error) {
	customer, exists := m.customers[string(id)]
	if !exists {
		return nil, errors.New("customer not found")
	}
	return customer, nil
}

func (m *InMemoryCustomerRepository) ExistsByCPF(cpf common.CPF, _ context.Context) (bool, error) {
	for _, c := range m.customers {
		if c.GetCPF().IsEqual(cpf) {
			return true, nil
		}
	}
	return false, nil
}
