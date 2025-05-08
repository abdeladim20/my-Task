package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	// Health check route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Social Network!"))
	}).Methods("GET")

	// Post routes
	router.HandleFunc("/posts", CreatePost).Methods("POST")
	router.HandleFunc("/posts", GetPosts).Methods("GET")
}
