package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"social-network/backend/pkg/models"
	"social-network/backend/utils"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		utils.CreateResponseAndLogger(w, http.StatusBadRequest, err, "Could not parse multipart form")
		return
	}

	var imagePath *string
	file, handler, err := r.FormFile("image")
	if err == nil {
		path, err := utils.SaveUploadedImage(file, handler)
		if err != nil {
			utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err, "Failed to save image")
			return
		}
		imagePath = &path
	}

	postIDStr := r.FormValue("post_id")
	userIDStr := r.FormValue("user_id")
	content := r.FormValue("content")

	postID, err := strconv.Atoi(postIDStr)
	if err != nil || postID == 0 {
		utils.CreateResponseAndLogger(w, http.StatusBadRequest, err, "Invalid or missing post_id")
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil || userID == 0 {
		utils.CreateResponseAndLogger(w, http.StatusBadRequest, err, "Invalid or missing user_id")
		return
	}
	if content == "" {
		utils.CreateResponseAndLogger(w, http.StatusBadRequest, err, "Content is required")
		return
	}

	comment := models.Comment{
		PostID:  postID,
		UserID:  userID,
		Content: content,
		Image:   imagePath,
	}

	if err := comment.Create(); err != nil {
		utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err, "Failed to create comment")
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

func GetCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	// Extract the postID from the URL
	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/posts/"), "/")
	if len(pathParts) < 2 || pathParts[1] != "comments" {
		http.Error(w, "Invalid URL", http.StatusNotFound)
		return
	}

	postIDStr := pathParts[0]
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		utils.CreateResponseAndLogger(w, http.StatusBadRequest, err, "Invalid post ID")
		return
	}

	comments, err := models.GetCommentsByPostID(postID)
	if err != nil {
		utils.CreateResponseAndLogger(w, http.StatusInternalServerError, err, "GetCommentsByPostID Failed")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}
