package persistence

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/ports"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/wrongs"
	"github.com/go-redis/redis/v8"
)

// accountKey method that generates a string key to save or verify the user's account
// example -> registeredAccount:94343721-6baa-4cd5-a0b4-6c5d0419c02d
func (a *accountStore) accountKey(id string) string {
	return "registeredAccount:" + id
}

// accountStore implements ports.Store interface
type accountStore struct {
	*redis.Client
}

// NewAccountStore returns an instance of the ports.Strore interface with the generic data initialized
//
// injects the database type to be used to persist the data
func NewAccountStore(rdb *redis.Client) ports.Store[domain.AccountId, domain.Account] {
	return &accountStore{
		Client: rdb,
	}
}

// Save method that persists in redis a user's account in Redis
func (a *accountStore) Save(ctx context.Context, id domain.AccountId, account domain.Account) error {
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

// Search method that searches in redis for a user account by an identifier
func (a *accountStore) Search(ctx context.Context, accountId domain.AccountId) (account domain.Account, err error) {
	values, err := a.HGetAll(ctx, a.accountKey(accountId.String())).Result()
	if err != nil {
		return
	}

	if len(values) != 10 {
		return
	}

	var details domain.Map
	if err = json.Unmarshal([]byte(values["details"]), &details); err != nil {
		return
	}

	password, err := domain.NewEncryptedAccountPassword(values["password"])
	if err != nil {
		return
	}

	if account, err = domain.NewAccount(
		values["id"],
		values["username"],
		values["name"],
		values["last_name"],
		values["email"],
		password,
		values["career"],
		values["ip"],
		values["active"],
		details,
	); err != nil {
		return
	}

	return
}

// Remove method that removes a redis account by means of a key created by the account ID
func (a *accountStore) Remove(ctx context.Context, id domain.AccountId) (err error) {
	if _, err = a.Del(ctx, a.accountKey(id.String())).Result(); err != nil {
		err = wrongs.StatusInternalServerError(err.Error())
		return
	}
	return
}
