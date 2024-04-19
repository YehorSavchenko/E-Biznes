package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text;not null"`
	Products    []Product `gorm:"foreignKey:CategoryID"`
}
