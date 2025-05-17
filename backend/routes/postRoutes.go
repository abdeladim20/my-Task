package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"social-network/backend/pkg/db/sqlite"
	"social-network/backend/pkg/models"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10MB max
	if err != nil {
		http.Error(w, "Cannot parse form", http.StatusBadRequest)
		return
	}

	// Get form values
	userIDStr := r.FormValue("user_id")
	title := r.FormValue("title")
	content := r.FormValue("content")
	privacy := r.FormValue("privacy")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID == 0 || title == "" || content == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	var imagePath *string

	// Get the file from form input "image"
	file, handler, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
		// Create a destination file
		filename := fmt.Sprintf("uploads/%d_%s", time.Now().UnixNano(), handler.Filename)
		dst, err := os.Create(filename)
		if err != nil {
			http.Error(w, "Unable to save the file", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		// Copy uploaded file to destination
		if _, err := io.Copy(dst, file); err != nil {
			http.Error(w, "Unable to save the file", http.StatusInternalServerError)
			return
		}

		imagePath = &filename
	}

	post := models.Post{
		UserID:  userID,
		Title:   title,
		Content: content,
		Privacy: privacy,
		Image:   imagePath,
	}

	if err := post.Create(); err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

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
