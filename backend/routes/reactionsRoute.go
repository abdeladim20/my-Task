package routes

import (
	"encoding/json"
	"net/http"

	"social-network/backend/pkg/models"
	"social-network/backend/utils"
)

// POST /api/posts/react
func HandleReactToPost(w http.ResponseWriter, r *http.Request) {
	type RequestBody struct {
		PostID int    `json:"post_id"`
		UserID int    `json:"user_id"`
		Type   string `json:"type"` // "like" or "dislike"
	}

	var req RequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.CreateResponseAndLogger(w, http.StatusBadRequest, err, "Invalid request body")
		return
	}

	if req.Type != "like" && req.Type != "dislike" {
		http.Error(w, "Invalid reaction type", http.StatusBadRequest)
		return
	}

	reaction := models.Reaction{
		PostID: req.PostID,
		UserID: req.UserID,
		Type:   req.Type,
	}

	if err := reaction.SaveOrUpdate(); err != nil {
		utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err, "Could not save reaction")
		return
	}

	likes, err := models.CountReactions(req.PostID, "like")
	if err != nil {
		utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err, "Failed to count likes")
		return
	}
	dislikes, err := models.CountReactions(req.PostID, "dislike")
	if err != nil {
		utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err, "Failed to count dislikes")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"message":  "Reaction saved",
		"likes":    likes,
		"dislikes": dislikes,
	})

	// json.NewEncoder(w).Encode(map[string]string{"message": "Reaction saved"})
}
