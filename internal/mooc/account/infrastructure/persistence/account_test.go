package persistence

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/ports"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/wrongs"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/infrastructure/persistence"
)

func TestStorer_Save(t *testing.T) {
	tsc := map[string]struct {
		ports.Storer[domain.AccountId, domain.Account]
		domain.AccountFunc
		expectedError error
	}{
		"Given a valid non-existing account, it will be persisted": {
			NewAccountStorer(persistence.NewRedisDataBase(persistence.NewRedisDBConfiguration())),
			domain.NewFirstAccount(),
			nil,
		},
		"Given an existing valid account, it will not be persisted": {
			NewAccountStorer(persistence.NewRedisDataBase(persistence.NewRedisDBConfiguration())),
			domain.NewTwoAccount(),
			func() error {
				account, _ := domain.NewTwoAccount().Account()

				return wrongs.StatusBadRequest(
					fmt.Sprintf("Account with id %s already exists", account.AccountId().String()))
			}(),
		},
	}

	storer := NewAccountStorer(persistence.NewRedisDataBase(persistence.NewRedisDBConfiguration()))
	account, err := domain.NewTwoAccount().Account()
	if err != nil {
		t.Fatal(err)
	}

	// SetUp prepare configuration before running integration tests all
	if err := func() error {
		return storer.Save(context.Background(), account.AccountId(), account)
	}(); err != nil {
		t.Fatal(err)
	}

	// Teardown reset configuration after running integration tests all
	t.Cleanup(func() {
		_ = storer.Remove(context.TODO(), account.AccountId())
	})

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			account, err := ts.AccountFunc.Account()
			if err != nil {
				t.Error(err)
				t.SkipNow()
			}

			t.Cleanup(func() {
				_ = ts.Remove(context.TODO(), account.AccountId())
			})

			err = ts.Save(context.TODO(), account.AccountId(), account)
			if !errors.Is(err, ts.expectedError) {
				t.Errorf("%v error was expected, but %v error was obtained", ts.expectedError, err)
			}
		})
	}
}
