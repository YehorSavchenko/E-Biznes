package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCartCreation(t *testing.T) {
	cart := Cart{
		UserID: 1,
		Items:  []CartItem{},
	}

	assert.Equal(t, uint(1), cart.UserID)
	assert.Equal(t, 0, len(cart.Items))
}

func TestAddItemToCart(t *testing.T) {
	cart := Cart{
		UserID: 1,
		Items:  []CartItem{},
	}

	product := Product{
		Name:        "Test Product",
		Description: "Test Description",
		Price:       100.0,
	}

	cartItem := CartItem{
		ProductID: product.ID,
		Product:   product,
		Quantity:  1,
	}

	cart.Items = append(cart.Items, cartItem)
	assert.Equal(t, 1, len(cart.Items))
	assert.Equal(t, product.Name, cart.Items[0].Product.Name)
	assert.Equal(t, 1, cart.Items[0].Quantity)
}

func TestRemoveItemFromCart(t *testing.T) {
	cart := Cart{
		UserID: 1,
		Items: []CartItem{
			{ProductID: 1, Quantity: 1},
			{ProductID: 2, Quantity: 1},
		},
	}

	cart.Items = cart.Items[1:]
	assert.Equal(t, 1, len(cart.Items))
	assert.Equal(t, uint(2), cart.Items[0].ProductID)
}
