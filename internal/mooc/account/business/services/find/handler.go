package find

import (
	"context"

	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/ports"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/services"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/query"
)

// FindAccountQuery implements the command.Query interface
var _ query.Query = FindAccountQuery{}

// FindAccountQuery represents the DTO with the primitive values
//
// represents the request that want performed
type FindAccountQuery struct {
	AccountId string
}

func (FindAccountQuery) QueryId() string {
	return "find_account_query"
}

// FindAccountQueryHandler implements the query.Handler interface
var _ query.Handler[FindAccountQuery, services.AccountResponse] = &FindAccountQueryHandler{}

type FindAccountQueryHandler struct {
	ports.AccountFinder[services.AccountResponse]
}

// Handler instantiates a domain.AccountId (Domain Object) with the query primitive value
func (h *FindAccountQueryHandler) Handler(ctx context.Context, qry FindAccountQuery) (services.AccountResponse, error) {
	id, err := domain.NewAccountId(qry.AccountId)

	if err != nil {
		return services.AccountResponse{}, err
	}

	return h.AccountFinder.Find(ctx, id)
}
