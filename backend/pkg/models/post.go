package models

import (
	"encoding/json"
	"net/http"
	"time"

	"social-network/backend/pkg/db/sqlite"
)

type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Image     *string   `json:"image,omitempty"`
	Privacy   string    `json:"privacy"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Comments  []Comment `json:"comments,omitempty"`
}

// Create a new post
func (p *Post) Create() error {
	res, err := sqlite.DB.Exec(
		`INSERT INTO posts (user_id, title ,content, image, privacy) VALUES (?, ?, ?, ?, ?)`,
		p.UserID, p.Title, p.Content, p.Image, p.Privacy,
	)
	if err != nil {
		return err
	}

	// Get the auto-generated ID from SQLite
	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = int(lastID)
	return nil
}

// Get all posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	rows, err := sqlite.DB.Query("SELECT id, user_id, title, content, image, privacy, created_at, updated_at FROM posts ORDER BY created_at DESC")
	if err != nil {
		http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.UserID,&post.Title ,&post.Content, &post.Image, &post.Privacy, &post.CreatedAt, &post.UpdatedAt); err != nil {
			continue
		}

		// Fetch comments for this post
		// comments, err := GetCommentsByPostID(post.ID)
		// if err == nil {
		// 	post.Comments = comments
		// }

		posts = append(posts, post)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
