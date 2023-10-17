package model

type Contact struct {
	ID           string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	UserID       string `gorm:"type:uuid;not null" json:"userID" binding:"required"`
	PhoneNumber  string `gorm:"type:varchar(255);not null;unique" json:"phoneNumber" binding:"required,numeric"`
	Relationship string `gorm:"type:varchar(255)" json:"relationship" binding:"alpha"`
	IsFavorite   bool   `gorm:"type:bool;default:false" json:"isFavorite" binding:"oneof=true false"`
}
