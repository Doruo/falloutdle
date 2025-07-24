package database

import (
	"fmt"
	"log"
	"os"

	"github.com/doruo/falloutdle/internal/character"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Single pattern for single database connection.
var instance *gorm.DB

// GetInstance returns a single instance.
// Creates a new one if nil.
func GetInstance() *gorm.DB {
	if instance == nil {
		instance = NewDatabaseConnection()
	}
	return instance
}

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

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migration
	err = db.AutoMigrate(&character.Character{})
	if err != nil {
		log.Fatal("Failed to migrate:", err)
	}

	return
}
