package domain

import "github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain"

// AccountName(Value Object) represent the name of account
type AccountName struct {
	value string
}

func NewAccountName(value string) (AccountName, error) {
	accountName, err := domain.String(value).Validate()
	if err != nil {
		return AccountName{}, err
	}

	return AccountName{
		value: accountName,
	}, nil
}

func (a *AccountName) String() string {
	return a.value
}
