package entities

import (
	"context"
	"ingressos/internal/common"
)

type CustomerRepository interface {
	Save(ctx context.Context, c *Customer) error
	GetById(ctx context.Context, id CustomerId) (*Customer, error)
	ExistsByCPF(ctx context.Context, cpf common.CPF) (bool, error)
}
