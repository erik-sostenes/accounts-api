package domain

import "github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain"

// AccountCareer(Value Object) represent the career of account
type AccountCareer struct {
	value string
}

func NewAccountCareer(value string) (AccountCareer, error) {
	accountCareer, err := domain.String(value).Validate()
	if err != nil {
		return AccountCareer{}, err
	}

	return AccountCareer{
		value: accountCareer,
	}, nil
}

func (a AccountCareer) String() string {
	return a.value
}
