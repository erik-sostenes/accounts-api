package auth

import (
	"net/http"
	"strings"

	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/mooc/auth/business/services/authenticate"
	rw "github.com/erik-sostenes/accounts-api/internal/shared/backend/controllers"
)

func (c *authController) Authenticate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		_ = rw.ResponseJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "Method Not Allowed"})
		return
	}

	id := r.URL.Query().Get("id")
	if strings.TrimSpace(id) == "" {
		_ = rw.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Missing query parameter id"})
		return
	}

	password := r.URL.Query().Get("password")
	if strings.TrimSpace(id) == "" {
		_ = rw.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Missing query parameter password"})
		return
	}

	query := authenticate.AuthenticateAccountQuery{
		Id:       id,
		Password: password,
	}

	authResponse, err := c.Bus.Ask(r.Context(), query)
	if err != nil {
		_ = rw.HandlerError(w, err)
		return
	}

	_ = rw.ResponseJSON(w, http.StatusOK, domain.Map{"response": authResponse})
}
