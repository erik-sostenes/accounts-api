package domain

import (
	"fmt"
	"regexp"

	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/wrongs"
)

// regexEmail regexp for validating an e-mail
const regexEmail = "^[\\w!#$%&'*+/=?`{|}~^-]+(?:\\.[\\w!#$%&'*+/=?`{|}~^-]+)*@(?:[a-zA-Z0-9-]+\\.)+[a-zA-Z]{2,6}$"

// AccountEmail(Value Object) represent the email of account
type AccountEmail struct {
	value string
}

func NewAccountEmail(value string) (AccountEmail, error) {
	ok, err := regexp.MatchString(regexEmail, value)
	if err != nil {
		return AccountEmail{}, wrongs.StatusInternalServerError(err.Error())
	}

	if !ok {
		return AccountEmail{}, wrongs.StatusUnprocessableEntity(
			fmt.Sprintf("the format %s value is incorrect, it should be an email", value))
	}

	return AccountEmail{
		value: value,
	}, nil
}

func (a *AccountEmail) String() string {
	return a.value
}
