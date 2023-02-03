package domain

import (
	"fmt"

	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/wrongs"
	"github.com/google/uuid"
)

// Identifier receives a value to verify if the format is correct
type Identifier string

// ParseUuID validate if the format the values is a UuID
func (i Identifier) ParseUuID() (string, error) {
	id, err := uuid.Parse(string(i))
	if err != nil {
		return "", wrongs.StatusUnprocessableEntity(fmt.Sprintf("incorrect %s uuid unique identifier", string(i)))
	}

	return id.String(), nil
}
