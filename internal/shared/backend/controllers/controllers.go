package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/wrongs"
)

// JSON method that sets up an http response by setting the Content-Type, statusCode
// and encoding the response data in the standard JSON format
func JSON(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(data)
}

// HandlerError handles the type of error
//
// responds with http status code  along with the message
func HandlerError(w http.ResponseWriter, err error) error {
	switch err.(type) {
	case wrongs.StatusBadRequest:
		return JSON(w, http.StatusBadRequest, domain.Map{"error": err})
	case wrongs.StatusUnprocessableEntity:
		return JSON(w, http.StatusUnprocessableEntity, domain.Map{"error": err})
	case wrongs.StatusForbidden:
		return JSON(w, http.StatusForbidden, domain.Map{"error": err})
	case wrongs.StatusNotFound:
		return JSON(w, http.StatusNotFound, domain.Map{"error": err})
	default:
		return JSON(w, http.StatusInternalServerError, domain.Map{"error": err})
	}
}
