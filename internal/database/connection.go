package database

import (
	"log"

	"github.com/doruo/falloutdle/internal/character"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDatabase() *gorm.DB {
	// Database
	dsn := "host=localhost user=username password=password dbname=falloutdle port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Logs SQL (dev)
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migration
	err = db.AutoMigrate(&character.Character{})
	if err != nil {
		log.Fatal("Failed to migrate:", err)
	}

	return db
}
