package domain

import "github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain"

// AccountCareer(Value Object) represent the token of account
type AccountToken struct {
	value string
}

func NewAccountToken(value string) (AccountToken, error) {
	accountToken, err := domain.String(value).Validate()
	if err != nil {
		return AccountToken{}, err
	}

	return AccountToken{
		value: accountToken,
	}, nil
}

func (a *AccountToken) String() string {
	return a.value
}
