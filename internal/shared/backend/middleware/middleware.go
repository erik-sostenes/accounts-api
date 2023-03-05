package middleware

import (
	"net/http"

	"github.com/erik-sostenes/accounts-api/internal/shared/backend/controllers"
	jwt2 "github.com/erik-sostenes/accounts-api/internal/shared/backend/jwt"
	"github.com/erik-sostenes/accounts-api/internal/shared/mooc/business/domain/wrongs"
)

// Authorization middleware that validates the token for each http request
// if the token is invalid the client is responded to with a StatusForbidden
//
// if the token is valid the requested HandlerFunc is executed
func Authorization(jwt jwt2.Token[jwt2.Claims], next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")

		err := jwt.Validate(token)
		if err != nil {
			_ = controllers.HandlerError(w, wrongs.StatusForbidden(err.Error()))
			return
		}

		next.ServeHTTP(w, r)
	}
}
