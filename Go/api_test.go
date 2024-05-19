package main

import (
	"bytes"
	"encoding/json"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go-app/database"
	"go-app/models"
)

func setupTestEcho() *echo.Echo {
	e := echo.New()
	setupTestData()
	os.Setenv("TEST_PARALLEL", "false")

	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)
	e.POST("/products", createProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.POST("/carts", createCart)
	e.GET("/carts/:id", getCart)
	e.POST("/carts/:id/items", addItem)
	e.DELETE("/carts/:id/items/:itemId", removeItem)

	e.POST("/categories", createCategory)
	e.GET("/categories", getCategories)
	e.GET("/categories/:id", getCategory)
	e.PUT("/categories/:id", updateCategory)
	e.DELETE("/categories/:id", deleteCategory)

	e.POST("/payment", processPayment)
	e.GET("/payment", getPayment)

	return e
}

func jsonPayload(v interface{}) *bytes.Buffer {
	b, _ := json.Marshal(v)
	return bytes.NewBuffer(b)
}

func setupTestData() {
	cleanUpTestData()
	database.DB, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	database.DB.AutoMigrate(&models.Product{}, &models.Category{}, &models.Cart{}, &models.CartItem{}, &models.Payment{})

	database.DB.Create(&models.Product{Model: gorm.Model{ID: 1}, Name: "Initial Product", Description: "Initial Description", Price: 50.0})
	database.DB.Create(&models.Product{Name: "Initial Product2", Description: "Initial Description2", Price: 50.0})
	database.DB.Create(&models.Category{Name: "Initial Category", Description: "Initial Description"})
	database.DB.Create(&models.Cart{UserID: 1})
}

func cleanUpTestData() {
	err := os.Remove("test.db")
	if err != nil {
		return
	}
}

func TestUpdateProduct(t *testing.T) {
	e := setupTestEcho()
	setupTestData()

	// Positive test case: Update product
	product := models.Product{Name: "Updated Product", Description: "Updated Description", Price: 150.0}
	req := httptest.NewRequest(http.MethodPut, "/products/1", jsonPayload(product))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	updateProduct(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Updated Product")

	// Negative test case: Non-existent product
	req = httptest.NewRequest(http.MethodPut, "/products/99999", jsonPayload(product))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("99999")

	updateProduct(c)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestDeleteProduct(t *testing.T) {
	e := setupTestEcho()
	setupTestData()

	// Positive test case: Delete existing product
	req := httptest.NewRequest(http.MethodDelete, "/products/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, deleteProduct(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Product deleted")
	}

	// Negative test case: Delete non-existent product
	req = httptest.NewRequest(http.MethodDelete, "/products/9991", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("9991")

	deleteProduct(c)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Contains(t, rec.Body.String(), "Product not found")
}

func TestGetProducts(t *testing.T) {
	e := setupTestEcho()
	setupTestData()

	// Positive test case: Get products
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, getProducts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Initial Product")
	}

	// Negative test case: Incorrect endpoint
	req = httptest.NewRequest(http.MethodGet, "/wrongendpoint", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestCreateProduct(t *testing.T) {
	e := setupTestEcho()
	setupTestData()

	// Positive test case: Create product
	product := models.Product{Name: "Test Product", Description: "Test Description", Price: 100.0}
	req := httptest.NewRequest(http.MethodPost, "/products", jsonPayload(product))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, createProduct(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "Test Product")
	}

	// Negative test case: Invalid payload
	req = httptest.NewRequest(http.MethodPost, "/products", bytes.NewBufferString("{invalid json}"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	// Here, we don't expect a nil error, but the status code should be 400
	createProduct(c)
	assert.Contains(t, rec.Body.String(), "Invalid JSON payload")
}

func TestCreateCart(t *testing.T) {
	e := setupTestEcho()
	setupTestData()

	// Positive test case: Create cart
	cart := models.Cart{UserID: 2}
	req := httptest.NewRequest(http.MethodPost, "/carts", jsonPayload(cart))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, createCart(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), `"UserID":2`)
	}

	// Negative test case: Invalid payload
	req = httptest.NewRequest(http.MethodPost, "/carts", bytes.NewBufferString("{invalid json}"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	createCart(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

}

func TestGetCart(t *testing.T) {
	e := setupTestEcho()
	setupTestData()

	// Positive test case: Get cart
	req := httptest.NewRequest(http.MethodGet, "/carts/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, getCart(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), `"UserID":1`)
	}

	// Negative test case: Non-existent cart
	req = httptest.NewRequest(http.MethodGet, "/carts/9992", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("9992")

	getCart(c)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)

}

func TestAddItem(t *testing.T) {
	e := setupTestEcho()
	setupTestData()

	// Positive test case: Add item
	item := models.CartItem{ProductID: 1, Quantity: 1}
	req := httptest.NewRequest(http.MethodPost, "/carts/1/items", jsonPayload(item))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, addItem(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), `"Quantity":1`)
	}

	// Negative test case: Invalid payload
	req = httptest.NewRequest(http.MethodPost, "/carts/1/items", bytes.NewBufferString("{invalid json}"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	addItem(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

}

func TestRemoveItem(t *testing.T) {
	e := setupTestEcho()
	setupTestData()

	// Positive test case: Remove item
	req := httptest.NewRequest(http.MethodDelete, "/carts/1/items/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id", "itemId")
	c.SetParamValues("1", "1")

	if assert.NoError(t, removeItem(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Item removed")
	}

	// Negative test case: Non-existent item
	req = httptest.NewRequest(http.MethodDelete, "/carts/1/items/9993", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id", "itemId")
	c.SetParamValues("1", "asdas")

	removeItem(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

}

func TestCreateCategory(t *testing.T) {
	e := setupTestEcho()
	setupTestData()

	// Positive test case: Create category
	category := models.Category{Name: "Test Category", Description: "Test Description"}
	req := httptest.NewRequest(http.MethodPost, "/categories", jsonPayload(category))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, createCategory(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "Test Category")
	}

	// Negative test case: Invalid payload
	req = httptest.NewRequest(http.MethodPost, "/categories", bytes.NewBufferString("{invalid json}"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	createCategory(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

}

func TestGetCategories(t *testing.T) {
	e := setupTestEcho()
	setupTestData()

	// Positive test case: Get categories
	req := httptest.NewRequest(http.MethodGet, "/categories", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, getCategories(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Initial Category")
	}

	// Negative test case: Incorrect endpoint
	req = httptest.NewRequest(http.MethodGet, "/wrongendpoint", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNotFound, rec.Code)

}

func TestUpdateCategory(t *testing.T) {
	e := setupTestEcho()
	setupTestData()

	// Positive test case: Update category
	category := models.Category{Name: "Updated Category", Description: "Updated Description"}
	req := httptest.NewRequest(http.MethodPut, "/categories/1", jsonPayload(category))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, updateCategory(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Updated Category")
	}

	req = httptest.NewRequest(http.MethodPut, "/categories/9994", bytes.NewBufferString("{invalid json}"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("9994")

	updateCategory(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

}

func TestDeleteCategory(t *testing.T) {
	e := setupTestEcho()
	setupTestData()

	// Positive test case: Delete category
	req := httptest.NewRequest(http.MethodDelete, "/categories/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, deleteCategory(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Category deleted")
	}

	// Negative test case: Non-existent category
	req = httptest.NewRequest(http.MethodDelete, "/categories/9995", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("ab")

	deleteCategory(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

}

func TestProcessPayment(t *testing.T) {
	e := setupTestEcho()
	setupTestData()

	// Positive test case: Process payment
	payment := models.Payment{
		TransactionID: "test-transaction",
		Amount:        100.0,
		Currency:      "USD",
		Status:        "Pending",
		PaymentDate:   time.Now(),
	}
	req := httptest.NewRequest(http.MethodPost, "/payment", jsonPayload(payment))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, processPayment(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Payment processed successfully")
	}

	// Negative test case: Invalid payload
	req = httptest.NewRequest(http.MethodPost, "/payment", bytes.NewBufferString("{invalid json}"))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	processPayment(c)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

}

func TestGetPayment(t *testing.T) {
	e := setupTestEcho()
	setupTestData()

	// Positive test case: Get payment
	req := httptest.NewRequest(http.MethodGet, "/payment", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, getPayment(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), `"TransactionID":"test-transaction"`)
	}

	// Negative test case: Incorrect endpoint
	req = httptest.NewRequest(http.MethodGet, "/wrongendpoint", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	e.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNotFound, rec.Code)

}
