package create

import (
	"context"

	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/ports"
)

// accountCreator implements ports.Store interface
type accountCreator struct {
	ports.Store[domain.AccountId, domain.Account]
}

// NewAccountCreator returns an instance of the ports.AccountCreator interface by injecting the store
func NewAccountCreator(store ports.Store[domain.AccountId, domain.Account]) ports.AccountCreator {
	return &accountCreator{
		Store: store,
	}
}

// Create method that receives a domain.Account, applies the business logic and sends it to the storer
func (a *accountCreator) Create(ctx context.Context, account domain.Account) (err error) {
	return a.Store.Save(ctx, account.AccountId(), account)
}
