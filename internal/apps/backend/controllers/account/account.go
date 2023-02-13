package account

import (
	"net/http"

	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/services/create"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/command"
)

// AccountController contains the contracts for managing user accounts via http requests
type AccountController interface {
	// Create http handler that receives an http request to create an account
	Create(http.ResponseWriter, *http.Request)
}

// accountController implements the AccountController interface
type accountController struct {
	command.Bus[create.CreateAccountCommand]
}

// NewAccountController injects all dependencies to create the AccountController instance
func NewAccountController(bus command.Bus[create.CreateAccountCommand]) AccountController {
	return &accountController{
		Bus: bus,
	}
}
