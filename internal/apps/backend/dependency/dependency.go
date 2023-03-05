package dependency

import (
	"fmt"
	"github.com/erik-sostenes/accounts-api/internal/apps/backend/controllers/auth"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/ports"
	account2 "github.com/erik-sostenes/accounts-api/internal/mooc/account/business/services"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/services/find"
	"github.com/erik-sostenes/accounts-api/internal/mooc/auth/business/services"
	"github.com/erik-sostenes/accounts-api/internal/mooc/auth/business/services/authenticate"
	"github.com/erik-sostenes/accounts-api/internal/shared/backend/jwt"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/query"
	"net/http"

	"github.com/erik-sostenes/accounts-api/internal/apps/backend"
	"github.com/erik-sostenes/accounts-api/internal/apps/backend/controllers"
	"github.com/erik-sostenes/accounts-api/internal/apps/backend/controllers/account"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/services/create"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/infrastructure/persistence"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/command"
	p2 "github.com/erik-sostenes/accounts-api/internal/shared/mooc/infrastructure/persistence"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAxe8efqXwuRZ4V1FoCmUkRln6loUfp8XtqdRWzJciVivkJZld
/jHONSqfTw4SpY0qQkD5G+uNYb3MSEOkBH9UCwCHoSdl1sjYDaWbx/Je4d/NM6YG
E3rsWgwvIcujGCLz3BQ2hO+57NqEAHQI3LkWXI+Rie4K5HQAVvMCat4UMe4CL++Z
52PC5yh1sgAbLVnoPBoLI7cYCERPhuexWm+gnjez1GKrWkvBACB2t+XAC+hU1TTc
Ry6aEHUvJ5+dlBglfPWT6UQJw7PJRaSELD8YLih4C0saxpq+M+ZMoriv1/phY/kk
g1ZafT+lydgBlMVxLnESqgm20zJqplsIg2X7EQIDAQABAoIBABtkooXIlW4oK/N5
srptkP2jiki2l9DyVZgBaRnbeMcQP/zsItQBNJarFW0td2suBEEzGMbCbMiwKct+
gP6WWJ1FL4AgIbn+BditqMedRYBhJtcVDRY5FujHcuZsdl/qxnEY4wq22rZq74XY
iTly7CNXQz8hkKRZYYqnCxibL5RRK6Z5Zo9ohTNFOZIw897k4s9FZEOmEG8pWho9
TNE9tgHBuuYQzaJjLgWqIM97TatDNv7KA996tFNNTmtPzQl1B4MFO5YdpL+OjJ2L
IQAzs8vD3fbCEpYb7UTWHCU8d9bpWpONByO8ZOo49wUg/bZIHxZwb5uQtg63E5z3
+gdmiFkCgYEA+A1m2tAMfTHScNSKy+pir9xHEjR8JdUABWeYmbtxES6Iu8uTy8Vy
i40J6YfgRtyzUqEeo18MKfJWzk11RBPrSO8uNFFpQPDoL3FcknFxbBlhmngpcAUG
vqYNCrVLyUNXI1vDXnFZW+eg6E9CKf1pJ1R9BPc/qALRDj/lzFxDW+sCgYEAzEah
5NVSD83+GlUd+t+Xtyf9UqkNyKzPLUm4On3mXVLmx35J4tJpihm5SHFUo9gxY3dn
JxU4vbxqFd6P8xCvWKEQShqrrlCHZ5YUnbHdTxEd+atKf64IMIfwX87P8Jb5kY1i
6gHbVP3i5X/HRLqz6prMpnOc6J0Xg+nw8eHpcfMCgYAliPmgcM0DANAETNU36B7I
179VbOXAX8viBXwc/zUr0WvVZwfVVOpxXYU7dlkkv+7OuRzGwfI4QriJ/USaaZ03
6yGFvy/7KLkpvLCyZEIyhmCznC1BCzGrFbtxfF+cc/kym4cjumk4NAOwQ5YSfosz
7WABqVxTkyGJU3f1hZyXwwKBgFA/x0Xwj8ZptFN/8MEnqaBoc1pP03xsdw9hkKBZ
6W/sK4FfmYMkChYYuPM+onOjcPOUas+txJa1OC/TOVXRzjDRRWb3R065kBgfm4W/
5CM1pEL7Cc9S/SCjpsjcpE/t36lQk/U+OX4QJ1zlb9EOT7PwkEkrzg6L+Dr4YpGD
oIQFAoGBANkKowkWZwTQjj+Rcsr2Tl99BWQouaX92hMYwHSLlV+IxA+gVCz3d5o8
XX6jW7U9o45KANn9TVbBV+TLR2iJyZWjUPgNSVRoPnlPJfFYFbIYbS9mnZ+x9JDf
Tkqn9t2JRnRp7OXIpdXPxe+glrrYZA365rb4aZ048RTT3omcREG4
-----END RSA PRIVATE KEY-----`

const publicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxe8efqXwuRZ4V1FoCmUk
Rln6loUfp8XtqdRWzJciVivkJZld/jHONSqfTw4SpY0qQkD5G+uNYb3MSEOkBH9U
CwCHoSdl1sjYDaWbx/Je4d/NM6YGE3rsWgwvIcujGCLz3BQ2hO+57NqEAHQI3LkW
XI+Rie4K5HQAVvMCat4UMe4CL++Z52PC5yh1sgAbLVnoPBoLI7cYCERPhuexWm+g
njez1GKrWkvBACB2t+XAC+hU1TTcRy6aEHUvJ5+dlBglfPWT6UQJw7PJRaSELD8Y
Lih4C0saxpq+M+ZMoriv1/phY/kkg1ZafT+lydgBlMVxLnESqgm20zJqplsIg2X7
EQIDAQAB
-----END PUBLIC KEY-----`

const defaultPort = "9090"

// NewInjector injects all the dependencies to start the app
func NewInjector() error {
	svr := &http.Server{
		Addr: fmt.Sprintf("127.0.0.1:%s", defaultPort),
	}

	store := persistence.NewAccountStore(p2.NewRedisDataBase(p2.NewRedisDBConfiguration()))

	jwt := NewJWT()
	accountController, err := injectsAccountCreatorHandlerDependencies(store)
	if err != nil {
		return err
	}

	accountFinder, err := injectsAccountFinderHandlerDependencies(store)
	if err != nil {
		return err
	}

	authController, err := injectsAuthenticateAccount(&accountFinder, jwt)
	if err != nil {
		return err
	}

	h := backend.Server{
		Server: svr,
		Token:  jwt,
		Controllers: controllers.Controllers{
			AccountController: accountController,
			AuthController:    authController,
		},
	}

	return h.Start()
}

func NewJWT() jwt.Token[jwt.Claims] {
	return jwt.NewJWT([]byte(privateKey), []byte(publicKey))
}

func injectsAuthenticateAccount(bus query.Bus[find.FindAccountQuery, account2.AccountResponse], jwt jwt.Token[jwt.Claims]) (controller auth.AuthController, err error) {
	authenticator := authenticate.NewAccountAuthenticator(bus, jwt)

	queryHandler := authenticate.AuthenticateAccountQueryHandler{
		Authenticator: authenticator,
	}

	queryBus := make(query.QueryBus[authenticate.AuthenticateAccountQuery, services.AuthResponse])

	if err := queryBus.Record(authenticate.AuthenticateAccountQuery{}, &queryHandler); err != nil {
		return controller, err
	}

	return auth.NewAuthController(&queryBus), nil
}

func injectsAccountFinderHandlerDependencies(store ports.Store[domain.AccountId, domain.Account]) (bus query.QueryBus[find.FindAccountQuery, account2.AccountResponse], err error) {
	queryHandler := find.FindAccountQueryHandler{
		AccountFinder: find.NewAccountFinder(store),
	}

	bus = make(query.QueryBus[find.FindAccountQuery, account2.AccountResponse])

	if err = bus.Record(find.FindAccountQuery{}, &queryHandler); err != nil {
		return bus, err
	}

	return bus, nil
}

// injectsAccountHandlerDependencies prepares all dependencies to be able to create an account
// returns an account.AccountController instance with the command already entered
func injectsAccountCreatorHandlerDependencies(store ports.Store[domain.AccountId, domain.Account]) (controller account.AccountController, err error) {
	commandHandler := create.CreateAccountCommandHandler{
		AccountCreator: create.NewAccountCreator(store),
	}

	commandBus := make(command.CommandBus[create.CreateAccountCommand])

	if err = commandBus.Record(create.CreateAccountCommand{}, commandHandler); err != nil {
		return
	}

	return account.NewAccountController(&commandBus), nil
}
