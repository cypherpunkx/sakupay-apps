package model

type Contact struct {
	ID           string `gorm:"type:uuid;primaryKey;not null" json:"id"`
	UserID       string `gorm:"type:uuid;not null" json:"userID"`
	PhoneNumber  string `gorm:"type:varchar(255);not null" json:"phoneNumber"`
	Relationship string `gorm:"type:varchar(255)" json:"relationship"`
	IsFavorite   bool   `gorm:"type:bool;default:false" json:"isFavorite"`
}
