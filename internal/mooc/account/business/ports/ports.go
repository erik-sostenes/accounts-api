package ports

import (
	"context"

	"github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"
)

type (
	// AccountCreator represents the Left Side
	// manages business logic create user account requirements
	AccountCreator interface {
		// Create method that instantiates domain domain.Account object with DTO values
		Create(context.Context, domain.Account) error
	}

	// AccountFinder represents the Left Side
	// manages business logic find user account
	AccountFinder[V any] interface {
		// Find method that finds the user account by domain.AccountId and returns any type and an error
		// if an error occurs
		Find(context.Context, domain.AccountId) (V, error)
	}

	// Storer represents the Right Side
	// manages the persistence of any data in a DB or in memory, with the required operations
	Storer[K comparable, V any] interface {
		// Save method that saves any data to database
		Save(context.Context, K, V) error
		//Remove method that removes any data to the database by means of an identifier
		Remove(context.Context, K) error
		// Search method that searches for a resource by identifier
		Search(context.Context, K) (V, error)
	}
)
