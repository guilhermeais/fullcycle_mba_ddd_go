package entities_test

import (
	"errors"
	"ingressos/internal/common"
	"ingressos/internal/event/domain/entities"
	"testing"
	"time"
)

func TestCreateCustomer(t *testing.T) {
	t.Run("customer with invalid cpf", func(t *testing.T) {
		_, err := entities.CreateCustomer(entities.CreateCustomerCommand{
			CPF:      "00",
			Name:     "Guilherme Teixeira Ais",
			Birthday: time.Date(2003, 8, 26, 0, 0, 0, 0, time.UTC),
		}, common.RealClock{})

		if err == nil {
			t.Fatalf("error is expected when cpf is invalid")
		}
	})

	t.Run("customer with 0 years", func(t *testing.T) {
		_, err := entities.CreateCustomer(entities.CreateCustomerCommand{
			CPF:      "44407433825",
			Name:     "Guilherme Teixeira Ais",
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
		c, err := entities.CreateCustomer(entities.CreateCustomerCommand{
			CPF:      "44407433825",
			Name:     "Guilherme Teixeira Ais",
			Birthday: time.Now().AddDate(-23, 0, 0),
		}, common.RealClock{})

		if err != nil {
			t.Fatalf("error %v is not expected", err)
		}

		if c.GetCPF() != "44407433825" {
			t.Fatalf("expected cpf %s but received %s", "44407433825", c.GetCPF())
		}

		expectedAge := 23
		if c.GetBirtday().GetAge() != 23 {
			t.Fatalf("expected age of %d, but received %d", expectedAge, c.GetBirtday().GetAge())
		}
	})
}

func TestIsEqual(t *testing.T) {
	firstCustomer, _ := entities.CreateCustomer(entities.CreateCustomerCommand{
		CPF:      "44407433825",
		Name:     "Guilherme Teixeira Ais",
		Birthday: time.Date(2003, 8, 26, 0, 0, 0, 0, time.UTC),
	}, common.RealClock{})

	secondCustomer, _ := entities.RestoreCustomer(entities.RestoreCustomerCommand{
		Id:       string(firstCustomer.GetID()),
		CPF:      "44407433825",
		Name:     "Guilherme Teixeira Ais",
		Birthday: time.Date(2003, 8, 26, 0, 0, 0, 0, time.UTC),
	}, common.RealClock{})

	if !firstCustomer.IsEqual(secondCustomer) {
		t.Fatalf("customer %s shoud be equal to %s", firstCustomer.GetID(), secondCustomer.GetID())
	}
}
