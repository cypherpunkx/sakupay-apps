package model

type Wallet struct {
	ID      string  `gorm:"type:uuid;primaryKey;not null;unique" json:"id,omitempty"`
	UserID  string  `gorm:"type:uuid;not null;unique;references:ID" json:"userID,omitempty"`
	Name    string  `gorm:"type:varchar(255);not null;default:sakupay;" json:"name"`
	Balance float64 `gorm:"type:float;not null;default:0;check:Balance >= 0" json:"balance"`
}
