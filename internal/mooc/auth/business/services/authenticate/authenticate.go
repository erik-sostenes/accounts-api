package authenticate

import (
	"context"
	"fmt"
	account "github.com/erik-sostenes/accounts-api/internal/mooc/account/business/services"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/services/find"
	"github.com/erik-sostenes/accounts-api/internal/mooc/auth/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/mooc/auth/business/domain/wrongs"
	"github.com/erik-sostenes/accounts-api/internal/mooc/auth/business/ports"
	"github.com/erik-sostenes/accounts-api/internal/mooc/auth/business/services"
	"github.com/erik-sostenes/accounts-api/internal/shared/backend/jwt"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/query"
	wrongs2 "github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/wrongs"
)

// accountAuthenticator implements the port.Authenticator interface
var _ ports.Authenticator = (*accountAuthenticator)(nil)

type accountAuthenticator struct {
	query.Bus[find.FindAccountQuery, account.AccountResponse]
	jwt.Token[jwt.Claims]
}

// NewAccountAuthenticator returns an instance of accountAuthenticator with the dependencies that need to authenticate the account
func NewAccountAuthenticator(bus query.Bus[find.FindAccountQuery, account.AccountResponse], token jwt.Token[jwt.Claims]) ports.Authenticator {
	return &accountAuthenticator{
		Bus:   bus,
		Token: token,
	}
}

// Authenticate need to create a FindAccountQuery query to get the account of Data Base and valid the account and credentials math with the
// values that try check the authentication
func (a *accountAuthenticator) Authenticate(ctx context.Context, id domain.AuthId, password domain.AuthPassword) (authResponse services.AuthResponse, err error) {
	accountQuery := find.FindAccountQuery{
		AccountId: id.String(),
	}

	accountResponse, err := a.Bus.Ask(ctx, accountQuery)
	if _, ok := err.(wrongs2.StatusNotFound); ok {
		err = wrongs.InvalidAuthAccount(fmt.Sprintf("The user '%s' does not exists", accountResponse.AccountId))
		return
	}

	authAccount, err := domain.NewAuthAccount(accountResponse.AccountId, accountResponse.AccountPassword)
	if err != nil {
		return
	}

	if a.ensureUserExist(authAccount, id) {
		err = wrongs.InvalidAuthAccount(fmt.Sprintf("The user '%s' does not exists", authAccount.AuthId().String()))
		return
	}

	if a.ensureCredentialsAreValid(authAccount, password) {
		err = wrongs.InvalidAuthCredentials(fmt.Sprintf("The credentials for '%s' are invalid", authAccount.AuthId().String()))
		return
	}

	claims := jwt.NewClaims(accountResponse)

	token, err := a.Token.Generate(claims)

	return services.AuthResponse{
		Token: token,
	}, nil
}

func (a *accountAuthenticator) ensureUserExist(authAccount domain.AuthAccount, id domain.AuthId) bool {
	return authAccount.AuthId().String() != id.String()
}

func (a *accountAuthenticator) ensureCredentialsAreValid(authAccount domain.AuthAccount, password domain.AuthPassword) bool {
	return authAccount.PasswordMatches(password) != nil
}
