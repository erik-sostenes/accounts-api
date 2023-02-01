package account

import (
	"encoding/json"
	"net/http"
	"strings"

	rw "github.com/erik-sostenes/accounts-api/internal/shared/backend/controllers"
)

// Request represents a DTO(Data Transfers Object)
type Request struct {
	Name     string         `json:"name"`
	LastName string         `json:"last_name"`
	Email    string         `json:"email"`
	Password string         `json:"password"`
	Details  map[string]any `json:"details"`
}

// Create method that receives the request body, wraps the DTO and sends it to the service layer
func (c *accountController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		rw.JSON(w, 405, map[string]any{"error": "Method Not Allowed"})
		return
	}

	id := r.URL.Query().Get("id")
	if strings.TrimSpace(id) == "" {
		rw.JSON(w, 400, map[string]any{"error": "Invalid ID Format"})
		return
	}

	var request Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		rw.JSON(w, 422, map[string]any{"error": err})
		return
	}

	rw.JSON(w, 201, nil)
}
