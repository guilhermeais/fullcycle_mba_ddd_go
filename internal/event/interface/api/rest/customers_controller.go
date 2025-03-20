package rest

import (
	"encoding/json"
	"ingressos/internal/event/application"
	"net/http"
)

type CustomersController struct {
	cService *application.CustomerService
}

func NewCustomersController(mux *http.ServeMux, cService *application.CustomerService) *CustomersController {
	c := &CustomersController{cService: cService}
	mux.HandleFunc("POST /customers/register", c.Register)
	mux.HandleFunc("GET /customers/{id}", c.GetById)

	return c
}

func (c *CustomersController) Register(w http.ResponseWriter, r *http.Request) {
	var command application.RegisterCustomerCommand

	err := json.NewDecoder(r.Body).Decode(&command)
	if err != nil {
		HandleError(w, err)
		return
	}

	ctx := r.Context()

	result, err := c.cService.Register(ctx, command)
	if err != nil {
		HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (c *CustomersController) GetById(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	if idString == "" {
		writeJSONError(w, http.StatusBadRequest, "customer id should be provided")
	}
	ctx := r.Context()

	result, err := c.cService.GetById(ctx, idString)
	if err != nil {
		HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
