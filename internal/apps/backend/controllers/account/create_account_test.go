package account

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/services/create"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/infrastructure/persistence"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/command"
	p2 "github.com/erik-sostenes/accounts-api/internal/shared/mooc/infrastructure/persistence"
)

type funcController func() (AccountController, error)

func TestAccountController_Create(t *testing.T) {
	tsc := map[string]struct {
		*http.Request
		accountController  funcController
		expectedStatusCode int
	}{
		"Given a valid non-existing account, a status code 201 is expected": {
			httptest.NewRequest(http.MethodPut, "/v1/account/create?id=94343721-6baa-4cd5-a0b4-6c5d0419c02d",
				strings.NewReader(`
						{
							"user_name": "JaredNV",
							"name": "Jared Nicolas V", 
							"last_name": "Mitchell",
							"email": "jared.gibson@gmail.com",
							"password": "7or2m27yw6zrkao",
							"career": "ISIC",
							"ip": "192.168.10.0",
							"active": "true",
							"details": {
									"permissions": [1, 2, 3]
							} 
						}`,
				)),
			func() (controller AccountController, err error) {
				store := persistence.NewAccountStore(p2.NewRedisDataBase(p2.NewRedisDBConfiguration()))

				commandHandler := create.CreateAccountCommandHandler{
					AccountCreator: create.NewAccountCreator(store),
				}

				commandBus := make(command.CommandBus[create.CreateAccountCommand])

				if err = commandBus.Record(create.CreateAccountCommand{}, commandHandler); err != nil {
					return
				}

				return NewAccountController(&commandBus), nil
			},
			http.StatusCreated,
		},
		"Given an invalid http request, a status code 400 is expected": {
			httptest.NewRequest(http.MethodPut, "/v1/account/create",
				strings.NewReader(`
						{
							"user_name": "JaredNV",
							"name": "Jared Nicolas V", 
							"last_name": "Mitchell",
							"email": "jared.gibson@gmail.com",
							"password": "7or2m27yw6zrkao",
							"career": "ISIC",
							"ip": "192.168.10.0",
							"active": "true",
							"details": {
									"permissions": [1, 2, 3]
							} 
						}`,
				)),
			func() (controller AccountController, err error) {
				store := persistence.NewAccountStore(p2.NewRedisDataBase(p2.NewRedisDBConfiguration()))

				commandHandler := create.CreateAccountCommandHandler{
					AccountCreator: create.NewAccountCreator(store),
				}

				commandBus := make(command.CommandBus[create.CreateAccountCommand])

				if err = commandBus.Record(create.CreateAccountCommand{}, commandHandler); err != nil {
					return
				}

				return NewAccountController(&commandBus), nil
			},
			http.StatusBadRequest,
		},
	}

	for name, ts := range tsc {
		t.Run(name, func(t *testing.T) {
			req := ts.Request
			resp := httptest.NewRecorder()

			controller, err := ts.accountController()
			if err != nil {
				t.Fatal(err)
			}

			controller.Create(resp, req)

			if ts.expectedStatusCode != resp.Code {
				t.Errorf("status code was expected %d, but it was obtained %d", ts.expectedStatusCode, resp.Code)
			}
		})
	}
}
