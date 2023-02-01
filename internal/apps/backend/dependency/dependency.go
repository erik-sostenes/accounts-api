package dependency

import (
	"fmt"
	"net/http"

	"github.com/erik-sostenes/accounts-api/internal/apps/backend"
	"github.com/erik-sostenes/accounts-api/internal/apps/backend/controllers"
	"github.com/erik-sostenes/accounts-api/internal/apps/backend/controllers/account"
)

const defaultPort = "9090"

// NewInjector injects all the dependencies to start the app
func NewInjector() error {
	svr := &http.Server{
		Addr: fmt.Sprintf("127.0.0.1:%s", defaultPort),
	}

	h := backend.Server{
		Server: svr,
		Controllers: controllers.Controllers{
			AccountController: account.NewAccountController(),
		},
	}

	return h.Start()
}
