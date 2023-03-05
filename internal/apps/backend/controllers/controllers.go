package controllers

import (
	"github.com/erik-sostenes/accounts-api/internal/apps/backend/controllers/account"
	"github.com/erik-sostenes/accounts-api/internal/apps/backend/controllers/auth"
)

// Controllers contains all the http controllers to handle http requests
type Controllers struct {
	AccountController account.AccountController
	AuthController    auth.AuthController
}
