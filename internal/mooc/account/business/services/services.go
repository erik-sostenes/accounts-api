package services

import (
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/ports"
)

var _ ports.AccountManager = &accountManager{}

type accountManager struct {
	ports.Storer[domain.AccountId, domain.Account]
}

func NewAccountManager(storer ports.Storer[domain.AccountId, domain.Account]) ports.AccountManager {
	return &accountManager{
		Storer: storer,
	}
}
