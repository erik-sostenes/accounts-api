package create

import (
	"context"

	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/ports"
)

var _ ports.Storer[domain.AccountId, domain.Account] = accountCreator{}

// accountCreator implements ports.Storer interface
type accountCreator struct {
	ports.Storer[domain.AccountId, domain.Account]
}

// NewAccountCreator returns an instance of the ports.AccountCreator interface by injecting the storer
func NewAccountCreator(storer ports.Storer[domain.AccountId, domain.Account]) ports.AccountCreator {
	return &accountCreator{
		Storer: storer,
	}
}

// Create method that receives a domain.account, applies the business logic and sends it to the storer
func (a *accountCreator) Create(ctx context.Context, account domain.Account) (err error) {
	return a.Storer.Save(ctx, account.AccountId(), account)
}
