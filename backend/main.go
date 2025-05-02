package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	Content    string    `json:"content"`
	Visibility string    `json:"visibility"`
	Audience   []int     `json:"audience,omitempty"`
	Comments   []Comment `json:"comments,omitempty"`
}

type Comment struct {
	ID      int    `json:"id"`
	PostID  int    `json:"post_id"`
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}

var posts []Post
var comments []Comment

func main() {
	r := gin.Default()

	r.Use(corsMiddleware())

	r.POST("/api/posts", createPost)
	r.GET("/api/get_posts", getPosts)
	r.POST("/api/comments", createComment)

	r.Run(":8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func createPost(c *gin.Context) {
	var post Post
	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post"})
		return
	}
	post.ID = len(posts) + 1
	posts = append(posts, post)
	c.JSON(http.StatusOK, post)
}

func getPosts(c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, _ := strconv.Atoi(userIDStr)

	var visiblePosts []Post
	for _, post := range posts {
		if post.Visibility == "public" {
			visiblePosts = append(visiblePosts, withComments(post))
		} else if post.Visibility == "private" && post.UserID == userID {
			visiblePosts = append(visiblePosts, withComments(post))
		} else if post.Visibility == "custom" && contains(post.Audience, userID) {
			visiblePosts = append(visiblePosts, withComments(post))
		}
	}

	c.JSON(http.StatusOK, visiblePosts)
}

func createComment(c *gin.Context) {
	var comment Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment"})
		return
	}
	comment.ID = len(comments) + 1
	comments = append(comments, comment)
	c.JSON(http.StatusOK, comment)
}

func withComments(post Post) Post {
	var postComments []Comment
	for _, c := range comments {
		if c.PostID == post.ID {
			postComments = append(postComments, c)
		}
	}
	post.Comments = postComments
	return post
}

func contains(arr []int, val int) bool {
	for _, a := range arr {
		if a == val {
			return true
		}
	}
	return false
}