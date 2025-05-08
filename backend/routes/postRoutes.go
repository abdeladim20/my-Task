package routes

import (
    "encoding/json"
    "net/http"
    "social-network/backend/pkg/models"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
    var post models.Post
    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := post.Create(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(post)
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
    posts, err := models.GetAllPosts()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(posts)
}
