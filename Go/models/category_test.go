package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCategoryCreation(t *testing.T) {
	category := Category{
		Name:        "Test Category",
		Description: "Test Description",
	}

	assert.Equal(t, "Test Category", category.Name)
	assert.Equal(t, "Test Description", category.Description)
}
