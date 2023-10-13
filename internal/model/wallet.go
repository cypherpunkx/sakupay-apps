package model

type Wallet struct {
	ID      string  `gorm:"type:uuid;primaryKey;not null;unique" json:"id" binding:"required"`
	UserID  string  `gorm:"type:uuid;not null;unique;references:ID" json:"userID" binding:"required"`
	Name    string  `gorm:"type:varchar(255);not null;unique" json:"name" binding:"required,alpha,eq=sakupay"`
	Balance float64 `gorm:"type:float;not null;default:0" json:"balance" binding:"required,number"`
}
