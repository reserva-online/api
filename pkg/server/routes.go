package server

func (s Server) routes() {
	s.Router.HandleFunc("/health-check", handleHealthCheck(s)).Methods("GET")

	s.Router.HandleFunc("/login", handleUserLogin(s.user)).Methods("POST")
	s.Router.HandleFunc("/user", handleUserSave(s.user)).Methods("POST")

	authenticatedRouter := s.Router.NewRoute().Subrouter()
	authenticatedRouter.Use(s.authMiddleware)

	authenticatedRouter.HandleFunc("/user/{id:[0-9]+}", handleUserGetById(s.user)).Methods("GET")

}
