package persistence

import (
	"context"
	"fmt"

	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/ports"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/wrongs"
	"github.com/go-redis/redis/v8"
)

// accountKey method that generates a string key to save or verify the user's account
// example -> registeredAccount:94343721-6baa-4cd5-a0b4-6c5d0419c02d
func (a *accountStorer) accountKey(id string) string {
	return "registeredAccount:" + id
}

// accountStorer implements ports.Storer interface
type accountStorer struct {
	*redis.Client
}

// NewAccountStorer returns an instance of the ports.Strorer interface with the generic data initialized
//
// injects the database type to be used to persist the data
func NewAccountStorer(rdb *redis.Client) ports.Storer[domain.AccountId, domain.Account] {
	return &accountStorer{
		Client: rdb,
	}
}

// Save method that persists a user's account in Redis
func (a *accountStorer) Save(ctx context.Context, id domain.AccountId, account domain.Account) error {
	ok, err := a.HExists(ctx, a.accountKey(id.String()), "id").Result()
	if err != nil {
		err = wrongs.StatusInternalServerError(err.Error())
		return err
	}

	if ok {
		err = wrongs.StatusBadRequest(fmt.Sprintf("Account with id %s already exists", id.String()))
		return err
	}

	if _, err := a.HSet(ctx, a.accountKey(id.String()), map[string]any{
		"id":        id.String(),
		"username":  account.AccountUserName().String(),
		"name":      account.AccountName().String(),
		"last_name": account.AccountLastName().String(),
		"email":     account.AccountEmail().String(),
		"password":  account.AccountPassword().String(),
		"career":    account.AccountCareer().String(),
		"ip":        account.AccountIP().String(),
		"active":    account.AccountActive().Bool(),
		"details":   domain.NewMarshalJSON(account.AccountDetails().Value()),
	}).Result(); err != nil {
		err = wrongs.StatusInternalServerError(err.Error())
		return err
	}

	return nil
}

// Removes method that removes a redis account by means of a key created by the account ID
func (a accountStorer) Remove(ctx context.Context, id domain.AccountId) (err error) {
	if _, err = a.Del(ctx, a.accountKey(id.String())).Result(); err != nil {
		err = wrongs.StatusInternalServerError(err.Error())
		return
	}
	return
}
