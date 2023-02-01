package account

import "net/http"

// AccountController contains the contracts for managing user accounts via http requests
type AccountController interface {
	// Create http handler that receives an http request to create an account
	Create(http.ResponseWriter, *http.Request)
}

// accountController implements the AccountController interface
type accountController struct{}

// NewAccountController injects all dependencies to create the AccountController instance
func NewAccountController() AccountController {
	return &accountController{}
}
