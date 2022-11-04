package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tony-starikov/go-rest/initializers"
	"github.com/tony-starikov/go-rest/models"
	"log"
)

func CreateSinglePost(c *gin.Context) {
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

func ShowAllPosts(c *gin.Context) {
	// Get all posts from db
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Return all posts
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func ShowSinglePost(c *gin.Context) {
	// Get post id from url
	id := c.Param("id")

	// Get the post from db
	var post models.Post
	initializers.DB.First(&post, id)

	// Return the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdateSinglePost(c *gin.Context) {
	// Get post id from url
	id := c.Param("id")

	// Get data from request
	var updatePost struct{
		Title string
		Body string
	}

	c.Bind(&updatePost)

	// Get the post from db
	var post models.Post
	initializers.DB.First(&post, id)

	// Update the post
	initializers.DB.Model(&post).Updates(models.Post{Title: updatePost.Title , Body: updatePost.Body})

	// Return the post
	c.JSON(200, gin.H{
		"post": post,
	})
}

func DeleteSinglePost(c *gin.Context) {
	// Get post id from url
	id := c.Param("id")

	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)

	// Respond with the post
	c.Status(200)
}