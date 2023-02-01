package backend

import (
	"log"
	"net/http"

	"github.com/erik-sostenes/accounts-api/internal/apps/backend/controllers"
)

// Server contains the configuration of server to start and register all http handlers
type Server struct {
	*http.Server
	controllers.Controllers
}

// NewServer returns an instance of Server
func NewServer(server *http.Server, controllers controllers.Controllers) *Server {
	return &Server{
		server,
		controllers,
	}
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
	return mux
}
