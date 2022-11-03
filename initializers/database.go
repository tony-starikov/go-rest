package initializers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnectToDatabase() {
	var err error

	dbFileName := os.Getenv("DB_FILE")

	DB, err = gorm.Open(sqlite.Open(dbFileName), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
}