package server

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/schedule-api/pkg/user"
)

type Server struct {
	db     *sqlx.DB
	Router *mux.Router
	user   *user.Service
}

func NewServer() (*Server, error) {
	LoadVariablesIfNecessary()

	db, err := NewDatabase()
	if err != nil {
		return nil, err
	}
	server := &Server{
		db:     db,
		Router: mux.NewRouter(),
		user:   user.NewService(db),
	}
	server.routes()
	return server, nil
}

func LoadVariablesIfNecessary() {
	env := os.Getenv("APP_ENV")
	if env != "production" {
		godotenv.Load()
		env = os.Getenv("APP_ENV")
	}
	log.Printf("Starting environment %s server.", env)
}

func handleHealthCheck(server Server) http.HandlerFunc {
	type response struct {
		Message string `json:"message"`
	}
	return func(res http.ResponseWriter, req *http.Request) {

		err := server.db.Ping()
		if err != nil {
			res.Header().Add("Content-Type", "application/json")
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(response{Message: "Database connection Error: " + err.Error()})
			return
		}

		res.Header().Add("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(response{Message: "ok"})
	}
}
