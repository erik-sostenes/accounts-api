package dependency

import (
	"fmt"
	"net/http"

	"github.com/erik-sostenes/accounts-api/internal/apps/backend"
	"github.com/erik-sostenes/accounts-api/internal/apps/backend/controllers"
	"github.com/erik-sostenes/accounts-api/internal/apps/backend/controllers/account"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/services"
	p2 "github.com/erik-sostenes/accounts-api/internal/shared/mooc/infrastructure/persistence"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/infrastructure/persistence"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/command"
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

func injectsAccountHandlerDependencies() (controller account.AccountController, err error) {
	storer := persistence.NewAccountStorer(p2.NewRedisDataBase(p2.NewRedisDBConfiguration()))

	commandHandler := services.CreateAccountCommandHandler{
		AccountManager: services.NewAccountManager(storer),
	}

	commandBus := make(command.CommandBus[services.CreateAccountCommand])

	if err = commandBus.Record(services.CreateAccountCommand{}, commandHandler); err != nil {
		return
	}

	return account.NewAccountController(&commandBus), nil 
}
