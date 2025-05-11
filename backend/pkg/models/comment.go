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
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Create a new comment
func (c *Comment) Create() error {
	res, err := sqlite.DB.Exec(
		`INSERT INTO comments (post_id, user_id, content) VALUES (?, ?, ?)`,
		c.PostID, c.UserID, c.Content,
	)
	if err != nil {
		return err
	}
	// Get the auto-generated ID from SQLite
	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	c.ID = int(lastID)
	return nil
}

// Get comments by post ID
func GetCommentsByPostID(postID int) ([]Comment, error) {
	// rows, err := sqlite.DB.Query(`SELECT * FROM comments WHERE post_id = ? ORDER BY created_at ASC`, postID)
	rows, err := sqlite.DB.Query(`SELECT id, post_id, user_id, content, created_at, updated_at FROM comments WHERE post_id = ? ORDER BY created_at ASC`, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.ID, &comment.PostID, &comment.UserID, &comment.Content, &comment.CreatedAt, &comment.UpdatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
