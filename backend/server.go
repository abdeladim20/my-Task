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

	// Create a new router
	router := mux.NewRouter()

	// Register your routes
	routes.RegisterRoutes(router)

	// Set up the CORS options
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, 
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // Allow cookies
	})

	// Apply the CORS middleware
	handler := corsHandler.Handler(router)

	// Handle OPTIONS requests
	router.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") 
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(http.StatusOK)
			return
		}
		// POST/GET logic here
	})

	sqlite.HelpeMegration()

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
