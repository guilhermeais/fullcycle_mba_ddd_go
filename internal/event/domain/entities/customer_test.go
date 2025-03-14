package entities_test

import (
	"errors"
	"ingressos/internal/common"
	"ingressos/internal/event/domain/entities"
	"ingressos/internal/event/domain/events"
	"testing"
	"time"
)

func TestCustomer(t *testing.T) {
	t.Run("CreateCustomer()", func(t *testing.T) {
		t.Run("customer with invalid cpf", func(t *testing.T) {
			_, err := entities.CreateCustomer(entities.CreateCustomerCommand{
				CPF:      "00",
				Name:     "Guilherme Teixeira Ais",
				Email:    "guilhermeteixeiraais@gmail.com",
				Birthday: time.Date(2003, 8, 26, 0, 0, 0, 0, time.UTC),
			}, common.RealClock{})

			if err == nil {
				t.Fatalf("error is expected when cpf is invalid")
			}
		})

		t.Run("customer with invalid email", func(t *testing.T) {
			var invalidEmails = [...]string{"guilhermeteixeiraais", "gui@com"}
			for _, invalidEmail := range invalidEmails {
				_, err := entities.CreateCustomer(entities.CreateCustomerCommand{
					CPF:      "44407433825",
					Name:     "Guilherme Teixeira Ais",
					Email:    invalidEmail,
					Birthday: time.Date(2003, 8, 26, 0, 0, 0, 0, time.UTC),
				}, common.RealClock{})

				if err == nil {
					t.Fatalf("error is expected when email is invalid")
				}

				if !errors.Is(err, common.ErrValidation) {
					t.Fatal("should return a validation error when email is invalid")
				}

				expectedError := common.MakeErrEmailValidation(invalidEmail)
				if err.Error() != expectedError.Error() {
					t.Fatalf("expect error %q but received %q", expectedError.Error(), err.Error())
				}
			}
		})

		t.Run("customer with 0 years", func(t *testing.T) {
			_, err := entities.CreateCustomer(entities.CreateCustomerCommand{
				CPF:      "44407433825",
				Name:     "Guilherme Teixeira Ais",
				Email:    "guilhermeteixeiraais@gmail.com",
				Birthday: time.Now(),
			}, common.RealClock{})

			if err == nil {
				t.Fatalf("error is expected when birthday is invalid")
			}
			if !errors.Is(err, common.ErrValidation) {
				t.Fatalf("expected a validation err but received %v", err)
			}

			expectedErr := common.MakeErrInvalidBirthday(0)
			if expectedErr.Error() != err.Error() {
				t.Fatalf("expected error: %v but received error: %v", expectedErr, err)
			}
		})

		t.Run("customer with -1 years", func(t *testing.T) {
			_, err := entities.CreateCustomer(entities.CreateCustomerCommand{
				CPF:      "44407433825",
				Name:     "Guilherme Teixeira Ais",
				Email:    "guilhermeteixeiraais@gmail.com",
				Birthday: time.Now().AddDate(1, 0, 0),
			}, common.RealClock{})

			if err == nil {
				t.Fatalf("error is expected when birthday is invalid")
			}

			expectedErr := common.MakeErrInvalidBirthday(-1)
			if expectedErr.Error() != err.Error() && errors.Is(err, common.ErrValidation) {
				t.Fatalf("expected error: %v but received error: %v", expectedErr, err)
			}
		})

		t.Run("valid customer", func(t *testing.T) {
			birthday := time.Date(2003, 8, 26, 0, 0, 0, 0, time.UTC)
			now := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
			expectedAge := 21

			c, err := entities.CreateCustomer(entities.CreateCustomerCommand{
				CPF:      "44407433825",
				Name:     "Guilherme Teixeira Ais",
				Birthday: birthday,
				Email:    "guilhermeteixeiraais@gmail.com",
			}, common.FakeClock{MockedNow: now})

			if err != nil {
				t.Fatalf("error %v is not expected", err)
			}

			if c.GetCPF() != "44407433825" {
				t.Fatalf("expected cpf %s but received %s", "44407433825", c.GetCPF())
			}

			if c.GetBirtday().GetAge() != expectedAge {
				t.Fatalf("expected age of %d, but received %d", expectedAge, c.GetBirtday().GetAge())
			}

			domainEvents := c.GetDomainEvents()
			if len(domainEvents) <= 0 {
				t.Fatalf("should've emitted %T domain event", events.CustomerCreatedEvent{})
			}
			customerCreatedEvent, ok := domainEvents[0].(events.CustomerCreatedEvent)
			if !ok {
				t.Fatalf("expected event type %T but received %T", events.CustomerCreatedEvent{}, domainEvents[0])
			}

			if customerCreatedEvent.ID != string(c.GetID()) {
				t.Fatalf("expected customer ID %s but received %s", c.GetID(), customerCreatedEvent.ID)
			}

			if customerCreatedEvent.CPF != string(c.GetCPF()) {
				t.Fatalf("expected customer CPF %s but received %s", c.GetCPF(), customerCreatedEvent.CPF)
			}

			if customerCreatedEvent.Email != string(c.GetEmail()) {
				t.Fatalf("expected customer CPF %s but received %s", c.GetEmail(), customerCreatedEvent.Email)
			}

			if customerCreatedEvent.Name != string(c.GetName()) {
				t.Fatalf("expected customer CPF %s but received %s", c.GetName(), customerCreatedEvent.Name)
			}

			expectedBirtday := "2003-08-26"
			if customerCreatedEvent.Birthday != expectedBirtday {
				t.Fatalf("expected birthdate of %q but received %q", expectedBirtday, customerCreatedEvent.Birthday)
			}
		})
	})
	t.Run("IsEqual()", func(t *testing.T) {
		firstCustomer, _ := entities.CreateCustomer(entities.CreateCustomerCommand{
			CPF:      "44407433825",
			Name:     "Guilherme Teixeira Ais",
			Email:    "guilhermeteixeiraais@gmail.com",
			Birthday: time.Date(2003, 8, 26, 0, 0, 0, 0, time.UTC),
		}, common.RealClock{})

		secondCustomer, _ := entities.RestoreCustomer(entities.RestoreCustomerCommand{
			Id:       string(firstCustomer.GetID()),
			CPF:      "44407433825",
			Name:     "Guilherme Teixeira Ais",
			Email:    "guilhermeteixeiraais@gmail.com",
			Birthday: time.Date(2003, 8, 26, 0, 0, 0, 0, time.UTC),
		}, common.RealClock{})

		if !firstCustomer.IsEqual(secondCustomer) {
			t.Fatalf("customer %s shoud be equal to %s", firstCustomer.GetID(), secondCustomer.GetID())
		}
	})
}
