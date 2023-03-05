package domain

import "github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain"

// AuthId (Value Object) represents auth id account
type AuthId struct {
	value string
}

// NewAuthID returns an instance of AuthID
func NewAuthID(value string) (AuthId, error) {
	id, err := domain.Identifier(value).ParseUuID()
	if err != nil {
		return AuthId{}, err
	}

	return AuthId{
		value: id,
	}, nil
}

func (id AuthId) String() string {
	return id.value
}
