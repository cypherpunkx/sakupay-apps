package model

import "time"

type Card struct {
	ID             string    `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	UserID         string    `gorm:"type:uuid;not null;unique;references:ID" json:"userID" binding:"required"`
	CardNumber     string    `gorm:"type:varchar(255);size:16;not null;unique" json:"cardNumber" binding:"required,numeric"`
	CardholderName string    `gorm:"type:varchar(255);not null;unique" json:"cardholderName" binding:"required,alpha"`
	ExpirationDate time.Time `gorm:"type:timestamp;not null" json:"expirationDate" binding:"required"`
	Balance        float64   `gorm:"type:float;not null;default:0;check:Balance >= 0" json:"balance"`
	CVV            string    `gorm:"type:varchar(255);size:3" json:"cvv" binding:"required,numeric,len=3"`
}
