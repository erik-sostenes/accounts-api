package domain

import (
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/wrongs"
	"golang.org/x/crypto/bcrypt"
)

// AccountPassword (Value Object) represent the password of account
type AccountPassword struct {
	value string
}

func NewAccountPassword(value string) (AccountPassword, error) {
	password, err := domain.String(value).Validate()
	if err != nil {
		return AccountPassword{}, err
	}

	accountPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return AccountPassword{}, wrongs.StatusInternalServerError(err.Error())
	}

	return AccountPassword{
		value: string(accountPassword),
	}, nil
}

// NewEncryptedAccountPassword receives an already encrypted string
func NewEncryptedAccountPassword(value string) (string, error) {
	password, err := domain.String(value).Validate()
	if err != nil {
		return value, err
	}

	return password, nil
}

func (a AccountPassword) String() string {
	return a.value
}
