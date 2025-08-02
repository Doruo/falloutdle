package database

import "gorm.io/gorm"

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
