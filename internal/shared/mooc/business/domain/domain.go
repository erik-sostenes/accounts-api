package domain

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/wrongs"
)

// String receives a value to verify if the format is correct
type String string

// The Validate method validates if the value is a string and is not empty, if incorrect returns an errors.StatusUnprocessableEntity
func (s String) Validate() (string, error) {
	if strings.TrimSpace(string(s)) == "" {
		return "", wrongs.StatusUnprocessableEntity("Value not found")
	}
	return string(s), nil
}

// Bool receives a value to verify if the format is correct
type Bool string

// Validate method validates if the value is a bool, if incorrect returns an errors.StatusUnprocessableEntity
func (b Bool) Validate() (v bool, err error) {
	v, err = strconv.ParseBool(string(b))
	if err != nil {
		return v, wrongs.StatusUnprocessableEntity(fmt.Sprintf("incorrect %s value format, %v", string(b), err))
	}
	return
}
