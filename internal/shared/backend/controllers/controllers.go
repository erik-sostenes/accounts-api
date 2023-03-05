package controllers

import (
	"encoding/json"
	wrongs2 "github.com/erik-sostenes/accounts-api/internal/mooc/auth/business/domain/wrongs"
	"net/http"

	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/wrongs"
)

// ResponseJSON method that sets up a http response by setting the Content-Type, statusCode
// and encoding the response data in the standard ResponseJSON format
func ResponseJSON(w http.ResponseWriter, statusCode int, data any) error {
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
		return ResponseJSON(w, http.StatusBadRequest, domain.Map{"error": err})
	case wrongs.StatusUnprocessableEntity:
		return ResponseJSON(w, http.StatusUnprocessableEntity, domain.Map{"error": err})
	case wrongs.StatusForbidden:
		return ResponseJSON(w, http.StatusForbidden, domain.Map{"error": err})
	case wrongs.StatusNotFound:
		return ResponseJSON(w, http.StatusNotFound, domain.Map{"error": err})
	case wrongs2.InvalidAuthAccount:
		return ResponseJSON(w, http.StatusForbidden, domain.Map{"error": err})
	case wrongs2.InvalidAuthCredentials:
		return ResponseJSON(w, http.StatusUnauthorized, domain.Map{"error": err})
	default:
		return ResponseJSON(w, http.StatusInternalServerError, domain.Map{"error": err})
	}
}
