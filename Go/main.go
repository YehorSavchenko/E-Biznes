package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/products", getProducts)
	e.GET("/products/:id", getProduct)
	e.POST("/products", createProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.Logger.Fatal(e.Start(":8080"))
}

func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get Products")
}

func getProduct(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get Product")
}

func createProduct(c echo.Context) error {
	return c.JSON(http.StatusOK, "Product Created")
}

func updateProduct(c echo.Context) error {
	return c.JSON(http.StatusOK, "Product Updated")
}

func deleteProduct(c echo.Context) error {
	return c.JSON(http.StatusOK, "Product Deleted")
}
