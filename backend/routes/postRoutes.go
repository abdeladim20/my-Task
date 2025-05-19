package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"social-network/backend/pkg/db/sqlite"
	"social-network/backend/pkg/models"
	"social-network/backend/utils"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10MB max
	if err != nil {
		utils.CreateResponseAndLogger(w, http.StatusBadRequest, err, "Cannot parse form")
		return
	}

	// Get form values
	userIDStr := r.FormValue("user_id")
	title := r.FormValue("title")
	content := r.FormValue("content")
	privacy := r.FormValue("privacy")

	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID == 0 || title == "" || content == "" {
		utils.CreateResponseAndLogger(w, http.StatusBadRequest, err, "Missing required fields")
		return
	}

	var imagePath *string

	// Get the file from form input "image"
	file, handler, err := r.FormFile("image")
	if err == nil {
		defer file.Close()
		// Create a destination file
		filename := fmt.Sprintf("backend/static/uploads/%d_%s", time.Now().UnixNano(), handler.Filename)
		// filename := fmt.Sprintf("static/uploads/%d_%s", time.Now().UnixNano(), handler.Filename)
		dst, err := os.Create(filename)
		if err != nil {
			utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err, "Unable to save the file")
			return
		}
		defer dst.Close()

		// Copy uploaded file to destination
		if _, err := io.Copy(dst, file); err != nil {
			utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err, "Unable to save the file")
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
		utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err, "Failed to create post")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	rows, err := sqlite.DB.Query("SELECT id, user_id, title, content, image, privacy, created_at, updated_at FROM posts ORDER BY created_at DESC")
	if err != nil {
		utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err, "Failed to fetch posts")
		return
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Image, &post.Privacy, &post.CreatedAt, &post.UpdatedAt); err != nil {
			utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err, "Failed to scan post")
			// log.Println("Failed to scan post:", err)
			continue
		}
		posts = append(posts, post)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
