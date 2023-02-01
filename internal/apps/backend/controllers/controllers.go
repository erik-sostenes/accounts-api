package controllers

import "github.com/erik-sostenes/accounts-api/internal/apps/backend/controllers/account"

// Controllers contains all the http controllers to handle http requests
type Controllers struct {
	AccountController account.AccountController
}
