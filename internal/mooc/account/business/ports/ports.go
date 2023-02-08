package ports

import (
	"context"

	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"
)

type (
	// AccountManager represents the Left Side
	// manages business logic resolves user account requirements
	AccountManager interface {
		// Create method that instantiates domain domain.Account object with DTO values
		Create(context.Context, domain.Account) error
	}

	// Storer represents the Right Side
	// manages the persistence of any data in a DB or in memory, with the required operations
	Storer[K comparable, V any] interface {
		// Save method that saves any data to database
		Save(context.Context, K, V) error
		//Remove method that removes any data to the database by means of an identifier
		Remove(context.Context, K) (err error) 
	}
)
