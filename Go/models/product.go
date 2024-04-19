package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `gorm:"type:varchar(100);not null"`
	Description string  `gorm:"type:text;not null"`
	Price       float64 `gorm:"type:decimal(10,2);not null"`
	CategoryID  uint
	Category    Category `gorm:"foreignKey:CategoryID"`
}
