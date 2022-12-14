package migrations

import (
	"fmt"
	"github.com/tony-starikov/go-rest/initializers"
	"github.com/tony-starikov/go-rest/models"
	"log"
)

func Migrate() {
	err := initializers.DB.AutoMigrate(&models.Post{})

	if err != nil {
		log.Fatal("Failed to migrate")
	} else {
		fmt.Println("Migration success")
	}
}