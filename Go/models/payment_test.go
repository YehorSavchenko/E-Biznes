package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPaymentCreation(t *testing.T) {
	payment := Payment{
		TransactionID: "12345",
		Amount:        100.0,
		Currency:      "USD",
		Status:        "Pending",
		PaymentDate:   time.Now(),
	}

	assert.Equal(t, "12345", payment.TransactionID)
	assert.Equal(t, 100.0, payment.Amount)
	assert.Equal(t, "USD", payment.Currency)
	assert.Equal(t, "Pending", payment.Status)
}
