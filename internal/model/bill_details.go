package model

type BillDetails struct {
	ID          string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	BillID      string `gorm:"type:uuid;not null;unique" json:"billID" binding:"required"`
	Name        string `gorm:"type:varchar(255);not null" json:"name" binding:"required,alpha"`
	Category    string `gorm:"type:varchar(255);not null" json:"category" binding:"required,alpha"`
	Description string `gorm:"type:text" json:"description" binding:"len=100"`
	Website     string `gorm:"type:varchar(100);not null;" json:"website" binding:"required,url"`
}