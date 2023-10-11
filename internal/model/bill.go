package model

import "time"

type Bill struct {
	ID            string    `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	UserID        string    `gorm:"type:uuid;not null;unique" json:"userID" binding:"required"`
	BilldetailsID string    `gorm:"type:uuid;not null;unique" json:"billdetailsID" binding:"required"`
	Total         float64   `gorm:"type:float;not null;default:0" json:"total" binding:"number"`
	DueDate       time.Time `gorm:"type:timestamp;not null" json:"dueDate" binding:"datetime"`
}
