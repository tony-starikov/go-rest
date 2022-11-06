package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tony-starikov/go-rest/controllers"
	"github.com/tony-starikov/go-rest/initializers"
	"github.com/tony-starikov/go-rest/migrations"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
	migrations.Migrate()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.CreateSinglePost)
	r.GET("/posts", controllers.ShowAllPosts)
	r.GET("/posts/:id", controllers.ShowSinglePost)
	r.PUT("/posts/:id", controllers.UpdateSinglePost)
	r.DELETE("/posts/:id", controllers.DeleteSinglePost)

	r.Run()
}