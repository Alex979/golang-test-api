package main

import "gorm.io/gorm"

// Post - Model defining a blog post
type Post struct {
	gorm.Model
	User     string
	Title    string
	Content  string
	Comments []Comment
}

// PostJSON - JSON format for a POST request creating a new Post
type PostJSON struct {
	User    string `json:"user" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// Comment - Model defining a comment on a post
type Comment struct {
	gorm.Model
	User    string
	Content string
	PostID  uint
}

// CommentJSON - JSON format for a POST request creating a new Comment
type CommentJSON struct {
	User    string `json:"user" binding:"required"`
	Content string `json:"content" binding:"required"`
}
