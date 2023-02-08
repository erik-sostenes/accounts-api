package services

import (
	"context"

	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"
)

func (a *accountManager) Create(ctx context.Context, account domain.Account) (err error) {
	return a.Storer.Save(ctx, account.AccountId(), account)
}
