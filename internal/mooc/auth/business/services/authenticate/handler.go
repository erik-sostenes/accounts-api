package authenticate

import (
	"context"

	"github.com/erik-sostenes/accounts-api/internal/mooc/auth/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/mooc/auth/business/ports"
	"github.com/erik-sostenes/accounts-api/internal/mooc/auth/business/services"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/query"
)

// AuthenticateAccountQuery implements the command.Command interface
var _ query.Query = AuthenticateAccountQuery{}

// AuthenticateAccountQuery represents the DTO with values primitives
type AuthenticateAccountQuery struct {
	Id       string
	Password string
}

func (AuthenticateAccountQuery) QueryId() string {
	return "authenticate_account_query"
}

// AuthenticateAccountQueryHandler implements the command.Handler interface
var _ query.Handler[AuthenticateAccountQuery, services.AuthResponse] = (*AuthenticateAccountQueryHandler)(nil)

type AuthenticateAccountQueryHandler struct {
	ports.Authenticator
}

// Handler instance a domain.AuthAccount (Domain Object) with the command primitives values
// after domain.AuthAccount is sent t Authenticator port
func (h *AuthenticateAccountQueryHandler) Handler(ctx context.Context, qry AuthenticateAccountQuery) (services.AuthResponse, error) {
	id, err := domain.NewAuthID(qry.Id)
	if err != nil {
		return services.AuthResponse{}, err
	}

	password, err := domain.NewAuthPassword(qry.Password)
	if err != nil {
		return services.AuthResponse{}, err
	}

	return h.Authenticator.Authenticate(ctx, id, password)
}
