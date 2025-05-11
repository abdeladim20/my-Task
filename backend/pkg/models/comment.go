package models

import (
	"strconv"
	"time"

	"social-network/backend/pkg/db/sqlite"

	"github.com/gofrs/uuid"
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
	id, _ := uuid.NewV4()
    var err error
	c.ID, err = strconv.Atoi(id.String()) // need to be int
    if err != nil {
        return err
    }
	_, err = sqlite.DB.Exec(
		`INSERT INTO comments (id, post_id, user_id, content) VALUES (?, ?, ?, ?)`,
		c.ID, c.PostID, c.UserID, c.Content,
	)
	return err
}

// Get comments by post ID
func GetCommentsByPostID(postID string) ([]Comment, error) {
	rows, err := sqlite.DB.Query(`SELECT * FROM comments WHERE post_id = ? ORDER BY created_at ASC`, postID)
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
