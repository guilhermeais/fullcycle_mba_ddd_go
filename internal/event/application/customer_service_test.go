package application_test

import (
	"context"
	"errors"
	"ingressos/internal/common"
	"ingressos/internal/event/application"
	"ingressos/internal/event/infra/repositories"
	"testing"
	"time"
)

func TestCustomerService(t *testing.T) {
	t.Run("Register()", func(t *testing.T) {
		t.Run("should create a customer", func(t *testing.T) {
			cService := application.NewCustomerSerivce(repositories.NewInMemoryCustomerRepository(), common.RealClock{})
			res, err := cService.Register(context.Background(), application.RegisterCustomerCommand{
				CPF:      "16211571801",
				Name:     "Testando customer",
				Email:    "testando@gmail.com",
				Birthday: time.Date(1974, 4, 7, 0, 0, 0, 0, time.UTC),
			})
			if err != nil {
				t.Fatalf("error %q is not expected", err.Error())
			}

			if !common.ValidateUUID(res.CustomerId) {
				t.Fatal("exepcted a valid customer id")
			}
		})

		t.Run("should not create a customer with same CPF", func(t *testing.T) {
			cService := application.NewCustomerSerivce(repositories.NewInMemoryCustomerRepository(), common.RealClock{})
			cService.Register(context.Background(), application.RegisterCustomerCommand{
				CPF:      "16211571801",
				Name:     "Testando customer",
				Email:    "testando@gmail.com",
				Birthday: time.Date(1974, 4, 7, 0, 0, 0, 0, time.UTC),
			})

			_, err := cService.Register(context.Background(), application.RegisterCustomerCommand{
				CPF:      "16211571801",
				Name:     "Testando customer",
				Email:    "testando@gmail.com",
				Birthday: time.Date(1974, 4, 7, 0, 0, 0, 0, time.UTC),
			})

			if err != nil {
				t.Fatal("error is expected")
			}

			if !errors.Is(err, common.ErrConflict) {
				t.Fatal("expected a conflict err")
			}
		})
	})
}
