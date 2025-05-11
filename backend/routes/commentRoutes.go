package routes

import (
	"encoding/json"
	"net/http"
	"social-network/backend/pkg/models"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := comment.Create(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

func GetCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// postID := vars["postID"]
	var postID models.Post

	comments, err := models.GetCommentsByPostID(postID.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comments)
}
