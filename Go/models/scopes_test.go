package models

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{}, &Category{})

	return db
}

func TestMinPrice(t *testing.T) {
	db := setupTestDB()

	products := []Product{
		{Name: "Cheap Product", Price: 10.0},
		{Name: "Affordable Product", Price: 50.0},
		{Name: "Expensive Product", Price: 100.0},
	}
	for _, product := range products {
		db.Create(&product)
	}

	var result []Product
	db.Scopes(MinPrice(30)).Find(&result)

	assert.Equal(t, 2, len(result))
	assert.Equal(t, "Affordable Product", result[0].Name)
	assert.Equal(t, 50.0, result[0].Price)
	assert.Equal(t, "Expensive Product", result[1].Name)
	assert.Equal(t, 100.0, result[1].Price)

	var allProducts []Product
	db.Find(&allProducts)
	assert.Equal(t, 3, len(allProducts))
	assert.Equal(t, "Cheap Product", allProducts[0].Name)
	assert.Equal(t, 10.0, allProducts[0].Price)
	assert.Equal(t, "Affordable Product", allProducts[1].Name)
	assert.Equal(t, "Expensive Product", allProducts[2].Name)
	assert.True(t, allProducts[2].Price > allProducts[1].Price)
}

func TestByCategory(t *testing.T) {
	db := setupTestDB()

	category1 := Category{Name: "Category 1"}
	category2 := Category{Name: "Category 2"}
	db.Create(&category1)
	db.Create(&category2)

	products := []Product{
		{Name: "Product 1", CategoryID: category1.ID},
		{Name: "Product 2", CategoryID: category2.ID},
		{Name: "Product 3", CategoryID: category1.ID},
	}
	for _, product := range products {
		db.Create(&product)
	}

	var result []Product
	db.Scopes(ByCategory(category1.ID)).Find(&result)

	assert.Equal(t, 2, len(result))
	assert.Equal(t, "Product 1", result[0].Name)
	assert.Equal(t, category1.ID, result[0].CategoryID)
	assert.Equal(t, "Product 3", result[1].Name)
	assert.Equal(t, category1.ID, result[1].CategoryID)

	var allProducts []Product
	db.Find(&allProducts)
	assert.Equal(t, 3, len(allProducts))
	assert.Equal(t, "Product 1", allProducts[0].Name)
	assert.Equal(t, "Product 2", allProducts[1].Name)
	assert.Equal(t, "Product 3", allProducts[2].Name)
	assert.Equal(t, category2.ID, allProducts[1].CategoryID)
}

func TestSortProducts(t *testing.T) {
	db := setupTestDB()

	products := []Product{
		{Name: "Product A", Price: 30.0},
		{Name: "Product B", Price: 10.0},
		{Name: "Product C", Price: 20.0},
	}
	for _, product := range products {
		db.Create(&product)
	}

	var resultAsc []Product
	db.Scopes(SortProducts("price", false)).Find(&resultAsc)

	assert.Equal(t, 3, len(resultAsc))
	assert.Equal(t, "Product B", resultAsc[0].Name)
	assert.Equal(t, 10.0, resultAsc[0].Price)
	assert.Equal(t, "Product C", resultAsc[1].Name)
	assert.Equal(t, 20.0, resultAsc[1].Price)
	assert.Equal(t, "Product A", resultAsc[2].Name)
	assert.Equal(t, 30.0, resultAsc[2].Price)

	assert.True(t, resultAsc[0].Price < resultAsc[1].Price)
	assert.True(t, resultAsc[1].Price < resultAsc[2].Price)

	var resultDesc []Product
	db.Scopes(SortProducts("price", true)).Find(&resultDesc)

	assert.Equal(t, 3, len(resultDesc))
	assert.Equal(t, "Product A", resultDesc[0].Name)
	assert.Equal(t, 30.0, resultDesc[0].Price)
	assert.Equal(t, "Product C", resultDesc[1].Name)
	assert.Equal(t, 20.0, resultDesc[1].Price)
	assert.Equal(t, "Product B", resultDesc[2].Name)
	assert.Equal(t, 10.0, resultDesc[2].Price)

	assert.True(t, resultDesc[0].Price > resultDesc[1].Price)
	assert.True(t, resultDesc[1].Price > resultDesc[2].Price)
}
