package controllers

import (
	"encoding/json"
	"net/http"
)

// JSON method that sets up an http response by setting the Content-Type, statusCode
// and encoding the response data in the standard JSON format
func JSON(w http.ResponseWriter, statusCode int, data any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(data)
}
