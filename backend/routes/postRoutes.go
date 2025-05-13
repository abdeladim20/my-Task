package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"social-network/backend/pkg/db/sqlite"
	"social-network/backend/pkg/models"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	// Decode JSON
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check required fields
	if post.UserID == 0 || post.Content == "" || post.Title == "" {
		http.Error(w, "UserID, title and Content are required", http.StatusBadRequest)
		return
	}

	// Create post using model
	if err := post.Create(); err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	// Return created post
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	rows, err := sqlite.DB.Query("SELECT id, user_id, title, content, image, privacy, created_at, updated_at FROM posts ORDER BY created_at DESC")
	if err != nil {
		log.Println("Failed to fetch posts:", err)
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Image, &post.Privacy, &post.CreatedAt, &post.UpdatedAt); err != nil {
			log.Println("Failed to scan post:", err)
			continue
		}
		posts = append(posts, post)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
