package controllers

import (
	"context"
	"encoding/json"
	"ingressos/internal/event/application"
	"net/http"
)

type CustomersController struct {
	cService application.CustomerService
}

func NewCustomersController(cService application.CustomerService) *CustomersController {
	return &CustomersController{cService: cService}
}

func (c *CustomersController) Register(w http.ResponseWriter, r *http.Request) {
	var command application.RegisterCustomerCommand

	err := json.NewDecoder(r.Body).Decode(&command)
	if err != nil {
		HandleError(w, err)
		return
	}

	ctx := context.Background()

	result, err := c.cService.Register(ctx, command)
	if err != nil {
		HandleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
