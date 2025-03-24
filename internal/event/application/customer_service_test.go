package application_test

import (
	"context"
	"errors"
	"ingressos/internal/common"
	"ingressos/internal/event/application"
	"ingressos/internal/event/domain/entities"
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
			cService := application.NewCustomerService(repositories.NewInMemoryCustomerRepository(uow), common.NewDomainEventManager(), common.RealClock{})
			res, err := cService.Register(context.Background(), application.RegisterCustomerCommand{
				CPF:      "16211571801",
				Name:     "Testando customer",
				Email:    "testando@gmail.com",
				Birthday: time.Date(1974, 4, 7, 0, 0, 0, 0, time.UTC),
			})
			if err != nil {
				t.Fatalf("error %q is not expected", err.Error())
			}

			customerId, err := common.RestoreUUID(res.CustomerId)

			if err != nil {
				t.Fatal("exepcted a valid customer id")
			}
			expectedCustomerCreatedEvent := events.CustomerCreatedEvent{
				ID:       customerId,
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
			cService := application.NewCustomerService(repositories.NewInMemoryCustomerRepository(uow), common.NewDomainEventManager(), common.RealClock{})
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

			if err == nil {
				t.Fatal("error is expected")
			}

			if !errors.Is(err, common.ErrConflict) {
				t.Fatal("expected a conflict err")
			}
		})
	})

	t.Run("Update()", func(t *testing.T) {
		t.Run("should return not found if customer does not exists", func(t *testing.T) {
			uow := &common.UnitOfWork{}
			ctx := context.Background()
			cRepository := repositories.NewInMemoryCustomerRepository(uow)
			cService := application.NewCustomerService(cRepository, common.NewDomainEventManager(), common.FakeClock{
				MockedNow: time.Date(2025, 3, 16, 0, 0, 0, 0, time.UTC),
			})
			id, _ := common.CreateUUID()

			err := cService.Update(ctx, string(id), application.UpdateCustomerCommand{
				Name:      "new name",
				Birthdate: "1975-04-07",
			})

			if err == nil {
				t.Fatalf("error %q is expected", common.ErrNotFound)
			}

			if !errors.Is(err, common.ErrNotFound) {
				t.Fatalf("error %q is expected but received %q", common.ErrNotFound, err.Error())
			}
		})
		t.Run("should update a existing customer", func(t *testing.T) {
			uow := &common.UnitOfWork{}
			ctx := context.Background()
			cRepository := repositories.NewInMemoryCustomerRepository(uow)
			cService := application.NewCustomerService(cRepository, common.NewDomainEventManager(), common.FakeClock{
				MockedNow: time.Date(2025, 3, 16, 0, 0, 0, 0, time.UTC),
			})
			createdRes, err := cService.Register(ctx, application.RegisterCustomerCommand{
				CPF:      "16211571801",
				Name:     "Testando customer",
				Email:    "testando@gmail.com",
				Birthday: time.Date(1974, 4, 7, 0, 0, 0, 0, time.UTC),
			})
			if err != nil {
				t.Fatalf("error %q is not expected", err.Error())
			}
			id := entities.CustomerId(createdRes.CustomerId)
			customerAfterCreate, err := cRepository.GetById(ctx, id)
			if err != nil {
				t.Fatalf("error %q is not expected", err.Error())
			}
			ageBeforeUpdate := customerAfterCreate.GetBirtday().GetAge()
			expectedAge := 50
			if ageBeforeUpdate != expectedAge {
				t.Fatalf("expected age of %d before update but received %d", expectedAge, ageBeforeUpdate)
			}

			err = cService.Update(ctx, string(id), application.UpdateCustomerCommand{
				Name:      "new name",
				Birthdate: "1975-04-07",
			})
			if err != nil {
				t.Fatalf("error %q is not expected", err.Error())
			}
			customerAfterUpdate, err := cRepository.GetById(ctx, id)
			if err != nil {
				t.Fatalf("error %q is not expected", err.Error())
			}
			ageAfterUpdate := customerAfterUpdate.GetBirtday().GetAge()
			expectedAge = 49
			if ageAfterUpdate != expectedAge {
				t.Fatalf("expected age of %d after update but received %d", expectedAge, ageBeforeUpdate)
			}
			if customerAfterUpdate.GetName() != "new name" {
				t.Fatal("should have updated the name")
			}
		})
	})

	t.Run("GetById()", func(t *testing.T) {
		t.Run("should return not found if customer does not exists", func(t *testing.T) {
			uow := &common.UnitOfWork{}
			ctx := context.Background()
			cRepository := repositories.NewInMemoryCustomerRepository(uow)
			cService := application.NewCustomerService(cRepository, common.NewDomainEventManager(), common.FakeClock{
				MockedNow: time.Date(2025, 3, 16, 0, 0, 0, 0, time.UTC),
			})
			id, _ := common.CreateUUID()

			_, err := cService.GetById(ctx, string(id))

			if err == nil {
				t.Fatalf("error %q is expected", common.ErrNotFound)
			}

			if !errors.Is(err, common.ErrNotFound) {
				t.Fatalf("error %q is expected but received %q", common.ErrNotFound, err.Error())
			}
		})
		t.Run("should return an existing customer", func(t *testing.T) {
			uow := &common.UnitOfWork{}
			ctx := context.Background()
			cRepository := repositories.NewInMemoryCustomerRepository(uow)
			cService := application.NewCustomerService(cRepository, common.NewDomainEventManager(), common.FakeClock{
				MockedNow: time.Date(2025, 3, 16, 0, 0, 0, 0, time.UTC),
			})
			createdRes, err := cService.Register(ctx, application.RegisterCustomerCommand{
				CPF:      "16211571801",
				Name:     "Testando customer",
				Email:    "testando@gmail.com",
				Birthday: time.Date(1974, 4, 7, 0, 0, 0, 0, time.UTC),
			})
			if err != nil {
				t.Fatalf("error %q is not expected", err.Error())
			}
			id := entities.CustomerId(createdRes.CustomerId)
			createdCustomer, err := cRepository.GetById(ctx, id)
			if err != nil {
				t.Fatalf("error %q is not expected", err.Error())
			}

			getCustomerResult, err := cService.GetById(ctx, string(id))
			if err != nil {
				t.Fatalf("error %q is not expected", err.Error())
			}

			if getCustomerResult.ID != string(createdCustomer.GetID()) {
				t.Fatalf("expected customer id %q but received %q", getCustomerResult.ID, createdCustomer.GetID())
			}

			if getCustomerResult.Name != string(createdCustomer.GetName()) {
				t.Fatalf("expected customer name %q but received %q", getCustomerResult.ID, createdCustomer.GetName())
			}

			if getCustomerResult.Email != string(createdCustomer.GetEmail()) {
				t.Fatalf("expected customer email %q but received %q", getCustomerResult.ID, createdCustomer.GetEmail())
			}

			if getCustomerResult.CPF != string(createdCustomer.GetCPF()) {
				t.Fatalf("expected customer CPF %q but received %q", getCustomerResult.ID, createdCustomer.GetCPF())
			}

			expectedBDate := "1974-04-07"
			if getCustomerResult.Birthdate != expectedBDate {
				t.Fatalf("expected customer birthdate %q but received %q", getCustomerResult.Birthdate, expectedBDate)
			}
		})
	})
}
