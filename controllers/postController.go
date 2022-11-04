package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tony-starikov/go-rest/initializers"
	"github.com/tony-starikov/go-rest/models"
	"log"
)

func CreatePost(c *gin.Context) {
	// Get data from request
	var newPost struct{
		Title string
		Body string
	}

	c.Bind(&newPost)

	// Create a post
	post := models.Post{Title: newPost.Title, Body: newPost.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		log.Fatal("Failed to crate a post")
		return
	}

	// Return created post
	c.JSON(200, gin.H{
		"post": post,
	})
}