package database

import (
	"go-app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!", err)
	}

	err = database.AutoMigrate(&models.Product{})
	if err != nil {
		return
	}

	DB = database
}
