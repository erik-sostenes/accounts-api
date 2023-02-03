package domain

import "github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain"

// AccountActive(Value Object) represents whether the account is active.
type AccountActive struct {
	value bool
}

func NewAccounActive(value string) (AccountActive, error) {
	accountActive, err := domain.Bool(value).Validate()
	if err != nil {
		return AccountActive{}, err
	}

	return AccountActive{
		value: accountActive,
	}, nil
}

func (a *AccountActive) Bool() bool {
	return a.value
}
