package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-app/database"
	"go-app/models"
	"net/http"
	"strconv"
)

func main() {
	database.ConnectDataBase()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)
	e.POST("/products", createProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.Logger.Fatal(e.Start(":8080"))
}

func getProducts(c echo.Context) error {
	var products []models.Product
	result := database.DB.Find(&products)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, products)
}

func getProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product := new(models.Product)
	result := database.DB.First(&product, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, product)
}

func createProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return err
	}

	result := database.DB.Create(&product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}

	return c.JSON(http.StatusOK, product)
}

func updateProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	product := new(models.Product)
	if err := c.Bind(product); err != nil {
		return err
	}
	product.ID = uint(id)
	result := database.DB.Save(&product)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, product)
}

func deleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	result := database.DB.Delete(&models.Product{}, id)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Product deleted"})
}
