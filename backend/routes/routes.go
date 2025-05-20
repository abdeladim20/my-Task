package routes

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	// Test check route
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Write([]byte("Welcome to the Social Network!"))
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Post routes
	mux.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			CreatePost(w, r)
		case http.MethodGet:
			GetPosts(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Comment routes
	mux.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			CreateComment(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	// Handle dynamic post comment routes
	mux.HandleFunc("/posts/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			GetCommentsByPostID(w, r)
			return
		}
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	})

	// Serve uploaded images statically
	mux.Handle("/static/uploads/", http.StripPrefix("/static/uploads", http.FileServer(http.Dir("./backend/static/uploads"))))

}
