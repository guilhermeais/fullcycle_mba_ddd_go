package controllers

import (
	"errors"
	"ingressos/internal/common"
	"net/http"
)

func HandleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, common.ErrValidation):
		http.Error(w, err.Error(), http.StatusBadRequest)
	case errors.Is(err, common.ErrConflict):
		http.Error(w, err.Error(), http.StatusConflict)
	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
