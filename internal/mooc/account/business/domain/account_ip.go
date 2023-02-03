package domain

import (
	"fmt"
	"regexp"

	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/wrongs"
)

// regexIP regexp for validating an ipv4
const regexpIP = "(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])"

// AccountIP(Value Object) represent the ip of account
type AccountIP struct {
	value string
}

func NewAccountIP(value string) (AccountIP, error) {
	ok, err := regexp.MatchString(regexpIP, value)
	if err != nil {
		return AccountIP{}, wrongs.StatusInternalServerError(err.Error())
	}

	if !ok {
		return AccountIP{}, wrongs.StatusUnprocessableEntity(
			fmt.Sprintf("the format %s value is incorrect, it should be an ipv4", value))
	}

	return AccountIP{
		value: value,
	}, nil
}

func (a *AccountIP) String() string {
	return a.value
}
