package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tony-starikov/go-rest/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello!",
		})
	})
	r.Run()
}