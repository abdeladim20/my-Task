package main

import (
	"log"
	"net/http"

	"social-network/backend/pkg/db/sqlite"
	"social-network/backend/routes"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Connect to the database
	sqlite.Connect()

	// Perform migration
	sqlite.HelpeMegration()

	// Create a new router
	router := mux.NewRouter()

	// Register routes
	routes.RegisterRoutes(router)

	// Set up CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Wrap router with CORS middleware
	handler := corsHandler.Handler(router)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}