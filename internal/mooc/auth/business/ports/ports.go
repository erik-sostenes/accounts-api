package ports

import (
	"context"

	"github.com/erik-sostenes/accounts-api/internal/mooc/auth/business/domain"
	"github.com/erik-sostenes/accounts-api/internal/mooc/auth/business/services"
)

type (
	// Authenticator represents the Left Side
	Authenticator interface {
		// Authenticate ...
		Authenticate(context.Context, domain.AuthId, domain.AuthPassword) (services.AuthResponse, error)
	}
)
