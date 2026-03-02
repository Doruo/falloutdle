package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"

	"github.com/doruo/falloutdle/internal/character"
	"github.com/doruo/falloutdle/internal/game"
	"gorm.io/gorm"
)

func NewDatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = db.AutoMigrate(&character.Character{}, &game.Game{}); err != nil {
		log.Fatal("Failed to migrate tables:", err)
	}

	return db
}
