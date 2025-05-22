package models

import (
	"database/sql"
	"social-network/backend/pkg/db/sqlite"
	"time"
)

type Reaction struct {
	ID        int       `json:"id"`
	PostID    int       `json:"post_id"`
	UserID    int       `json:"user_id"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// Create or update a reaction
func (r *Reaction) SaveOrUpdate() error {
	_, err := sqlite.DB.Exec(`
		INSERT INTO reactions (post_id, user_id, type)
		VALUES (?, ?, ?)
		ON CONFLICT(post_id, user_id) DO UPDATE SET type=excluded.type
	`, r.PostID, r.UserID, r.Type)
	return err
}

// Count reactions of a given type for a post
func CountReactions(postID int, reactionType string) (int, error) {
	var count int
	err := sqlite.DB.QueryRow(`
		SELECT COUNT(*) FROM reactions WHERE post_id = ? AND type = ?
	`, postID, reactionType).Scan(&count)
	return count, err
}

// Optionally: Get user reaction to a post
func GetUserReaction(postID, userID int) (string, error) {
	var reactionType string
	err := sqlite.DB.QueryRow(`
		SELECT type FROM reactions WHERE post_id = ? AND user_id = ?
	`, postID, userID).Scan(&reactionType)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return reactionType, err
}