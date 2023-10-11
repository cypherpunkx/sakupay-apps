package dto

import "github.com/sakupay-apps/internal/model"

type ContactResponse struct {
	ID           string     `json:"id"`
	User         model.User `json:"user"`
	PhoneNumber  string     `json:"phoneNumber"`
	Relationship string     `json:"relationship"`
	IsFavorite   bool       `json:"isFavorite"`
}
