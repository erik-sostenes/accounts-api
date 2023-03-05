package auth

import (
	"net/http"

	"github.com/erik-sostenes/accounts-api/internal/mooc/auth/business/services"
	"github.com/erik-sostenes/accounts-api/internal/mooc/auth/business/services/authenticate"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/query"
)

// AccountController ..
type AuthController interface {
	Authenticate(http.ResponseWriter, *http.Request)
}

// authController ...
type authController struct {
	query.Bus[authenticate.AuthenticateAccountQuery, services.AuthResponse]
}

// NewAccountController
func NewAuthController(bus query.Bus[authenticate.AuthenticateAccountQuery, services.AuthResponse]) AuthController {
	return &authController{
		Bus: bus,
	}
}
