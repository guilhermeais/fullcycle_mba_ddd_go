package rest_test

import (
	"bytes"
	"encoding/json"
	"ingressos/internal/common"
	"ingressos/internal/event/application"
	"ingressos/internal/event/infra/repositories"
	"ingressos/internal/event/interface/api/rest"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var mockedNow = time.Date(2025, 3, 18, 10, 0, 0, 0, time.UTC)

func setupTestServer() *httptest.Server {
	mux := http.NewServeMux()
	uow := &common.UnitOfWork{}
	domainEventManager := common.NewDomainEventManager()
	customerRepo := repositories.NewInMemoryCustomerRepository(uow)
	customerService := application.NewCustomerService(customerRepo, domainEventManager, common.FakeClock{MockedNow: mockedNow})
	rest.NewCustomersController(mux, customerService)
	return httptest.NewServer(mux)
}
func TestCustomersController(t *testing.T) {
	t.Run("POST /customers/register", func(t *testing.T) {
		t.Run("should register a valid customer", func(t *testing.T) {
			server := setupTestServer()
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
}
