package database

import (
	"fmt"
	"log"
	"os"

	"github.com/doruo/falloutdle/internal/character"
	"github.com/doruo/falloutdle/internal/game"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDatabaseConnection creates and returns a new database connection instance
func NewDatabaseConnection() (db *gorm.DB) {

	// Database connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, error := gorm.Open(postgres.Open(dsn))

	if error != nil {
		log.Fatal("Failed to connect to database:", error)
	}

	// Auto-migration
	error = db.AutoMigrate(&character.Character{})
	if error != nil {
		log.Fatal("Failed to migrate table Character:", error)
	}

	error = db.AutoMigrate(&game.Game{})
	if error != nil {
		log.Fatal("Failed to migrate table Game:", error)
	}

	return
}
