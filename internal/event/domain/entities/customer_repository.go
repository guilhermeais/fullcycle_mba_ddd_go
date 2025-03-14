package entities

import (
	"context"
	"ingressos/internal/common"
)

type CustomerRepository interface {
	Save(c *Customer, ctx context.Context) error
	GetById(id CustomerId, ctx context.Context) (*Customer, error)
	ExistsByCPF(cpf common.CPF, ctx context.Context) (bool, error)
}
