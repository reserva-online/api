package server

import (
	"net/http"

	"github.com/schedule-api/pkg/server/handler"
)

func (s Server) routes() {
	s.Router.HandleFunc("/health-check", handleHealthCheck(s)).Methods("GET")
	s.Router.HandleFunc("/login", handler.HandleUserLogin(s.user)).Methods("POST")
	s.Router.HandleFunc("/user", handler.HandleUserSave(s.user)).Methods("POST")

	authenticatedRouter := s.Router.NewRoute().Subrouter()
	authenticatedRouter.Use(s.authenticatedMiddleware)

	authenticatedRouter.HandleFunc("/user/{id:[0-9]+}", handler.HandleUserGetById(s.user)).Methods("GET")

}

func (s *Server) authenticatedMiddleware(httpHandler http.Handler) http.Handler {
	return handler.HandleAuthentication(httpHandler)
}
