package application_test

import (
	"context"
	"errors"
	"ingressos/internal/common"
	"ingressos/internal/event/application"
	"ingressos/internal/event/domain/events"
	"ingressos/internal/event/infra/repositories"
	"reflect"
	"testing"
	"time"
)

func TestCustomerService(t *testing.T) {
	t.Run("Register()", func(t *testing.T) {
		t.Run("should create a customer", func(t *testing.T) {
			uow := &common.UnitOfWork{}
			cService := application.NewCustomerService(repositories.NewInMemoryCustomerRepository(uow), common.RealClock{})
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
			expectedCustomerCreatedEvent := events.CustomerCreatedEvent{
				ID:       res.CustomerId,
				Name:     "Testando customer",
				Email:    "testando@gmail.com",
				CPF:      "16211571801",
				Birthday: "1974-04-07",
			}
			for _, aggregate := range uow.GetAggregateRoots() {
				if len(aggregate.GetDomainEvents()) == 0 {
					t.Fatalf("should've registered domain event of %T", expectedCustomerCreatedEvent)
				}

				receivedCustomerCreatedEvent := aggregate.GetDomainEvents()[0]
				if !reflect.DeepEqual(receivedCustomerCreatedEvent, expectedCustomerCreatedEvent) {
					t.Fatalf("customer created event should be %v, but received %v", expectedCustomerCreatedEvent, receivedCustomerCreatedEvent)
				}
			}
		})

		t.Run("should not create a customer with same CPF", func(t *testing.T) {
			uow := &common.UnitOfWork{}
			cService := application.NewCustomerService(repositories.NewInMemoryCustomerRepository(uow), common.RealClock{})
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
