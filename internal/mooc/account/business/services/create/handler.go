package create

import (
	"context"

	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/ports"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/command"
)

// CreateAccountCommand implements the command.Command interface
var _ command.Command = CreateAccountCommand{}

// CreateAccountCommand represents the DTO with the primitive values
type CreateAccountCommand struct {
	AccountId       string
	AccountUserName string
	AccountName     string
	AccountLastName string
	AccountEmail    string
	AccountPassword string
	AccountCareer   string
	AccountIP       string
	AccountActive   string
	AccountDetails  domain.Map
}

func (CreateAccountCommand) CommandId() string {
	return "create_account_command"
}

// CreateAccountCommandHandler implements the command.Handler interface
var _ command.Handler[CreateAccountCommand] = &CreateAccountCommandHandler{}

type CreateAccountCommandHandler struct {
	ports.AccountCreator
}

// Handler instantiates a domain.Account (Domain Object) with the command primitive values
func (h CreateAccountCommandHandler) Handler(ctx context.Context, cmd CreateAccountCommand) (err error) {
	var account domain.Account

	if account, err = domain.NewAccount(
		cmd.AccountId,
		cmd.AccountUserName,
		cmd.AccountName,
		cmd.AccountLastName,
		cmd.AccountEmail,
		cmd.AccountPassword,
		cmd.AccountCareer,
		cmd.AccountIP,
		cmd.AccountActive,
		cmd.AccountDetails,
	); err != nil {
		return
	}

	return h.AccountCreator.Create(ctx, account)
}
