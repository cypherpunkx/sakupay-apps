package model

import "time"

type Bill struct {
	ID          string      `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	UserID      string      `gorm:"type:uuid;not null:references:ID" json:"userID" binding:"required"`
	Billdetails BillDetails `gorm:"foreignKey:BillID" json:"billDetails"`
	Total       float64     `gorm:"type:float;not null;default:0;check:Total >= 0" json:"total" binding:"number"`
	DueDate     time.Time   `gorm:"type:timestamp;not null" json:"dueDate"`
	Status      string      `gorm:"type:varchar(255);not null;default:pending" json:"status" binding:"required,alpha,oneof=pending cancel paid"`
	Notified    bool        `gorm:"type:bool;default:false" json:"notified"`
}
