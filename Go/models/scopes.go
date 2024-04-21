package models

import (
	"gorm.io/gorm"
)

func MinPrice(price float64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("price >= ?", price)
	}
}

func ByCategory(categoryID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("category_id = ?", categoryID)
	}
}

func SortProducts(sortField string, descending bool) func(db *gorm.DB) *gorm.DB {
	sortOrder := "ASC"
	if descending {
		sortOrder = "DESC"
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(sortField + " " + sortOrder)
	}
}
