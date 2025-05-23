package models

import (
	"encoding/json"
	"net/http"
	"time"

	"social-network/backend/pkg/db/sqlite"
	"social-network/backend/utils"
)

type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Image     *string   `json:"image,omitempty"`
	Privacy   string    `json:"privacy"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
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

func GetReactionCounts(postID int) (likes int, dislikes int, err error) {
	err = sqlite.DB.QueryRow(`
		SELECT COUNT(*) FROM reactions WHERE post_id = ? AND type = 'like'
	`, postID).Scan(&likes)
	if err != nil {
		return
	}

	err = sqlite.DB.QueryRow(`
		SELECT COUNT(*) FROM reactions WHERE post_id = ? AND type = 'dislike'
	`, postID).Scan(&dislikes)
	return
}

// Get all posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	rows, err := sqlite.DB.Query("SELECT id, user_id, title, content, image, privacy, created_at, updated_at FROM posts ORDER BY created_at DESC")
	if err != nil {
		utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err, "Failed to fetch posts")
		// http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content, &post.Image, &post.Privacy, &post.CreatedAt, &post.UpdatedAt); err != nil {
			continue
		}
		// Count reactions
		likes, dislikes, err := GetReactionCounts(post.ID)
		if err == nil {
			post.Likes = likes
			post.Dislikes = dislikes
		}

		posts = append(posts, post)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}


