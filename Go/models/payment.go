package models

import (
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	gorm.Model
	TransactionID string
	Amount        float64
	Currency      string
	Status        string
	PaymentDate   time.Time
}
