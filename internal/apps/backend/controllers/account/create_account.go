package account

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/services/create"
	rw "github.com/erik-sostenes/accounts-api/internal/shared/backend/controllers"
)

// Request represents a DTO(Data Transfers Object)
type Request struct {
	UserName string     `json:"user_name"`
	Name     string     `json:"name"`
	LastName string     `json:"last_name"`
	Email    string     `json:"email"`
	Password string     `json:"password"`
	Career   string     `json:"career"`
	Ip       string     `json:"ip"`
	Active   string     `json:"active"`
	Details  domain.Map `json:"details"`
}

// Create method that receives the request body, wraps the DTO and sends it to the service layer
func (c *accountController) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		rw.JSON(w, http.StatusMethodNotAllowed, domain.Map{"error": "Method Not Allowed"})
		return
	}

	id := r.URL.Query().Get("id")
	if strings.TrimSpace(id) == "" {
		rw.JSON(w, http.StatusBadRequest, domain.Map{"error": "Invalid ID Format"})
		return
	}

	var request Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		rw.JSON(w, http.StatusUnprocessableEntity, domain.Map{"error": err})
		return
	}

	command := create.CreateAccountCommand{
		AccountId:       id,
		AccountUserName: request.UserName,
		AccountName:     request.Name,
		AccountLastName: request.LastName,
		AccountEmail:    request.Email,
		AccountPassword: request.Password,
		AccountCareer:   request.Career,
		AccountIP:       request.Ip,
		AccountActive:   request.Active,
		AccountDetails:  request.Details,
	}

	if err := c.Dispatch(r.Context(), command); err != nil {
		rw.HandlerError(w, err)
		return
	}

	rw.JSON(w, http.StatusCreated, nil)
}
