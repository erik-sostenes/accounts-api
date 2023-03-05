package backend

import (
	"fmt"
	"github.com/erik-sostenes/accounts-api/internal/shared/backend/jwt"
	"github.com/erik-sostenes/accounts-api/internal/shared/backend/middleware"
	"log"
	"net/http"

	"github.com/erik-sostenes/accounts-api/internal/apps/backend/controllers"
)

// Server contains the configuration of server to start and register all http handlers
type Server struct {
	*http.Server
	jwt.Token[jwt.Claims]
	controllers.Controllers
}

// Start initialize the server with all http handlers
func (s *Server) Start() error {
	s.Server.Handler = s.setRoutes()

	log.Println(s.Server.Addr)

	return s.Server.ListenAndServe()
}

// Routes register all endpoints
//
// configure the middlewares CORS
func (s *Server) setRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/v1/account/create", s.Controllers.AccountController.Create)
	mux.HandleFunc("/v1/account/authenticate", s.AuthController.Authenticate)

	mux.HandleFunc("/v1/account/authorize", middleware.Authorization(s.Token, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "protected route")
	}))

	return mux
}
