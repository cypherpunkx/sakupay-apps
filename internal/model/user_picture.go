package model

type UserPicture struct {
	ID           string `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	UserID       string `gorm:"type:uuid;not null:references:ID" json:"userID" binding:"required"`
	FileLocation string `gorm:"type:varchar(255);not null" json:"fileLocation" binding:"required"`
}
