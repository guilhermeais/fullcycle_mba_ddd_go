package rest_test

import (
	"bytes"
	"context"
	"encoding/json"
	"ingressos/internal/common"
	"ingressos/internal/event/application"
	"ingressos/internal/event/domain/entities"
	"ingressos/internal/event/infra/repositories"
	"ingressos/internal/event/interface/api/rest"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var mockedNow = time.Date(2025, 3, 18, 10, 0, 0, 0, time.UTC)

func setupTestServer() (*httptest.Server, *entities.CustomerRepository, common.Clock) {
	mux := http.NewServeMux()
	uow := &common.UnitOfWork{}
	domainEventManager := common.NewDomainEventManager()
	customerRepo := repositories.NewInMemoryCustomerRepository(uow)
	clockProvider := &common.FakeClock{MockedNow: mockedNow}
	customerService := application.NewCustomerService(customerRepo, domainEventManager, clockProvider)
	rest.NewCustomersController(mux, customerService)
	return httptest.NewServer(mux), customerRepo, clockProvider
}
func TestCustomersController(t *testing.T) {
	t.Run("POST /customers/register", func(t *testing.T) {
		t.Run("should register a valid customer", func(t *testing.T) {
			server, _, _ := setupTestServer()
			defer server.Close()

			command := application.RegisterCustomerCommand{
				CPF:      "16211571801",
				Name:     "Testando customer",
				Email:    "testando@gmail.com",
				Birthday: time.Date(1974, 4, 7, 0, 0, 0, 0, time.UTC),
			}

			body, err := json.Marshal(command)
			if err != nil {
				t.Fatalf("failed to marshal command: %v", err)
			}

			resp, err := http.Post(server.URL+"/customers/register", "application/json", bytes.NewBuffer(body))
			if err != nil {
				t.Fatalf("failed to send request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				t.Errorf("expected status OK, got %v", resp.Status)
			}

			var result application.RegisterCustomerResult
			err = json.NewDecoder(resp.Body).Decode(&result)
			if err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if result.CustomerId == "" {
				t.Errorf("expected non-empty CustomerId")
			}
		})
	})

	t.Run("GET /customers/{id}", func(t *testing.T) {
		t.Run("should return 404 if customer does not exists", func(t *testing.T) {
			server, _, _ := setupTestServer()
			defer server.Close()

			nonExistingId, _ := common.CreateUUID()
			resp, err := http.Get(server.URL + "/customers/" + string(nonExistingId))
			if err != nil {
				t.Fatalf("failed to send request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusNotFound {
				t.Errorf("expected status NOT FOUND, got %v", resp.Status)
			}

			var result rest.JSONError
			err = json.NewDecoder(resp.Body).Decode(&result)
			if err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			const expectedErrorMsg = "recurso não encontrado: Cliente não encontrado"
			if result.Error != expectedErrorMsg {
				t.Errorf("expected message %q", expectedErrorMsg)
			}
		})

		t.Run("should return 200 with the customer when exists", func(t *testing.T) {
			server, cRepo, clock := setupTestServer()
			defer server.Close()
			ctx := context.Background()
			c, _ := entities.CreateCustomer(entities.CreateCustomerCommand{
				CPF:      "44407433825",
				Name:     "Guilherme Teixeira Ais",
				Email:    "guilhermeteixeiraais@gmail.com",
				Birthday: time.Date(2003, 8, 26, 0, 0, 0, 0, time.UTC),
			}, clock)
			cRepo.Save(ctx, c)

			resp, err := http.Get(server.URL + "/customers/" + string(c.GetID()))
			if err != nil {
				t.Fatalf("failed to send request: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				t.Errorf("expected status OK, got %v", resp.Status)
			}

			var result application.GetCustomerByIdResult
			err = json.NewDecoder(resp.Body).Decode(&result)
			if err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if result.ID != string(c.GetID()) {
				t.Fatalf("expected customer id of %q but received %q", c.GetID(), result.ID)
			}
			if result.Name != string(c.GetName()) {
				t.Fatalf("expected customer name of %q but received %q", c.GetName(), result.Name)
			}
			if result.Email != string(c.GetEmail()) {
				t.Fatalf("expected customer email of %q but received %q", c.GetEmail(), result.Email)
			}
			if result.CPF != string(c.GetCPF()) {
				t.Fatalf("expected customer CPF of %q but received %q", c.GetCPF(), result.CPF)
			}
			const expectedBDay = "2003-08-26"
			if result.Birthdate != expectedBDay {
				t.Fatalf("expected customer birthdate of %q but received %q", expectedBDay, result.Birthdate)
			}
		})
	})
}
