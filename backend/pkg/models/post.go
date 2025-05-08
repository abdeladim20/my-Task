package models

import (
	"strconv"
	"time"

	"social-network/backend/pkg/db/sqlite"

	"github.com/gofrs/uuid"
)

type Post struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	Image     *string   `json:"image,omitempty"`
	Privacy   string    `json:"privacy"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Create a new post
func (p *Post) Create() error {
	var err error
	id, _ := uuid.NewV4()
	p.ID, err = strconv.Atoi(id.String())
	if err != nil {
		return err
	}
	_, err = sqlite.DB.Exec(
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
