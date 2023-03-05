package find

import (
	"context"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/ports"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/services"
)

// accountCreator implements ports.Store interface
type accountFinder struct {
	ports.Store[domain.AccountId, domain.Account]
}

// NewAccountFinder returns an instance of the ports.AccountFinder interface by injecting the store
func NewAccountFinder(store ports.Store[domain.AccountId, domain.Account]) ports.AccountFinder[services.AccountResponse] {
	return &accountFinder{
		Store: store,
	}
}

// Find method that receives a domain.AccountId, applies the business logic and search the data in it to the store
func (a accountFinder) Find(ctx context.Context, id domain.AccountId) (services.AccountResponse, error) {
	account, err := a.Search(ctx, id)
	if err != nil {
		return services.AccountResponse{}, err
	}

	return services.AccountResponse{
		AccountId:       account.AccountId().String(),
		AccountUserName: account.AccountUserName().String(),
		AccountName:     account.AccountName().String(),
		AccountLastName: account.AccountLastName().String(),
		AccountEmail:    account.AccountEmail().String(),
		AccountPassword: account.AccountPassword().String(),
		AccountCareer:   account.AccountCareer().String(),
		AccountIP:       account.AccountIP().String(),
		AccountActive:   account.AccountActive().Bool(),
		AccountDetails:  account.AccountDetails().Value(),
	}, nil
}
