package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	_ "github.com/lib/pq"
	"github.com/schedule-api/pkg/server"
)

func main() {
	log.Printf("Starting server...")
	s, err := server.NewServer()

	if err != nil {
		log.Printf("Error while instantiating server: %s", err)
		os.Exit(1)
	}

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Printf("Server is up!")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, methods, origins)(s.Router)))
}
