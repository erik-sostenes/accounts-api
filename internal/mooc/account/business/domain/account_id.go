package domain

import "github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain"

// AccountId(Value Object) represent the id of account
type AccountId struct {
	value string
}

func NewAccountId(value string) (AccountId, error) {
	id, err := domain.Identifier(value).ParseUuID()
	if err != nil {
		return AccountId{}, err
	}

	return AccountId{
		value: id,
	}, nil
}

func (id *AccountId) String() string {
	return id.value
}
