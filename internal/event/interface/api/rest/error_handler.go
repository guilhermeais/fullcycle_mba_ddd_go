package rest

import (
	"encoding/json"
	"errors"
	"ingressos/internal/common"
	"net/http"
)

type JSONError struct {
	Error string `json:"error"`
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(JSONError{message})
}

func HandleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, common.ErrValidation):
		writeJSONError(w, http.StatusBadRequest, err.Error())
	case errors.Is(err, common.ErrConflict):
		writeJSONError(w, http.StatusConflict, err.Error())
	case errors.Is(err, common.ErrNotFound):
		writeJSONError(w, http.StatusNotFound, err.Error())
	default:
		writeJSONError(w, http.StatusInternalServerError, err.Error())
	}
}
