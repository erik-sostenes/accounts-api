package ports

import "github.com/erik-sostenes/accounts-api/internal/mooc/account/business/domain"

type (
	// AccountManager represents the Left Side
	// manages business logic resolves user account requirements
	AccountManager interface {
		// Create method that instantiates domain domain.Account object with DTO values
		Create(domain.Account) error
	}

	// Storer represents the Right Side
	// manages the persistence of any data in a DB or in memory, with the required operations
	AccountStorer[V domain.Account] interface {
		// Save method that saves any data to database
		Save(v V) error
	}
)
