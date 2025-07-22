package database

import (
	"fmt"
	"log"
	"os"

	"github.com/doruo/falloutdle/internal/character"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Logs SQL (only on dev)
	})

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
