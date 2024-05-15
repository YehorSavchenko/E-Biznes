package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductCreation(t *testing.T) {
	product := Product{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       100.0,
	}

	assert.Equal(t, "Test Product", product.Name)
	assert.Equal(t, "Test Description", product.Description)
	assert.Equal(t, 100.0, product.Price)
}
