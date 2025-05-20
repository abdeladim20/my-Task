package main

import (
	"log"
	"net/http"

	"social-network/backend/pkg/db/sqlite"
	"social-network/backend/routes"
)

func main() {
	// Connect to the database
	sqlite.Connect()

	// Perform migration
	sqlite.HelpeMegration()

	// Create a multiplexer (router) using the standard library
	mux := http.NewServeMux()

	// Register routes
	routes.RegisterRoutes(mux)

	// Wrap with CORS middleware
	handler := corsMiddleware(mux)

	log.Println("Server running on :8080", "http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

// Basic CORS middleware
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Pass to the next handler
		next.ServeHTTP(w, r)
	})
}
