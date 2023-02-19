package find

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/services"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/infrastructure/persistence"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/query"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/wrongs"
)

type (
	queryBusFunc func() (query.QueryBus[FindAccountQuery, services.AccountResponse], error)
)

var account = func() domain.Account {
	account, err := domain.NewFirstAccount().Account()
	if err != nil {
		panic(err)
	}

	return account
}()

// id represents an identifier no registered in the storer
var id string = "94343721-6baa-4cd5-a0b4-6c5d0419c02d"

func TestAccountHandler_Find(t *testing.T) {
	tsc := map[string]struct {
		// id represents the primitive value with which the FindAccountQuery query will be created
		id string
		queryBusFunc
		expectedError error
	}{
		"Given a valid account that does not exist, an error of type wrongs.NotFound is expected": {
			id: id,
			queryBusFunc: func() (bus query.QueryBus[FindAccountQuery, services.AccountResponse], err error) {
				queryHandler := FindAccountQueryHandler{
					NewAccountFinder(persistence.NewMockStorer[domain.AccountId, domain.Account]()),
				}

				bus = make(query.QueryBus[FindAccountQuery, services.AccountResponse])

				if err = bus.Record(FindAccountQuery{}, &queryHandler); err != nil {
					return
				}

				return
			},
			expectedError: wrongs.StatusNotFound(fmt.Sprintf("resource with id %v not found", id)),
		},
		"Given a valid account that exists, no errors are expected.": {
			id: account.AccountId().String(),
			queryBusFunc: func() (bus query.QueryBus[FindAccountQuery, services.AccountResponse], err error) {
				mockStorer := persistence.NewMockStorer[domain.AccountId, domain.Account]()
				err = mockStorer.Save(context.TODO(), account.AccountId(), account)
				if err != nil {
					return
				}

				queryHandler := FindAccountQueryHandler{
					NewAccountFinder(mockStorer),
				}

				bus = make(query.QueryBus[FindAccountQuery, services.AccountResponse])

				if err = bus.Record(FindAccountQuery{}, &queryHandler); err != nil {
					return
				}
				return
			},
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			bus, err := ts.queryBusFunc()
			if err != nil {
				t.Fatal(err)
			}

			query := FindAccountQuery{
				AccountId: ts.id,
			}

			_, err = bus.Ask(context.Background(), query)
			if !errors.Is(err, ts.expectedError) {
				t.Errorf("error was expected %v, but error it was obtained %v", ts.expectedError, err)
			}
		})
	}
}
