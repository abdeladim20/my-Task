package main

import (
	"log"
	"net/http"

	"social-network/backend/pkg/db/sqlite"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to the database
	sqlite.Connect()

	// Create a new router
	router := mux.NewRouter()

	// Register your routes
	// routes.RegisterRoutes(router)

	// Start the server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
