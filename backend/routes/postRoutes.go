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

	// Decode the JSON payload
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		log.Println("Invalid request payload:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Check required fields
	if post.UserID == 0 || post.Content == "" {
		log.Println("UserID and Content are required")
		http.Error(w, "UserID and Content are required", http.StatusBadRequest)
		return
	}

	// Set the image to NULL if it's empty
	if post.Image == nil {
		post.Image = nil
	}

	// Insert the post into the database
	res, err := sqlite.DB.Exec(
		"INSERT INTO posts (user_id, content, image, privacy) VALUES (?, ?, NULLIF(?, 'NULL'), ?)",
		post.UserID, post.Content, post.Image, post.Privacy,
	)
	if err != nil {
		log.Println("Failed to create post:", err)
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	post.ID = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	rows, err := sqlite.DB.Query("SELECT id, user_id, content, image, privacy, created_at, updated_at FROM posts ORDER BY created_at DESC")
	if err != nil {
		log.Println("Failed to fetch posts:", err)
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.Image, &post.Privacy, &post.CreatedAt, &post.UpdatedAt); err != nil {
			log.Println("Failed to scan post:", err)
			continue
		}
		posts = append(posts, post)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
