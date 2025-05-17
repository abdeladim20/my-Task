package models

import (
	"time"

	"social-network/backend/pkg/db/sqlite"
)

type Comment struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	Image     *string   `json:"image,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Comment) Create() error {
	res, err := sqlite.DB.Exec(
		`INSERT INTO comments (post_id, user_id, content, image) VALUES (?, ?, ?, ?)`,
		c.PostID, c.UserID, c.Content, c.Image,
	)
	if err != nil {
		return err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	c.ID = int(lastID)
	return nil
}

func GetCommentsByPostID(postID int) ([]Comment, error) {
	rows, err := sqlite.DB.Query(
		`SELECT id, post_id, user_id, content, image, created_at, updated_at FROM comments WHERE post_id = ? ORDER BY created_at ASC`, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var c Comment
		if err := rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.Content, &c.Image, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}
