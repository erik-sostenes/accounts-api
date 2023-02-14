package find

import (
	"context"

	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/ports"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/services"
)

// accountCreator implements ports.Storer interface
type accountFinder struct {
	ports.Storer[domain.AccountId, domain.Account]
}

// newAccountFinder returns an instance of the ports.AccountFinder interface by injecting the storer
func NewAccountFinder(storer ports.Storer[domain.AccountId, domain.Account]) ports.AccountFinder[services.AccountResponse] {
	return &accountFinder{
		Storer: storer,
	}
}

// Find method that receives a domain.AccountId, applies the business logic and search the data in it to the storer
func (a accountFinder) Find(ctx context.Context, id domain.AccountId) (services.AccountResponse, error) {
	account, err := a.Search(ctx, id)

	if err != nil {
		return services.AccountResponse{}, err
	}

	return services.AccountResponse{
		AccountId:       account.AccountIP().String(),
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
