package dependency

import (
	"fmt"
	"net/http"

	"github.com/erik-sostenes/accounts-api/internal/apps/backend"
	"github.com/erik-sostenes/accounts-api/internal/apps/backend/controllers"
	"github.com/erik-sostenes/accounts-api/internal/apps/backend/controllers/account"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/services/create"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/infrastructure/persistence"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/command"
	p2 "github.com/erik-sostenes/accounts-api/internal/shared/mooc/infrastructure/persistence"
)

const defaultPort = "9090"

// NewInjector injects all the dependencies to start the app
func NewInjector() error {
	svr := &http.Server{
		Addr: fmt.Sprintf("127.0.0.1:%s", defaultPort),
	}

	accountController, err := injectsAccountHandlerDependencies()
	if err != nil {
		return err
	}

	h := backend.Server{
		Server: svr,
		Controllers: controllers.Controllers{
			AccountController: accountController,
		},
	}

	return h.Start()
}

// injectsAccountHandlerDependencies prepares all dependencies to be able to create an account
// returns an account.AccountController instance with the command already entered
func injectsAccountHandlerDependencies() (controller account.AccountController, err error) {
	storer := persistence.NewAccountStorer(p2.NewRedisDataBase(p2.NewRedisDBConfiguration()))

	commandHandler := create.CreateAccountCommandHandler{
		AccountCreator: create.NewAccountCreator(storer),
	}

	commandBus := make(command.CommandBus[create.CreateAccountCommand])

	if err = commandBus.Record(create.CreateAccountCommand{}, commandHandler); err != nil {
		return
	}

	return account.NewAccountController(&commandBus), nil
}
