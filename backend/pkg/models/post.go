package models

import (
	"time"

	"social-network/backend/pkg/db/sqlite"

	"github.com/gofrs/uuid"
)

type Post struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Content   string    `json:"content"`
	Image     *string   `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Create a new post
func (p *Post) Create() error {
	id, _ := uuid.NewV4()
	p.ID = id.String()
	_, err := sqlite.DB.Exec(
		`INSERT INTO posts (id, user_id, content, image) VALUES (?, ?, ?, ?)`,
		p.ID, p.UserID, p.Content, p.Image,
	)
	return err
}

// Get all posts
func GetAllPosts() ([]Post, error) {
	rows, err := sqlite.DB.Query(`SELECT * FROM posts ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.UserID, &post.Content, &post.Image, &post.CreatedAt, &post.UpdatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
