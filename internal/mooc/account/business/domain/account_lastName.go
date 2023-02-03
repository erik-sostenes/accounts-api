package domain

import "github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain"

// AccountLastName(Value Object) represent the last name of account
type AccountLastName struct {
	value string
}

func NewAccountLastName(value string) (AccountLastName, error) {
	accountLastName, err := domain.String(value).Validate()
	if err != nil {
		return AccountLastName{}, err
	}

	return AccountLastName{
		value: accountLastName,
	}, nil
}

func (a *AccountLastName) String() string {
	return a.value
}
