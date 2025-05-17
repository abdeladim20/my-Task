package routes

import (
	"encoding/json"
	"net/http"
	"social-network/backend/pkg/models"
	"social-network/backend/utils"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Could not parse multipart form", http.StatusBadRequest)
		return
	}

	var imagePath *string
	file, handler, err := r.FormFile("image")
	if err == nil {
		path, err := utils.SaveUploadedImage(file, handler)
		if err != nil {
			http.Error(w, "Failed to save image", http.StatusInternalServerError)
			return
		}
		imagePath = &path
	}

	postIDStr := r.FormValue("post_id")
	userIDStr := r.FormValue("user_id")
	content := r.FormValue("content")

	postID, err := strconv.Atoi(postIDStr)
	if err != nil || postID == 0 {
		http.Error(w, "Invalid or missing post_id", http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID == 0 {
		http.Error(w, "Invalid or missing user_id", http.StatusBadRequest)
		return
	}
	if content == "" {
		http.Error(w, "Content is required", http.StatusBadRequest)
		return
	}

	comment := models.Comment{
		PostID:  postID,
		UserID:  userID,
		Content: content,
		Image:   imagePath,
	}

	if err := comment.Create(); err != nil {
		http.Error(w, "Failed to create comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

func GetCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postIdASstring := vars["postID"]
	postID, err := strconv.Atoi(postIdASstring)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	comments, err := models.GetCommentsByPostID(postID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comments)
}
