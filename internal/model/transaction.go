package model

import "time"

type Transaction struct {
	ID              string    `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	UserID          string    `gorm:"type:uuid;not null;references:ID" json:"userID" binding:"required"`
	TransactionType string    `gorm:"type:varchar(255);not null" json:"transactionType" binding:"required,alpha"`
	Amount          float64   `gorm:"type:float;not null;default:0" json:"amount" binding:"required,numeric"`
	Description     string    `gorm:"type:text" json:"description" binding:"omitempty"`
	Timestamp       time.Time `gorm:"type:timestamp;default:current_timestamp" json:"timestamp"`
}
