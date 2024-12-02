package models

import (
	"time"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Amount      float64   `json:"amount" binding:"required"`
	Category    string    `json:"category" binding:"required"`
	Description string    `json:"description"`
	Type        string    `json:"type" binding:"required,oneof=income expense"` // income or expense
	Date        time.Time `json:"date" binding:"required"`
}
