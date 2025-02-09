package config

import (
	"fmt"
	"inventory-book/models"
)

// MigrateDB migrates the database
func MigrateDB() {
	if !DB.Migrator().HasTable(&models.User{}) || !DB.Migrator().HasTable(&models.Book{}) {
		DB.AutoMigrate(&models.User{}, &models.Book{})
		fmt.Println("Database migrated successfully!")
	} else {
		fmt.Println("Database already migrated!")
	}
}
