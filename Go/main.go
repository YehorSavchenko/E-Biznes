package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-app/database"
	"go-app/models"
	"net/http"
	"strconv"
	"time"
)

func main() {
	database.ConnectDataBase()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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

	e.Logger.Fatal(e.Start(":8080"))
}

func getProducts(c echo.Context) error {
	var products []models.Product
	minPriceParam := c.QueryParam("minPrice")
	categoryParam := c.QueryParam("category")
	sortField := c.QueryParam("sortField")
	sortDesc := c.QueryParam("sortDesc") == "true"

	query := database.DB.Model(&models.Product{}).Preload("Category")

	if minPriceParam != "" {
		if minPrice, err := strconv.ParseFloat(minPriceParam, 64); err == nil {
			query = query.Scopes(models.MinPrice(minPrice))
		}
	}
	if categoryParam != "" {
		if categoryID, err := strconv.Atoi(categoryParam); err == nil {
			query = query.Scopes(models.ByCategory(uint(categoryID)))
		}
	}
	if sortField != "" {
		query = query.Scopes(models.SortProducts(sortField, sortDesc))
	}

	result := query.Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}
	return c.JSON(http.StatusOK, products)
}

func getProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product := new(models.Product)
	result := database.DB.Preload("Category").First(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, product)
}

func createProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid JSON payload"})
	}

	if product.CategoryID != 0 {
		var category models.Category
		if err := database.DB.First(&category, product.CategoryID).Error; err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Category not found"})
		}
	}

	result := database.DB.Preload("Category").Create(&product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusCreated, product)
}

func updateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request payload"})
	}
	product.ID = uint(id)

	var existingProduct models.Product
	if err := database.DB.Preload("Category").First(&existingProduct, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Product not found"})
	}

	if product.CategoryID != 0 {
		var category models.Category
		if err := database.DB.First(&category, product.CategoryID).Error; err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Category not found"})
		}
	}

	result := database.DB.Preload("Category").Save(&product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}
	return c.JSON(http.StatusOK, product)
}

func deleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Product not found"})
	}

	result := database.DB.Delete(&product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Product deleted"})
}

func createCart(c echo.Context) error {
	var cart = new(models.Cart)
	if err := c.Bind(cart); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid cart data"})
	}

	result := database.DB.Create(cart)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusCreated, cart)
}

func getCart(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid cart ID"})
	}

	cart := new(models.Cart)
	result := database.DB.Preload("Items").Preload("Items.Product").First(cart, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusOK, cart)
}

func addItem(c echo.Context) error {
	cartID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid cart ID"})
	}

	item := new(models.CartItem)
	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid item data"})
	}
	item.CartID = uint(cartID)

	result := database.DB.Create(item)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusCreated, item)
}

func removeItem(c echo.Context) error {
	cartID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid cart ID"})
	}

	itemID, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid item ID"})
	}

	result := database.DB.Where("cart_id = ?", cartID).Delete(&models.CartItem{}, itemID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Item removed"})
}

func createCategory(c echo.Context) error {
	var category models.Category
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid category data"})
	}
	result := database.DB.Create(&category)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}
	return c.JSON(http.StatusCreated, category)
}

func getCategories(c echo.Context) error {
	var categories []models.Category
	result := database.DB.Find(&categories)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}
	return c.JSON(http.StatusOK, categories)
}

func getCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var category models.Category
	result := database.DB.First(&category, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}
	return c.JSON(http.StatusOK, category)
}

func updateCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var category models.Category
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid category data"})
	}
	category.ID = uint(id)
	result := database.DB.Save(&category)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}
	return c.JSON(http.StatusOK, category)
}

func deleteCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid Category ID"})
	}
	result := database.DB.Delete(&models.Category{}, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Category deleted"})
}

func processPayment(c echo.Context) error {
	var payment models.Payment

	if err := c.Bind(&payment); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid payment data"})
	}

	payment.Status = "Completed"
	payment.PaymentDate = time.Now()

	result := database.DB.Create(&payment)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Payment processed successfully", "paymentID": payment.ID})
}

func getPayment(c echo.Context) error {
	var payment []models.Payment
	result := database.DB.Find(&payment)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}
	return c.JSON(http.StatusOK, payment)
}
