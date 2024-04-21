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

	err = database.AutoMigrate(&models.Category{}, &models.Product{}, &models.Cart{}, &models.CartItem{}, &models.Payment{})
	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}

	DB = database
}
