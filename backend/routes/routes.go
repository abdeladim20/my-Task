package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	// test check route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Social Network!"))
	}).Methods("GET")

	// Post routes
	router.HandleFunc("/posts", CreatePost).Methods("POST")
	router.HandleFunc("/posts", GetPosts).Methods("GET")

	// Comment routes
	router.HandleFunc("/comments", CreateComment).Methods("POST")
	router.HandleFunc("/posts/{postID}/comments", GetCommentsByPostID).Methods("GET")

	// Serve uploaded images statically
	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/"))))
}