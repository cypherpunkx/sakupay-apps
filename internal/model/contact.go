package model

type Contact struct {
	ID           string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	UserID       string `gorm:"type:uuid;not null;unique" json:"userID" binding:"required"`
	PhoneNumber  string `gorm:"type:varchar(255);not null;unique" json:"phoneNumber" binding:"required,numeric,unique"`
	Relationship string `gorm:"type:varchar(255)" json:"relationship" binding:"required,alpha"`
	IsFavorite   bool   `gorm:"type:bool;default:false" json:"isFavorite" binding:"oneof=true false,boolean"`
}
