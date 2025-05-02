package models

type Post struct {
    ID         int      `json:"id"`
    UserID     int      `json:"user_id"`
    Content    string   `json:"content"`
    Visibility string   `json:"visibility"`
    Audience   []int    `json:"audience,omitempty"`
    Comments   []Comment `json:"comments,omitempty"`
}

type Comment struct {
    ID      int    `json:"id"`
    PostID  int    `json:"post_id"`
    UserID  int    `json:"user_id"`
    Content string `json:"content"`
}