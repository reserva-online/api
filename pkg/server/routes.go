package server

func (s Server) routes() {
	s.Router.HandleFunc("/health-check", handleHealthCheck(s)).Methods("GET")
}
