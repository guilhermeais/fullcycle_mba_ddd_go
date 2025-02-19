package entities_test

import (
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
		})

		if err == nil {
			t.Fatalf("error is expected when cpf is invalid")
		}
	})

	t.Run("valid customer", func(t *testing.T) {
		c, err := entities.CreateCustomer(entities.CreateCustomerCommand{
			CPF:      "44407433825",
			Name:     "Guilherme Teixeira Ais",
			Birthday: time.Date(2003, 8, 26, 0, 0, 0, 0, time.UTC),
		})

		if err != nil {
			t.Fatalf("error %v is not expected", err)
		}

		if c.GetCPF() != "44407433825" {
			t.Fatalf("expected cpf %s but received %s", "44407433825", c.GetCPF())
		}
	})
}

func TestIsEqual(t *testing.T) {
	firstCustomer, _ := entities.CreateCustomer(entities.CreateCustomerCommand{
		CPF:      "44407433825",
		Name:     "Guilherme Teixeira Ais",
		Birthday: time.Date(2003, 8, 26, 0, 0, 0, 0, time.UTC),
	})

	secondCustomer, _ := entities.RestoreCustomer(entities.RestoreCustomerCommand{
		Id:       string(firstCustomer.GetID()),
		CPF:      "44407433825",
		Name:     "Guilherme Teixeira Ais",
		Birthday: time.Date(2003, 8, 26, 0, 0, 0, 0, time.UTC),
	})

	if !firstCustomer.IsEqual(secondCustomer) {
		t.Fatalf("customer %s shoud be equal to %s", firstCustomer.GetID(), secondCustomer.GetID())
	}
}
