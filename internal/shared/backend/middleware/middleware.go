package middleware

import (
	"errors"
	"net/http"

	"github.com/erik-sostenes/accounts-api/internal/shared/backend/controllers"
	jwt2 "github.com/erik-sostenes/accounts-api/internal/shared/backend/jwt"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/wrongs"
)

// Authenticated middleware that validates the token for each http request
// if the token is invalid the client is responded to with a StatusForbidden
//
// if the token is valid the requested HandlerFunc is executed
func Authenticated(jwt jwt2.JWT, next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")

		err := jwt.Validate(token)
		if errors.Is(err, jwt2.ValidationError) {
			controllers.HandlerError(w, wrongs.StatusForbidden(err.Error()))
			return
		}

		if err != nil {
			controllers.HandlerError(w, err)
			return
		}

		next.ServeHTTP(w, r)
	})
}
