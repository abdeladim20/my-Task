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
		AllowedOrigins:   []string{"http://localhost:5173"}, // Frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true, // Allow cookies if needed
	})

	// Apply the CORS middleware
	handler := corsHandler.Handler(router)

	// Handle OPTIONS requests explicitly (preflight request)
	router.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // Frontend URL
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.WriteHeader(http.StatusOK)
			return
		}
		// Your actual POST/GET logic here
	})

	sqlite.HelpeMegration()

	// Start the server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
