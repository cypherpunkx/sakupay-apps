package model

type Wallet struct {
	ID      string  `gorm:"type:uuid;primaryKey;not null;unique" json:"id"`
	UserID  string  `gorm:"type:uuid;not null;unique;references:ID" json:"userID"`
	Name    string  `gorm:"type:varchar(255);not null;default:sakupay;" json:"name"`
	Balance float64 `gorm:"type:float;not null;default:0" json:"balance"`
}
