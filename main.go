package main

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize gin and connect to the database
	r := gin.Default()
	InitDB()

	// Return a list of all posts
	r.GET("/posts", func(c *gin.Context) {
		posts := GetPosts()
		c.JSON(200, gin.H{"posts": posts})
	})

	// Returns a single post by ID
	r.GET("/posts/:id", func(c *gin.Context) {
		// Get id from URL
		idString := c.Param("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "id must be an unsigned integer"})
			return
		}

		// Find matching post
		post, err := GetPostByID(uint(id))
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"post": post})
	})

	// Upload a post
	r.POST("/posts", func(c *gin.Context) {
		// Map JSON to struct
		var json PostJSON
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Create post and upoad to database
		post := Post{
			User: json.User,
			Title: json.Title,
			Content: json.Content,
		}
		if err := CreatePost(&post); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"post": post})
	});

	// Upload a comment on a post
	r.POST("/posts/:id/comment", func(c *gin.Context) {
		// Get id from URL
		idString := c.Param("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "id must be an unsigned integer"})
			return
		}

		// Map JSON to struct
		var json CommentJSON
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Create comment and upload to database
		comment := Comment{
			User: json.User,
			Content: json.Content,
			PostID: uint(id),
		}
		if err := CreateComment(&comment); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"comment": comment})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}