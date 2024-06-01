package models

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

const (
	initialProduct1Name     = "Cheap Product"
	initialProduct1Price    = 10.0
	initialProduct2Name     = "Affordable Product"
	initialProduct2Price    = 50.0
	initialProduct3Name     = "Expensive Product"
	initialProduct3Price    = 100.0
	category1Name           = "Category 1"
	category2Name           = "Category 2"
	product1Name            = "Product 1"
	product2Name            = "Product 2"
	product3Name            = "Product 3"
	sortProductName1        = "Product A"
	sortProductName2        = "Product B"
	sortProductName3        = "Product C"
	sortProductPrice1       = 30.0
	sortProductPrice2       = 10.0
	sortProductPrice3       = 20.0
	invalidJSONPayload      = "Invalid JSON payload"
	invalidCategoryData     = "Invalid category data"
	productNotFoundMessage  = "Product not found"
	categoryNotFoundMessage = "Category not found"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(&Product{}, &Category{}); err != nil {
		panic("failed to migrate database")
	}

	return db
}

func TestMinPrice(t *testing.T) {
	db := setupTestDB()

	products := []Product{
		{Name: initialProduct1Name, Price: initialProduct1Price},
		{Name: initialProduct2Name, Price: initialProduct2Price},
		{Name: initialProduct3Name, Price: initialProduct3Price},
	}
	for _, product := range products {
		db.Create(&product)
	}

	var result []Product
	db.Scopes(MinPrice(30)).Find(&result)

	assert.Equal(t, 2, len(result))
	assert.Equal(t, initialProduct2Name, result[0].Name)
	assert.Equal(t, initialProduct2Price, result[0].Price)
	assert.Equal(t, initialProduct3Name, result[1].Name)
	assert.Equal(t, initialProduct3Price, result[1].Price)

	var allProducts []Product
	db.Find(&allProducts)
	assert.Equal(t, 3, len(allProducts))
	assert.Equal(t, initialProduct1Name, allProducts[0].Name)
	assert.Equal(t, initialProduct1Price, allProducts[0].Price)
	assert.Equal(t, initialProduct2Name, allProducts[1].Name)
	assert.Equal(t, initialProduct3Name, allProducts[2].Name)
	assert.True(t, allProducts[2].Price > allProducts[1].Price)
}

func TestByCategory(t *testing.T) {
	db := setupTestDB()

	category1 := Category{Name: category1Name}
	category2 := Category{Name: category2Name}
	db.Create(&category1)
	db.Create(&category2)

	products := []Product{
		{Name: product1Name, CategoryID: category1.ID},
		{Name: product2Name, CategoryID: category2.ID},
		{Name: product3Name, CategoryID: category1.ID},
	}
	for _, product := range products {
		db.Create(&product)
	}

	var result []Product
	db.Scopes(ByCategory(category1.ID)).Find(&result)

	assert.Equal(t, 2, len(result))
	assert.Equal(t, product1Name, result[0].Name)
	assert.Equal(t, category1.ID, result[0].CategoryID)
	assert.Equal(t, product3Name, result[1].Name)
	assert.Equal(t, category1.ID, result[1].CategoryID)

	var allProducts []Product
	db.Find(&allProducts)
	assert.Equal(t, 3, len(allProducts))
	assert.Equal(t, product1Name, allProducts[0].Name)
	assert.Equal(t, product2Name, allProducts[1].Name)
	assert.Equal(t, product3Name, allProducts[2].Name)
	assert.Equal(t, category2.ID, allProducts[1].CategoryID)
}

func TestSortProducts(t *testing.T) {
	db := setupTestDB()

	products := []Product{
		{Name: sortProductName1, Price: sortProductPrice1},
		{Name: sortProductName2, Price: sortProductPrice2},
		{Name: sortProductName3, Price: sortProductPrice3},
	}
	for _, product := range products {
		db.Create(&product)
	}

	var resultAsc []Product
	db.Scopes(SortProducts("price", false)).Find(&resultAsc)

	assert.Equal(t, 3, len(resultAsc))
	assert.Equal(t, sortProductName2, resultAsc[0].Name)
	assert.Equal(t, sortProductPrice2, resultAsc[0].Price)
	assert.Equal(t, sortProductName3, resultAsc[1].Name)
	assert.Equal(t, sortProductPrice3, resultAsc[1].Price)
	assert.Equal(t, sortProductName1, resultAsc[2].Name)
	assert.Equal(t, sortProductPrice1, resultAsc[2].Price)

	assert.True(t, resultAsc[0].Price < resultAsc[1].Price)
	assert.True(t, resultAsc[1].Price < resultAsc[2].Price)

	var resultDesc []Product
	db.Scopes(SortProducts("price", true)).Find(&resultDesc)

	assert.Equal(t, 3, len(resultDesc))
	assert.Equal(t, sortProductName1, resultDesc[0].Name)
	assert.Equal(t, sortProductPrice1, resultDesc[0].Price)
	assert.Equal(t, sortProductName3, resultDesc[1].Name)
	assert.Equal(t, sortProductPrice3, resultDesc[1].Price)
	assert.Equal(t, sortProductName2, resultDesc[2].Name)
	assert.Equal(t, sortProductPrice2, resultDesc[2].Price)

	assert.True(t, resultDesc[0].Price > resultDesc[1].Price)
	assert.True(t, resultDesc[1].Price > resultDesc[2].Price)
}
