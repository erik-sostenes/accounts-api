package domain

import "github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain"

// AccountUserName(Value Object) represent the user name of account
type AccountUserName struct {
	value string
}

func NewAccountUserName(value string) (AccountUserName, error) {
	accountUserName, err := domain.String(value).Validate()
	if err != nil {
		return AccountUserName{}, err
	}

	return AccountUserName{
		value: accountUserName,
	}, nil
}

func (a AccountUserName) String() string {
	return a.value
}
