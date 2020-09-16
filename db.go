package main

import (
	"fmt"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

var db *gorm.DB

// InitDB Connects to the database and migrates the schema
func InitDB() {
  var err error
  db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  db.AutoMigrate(&Post{}, &Comment{})
}

// CreatePost inserts a post into the database
func CreatePost(post *Post) error {
  result := db.Create(post)
  return result.Error
}

// GetPosts returns all posts in the database
func GetPosts() []Post {
  var posts []Post
  db.Find(&posts)
  return posts;
}

// GetPostByID returns a post with the given primary key
func GetPostByID(id int) (Post, error) {
  var post Post
  var comments []Comment
  // Get the post from the database
  result := db.First(&post, id)
  if result.Error != nil {
    return post, result.Error
  }

  // Get all comments for this post
  if db.Model(&post).Association("Comments").Error != nil {
    return post, db.Model(&post).Association("Comments").Error
  }
  err := db.Model(&post).Association("Comments").Find(&comments)
  // Assign comments to post
  post.Comments = comments
  return post, err
}

// CreateComment inserts a new comment into the database
func CreateComment(comment *Comment) error {
  // Ensure post exists
  result := db.First(&Post{}, comment.PostID)
  if result.Error != nil {
    return fmt.Errorf("No post with id of %d", comment.PostID)
  }

  result = db.Create(&comment)
  return result.Error
}