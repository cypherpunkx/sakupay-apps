package dto

import (
	"time"

	"github.com/sakupay-apps/internal/model"
)

type UserResponse struct {
	ID               string          `json:"id"`
	Username         string          `json:"username"`
	Email            string          `json:"email"`
	Password         string          `json:"password"`
	FirstName        string          `json:"firstName"`
	LastName         string          `json:"lastName"`
	PhoneNumber      string          `json:"phoneNumber"`
	RegistrationDate time.Time       `json:"registrationDate"`
	ProfilePicture   []byte          `json:"profilePicture"`
	LastLogin        time.Time       `json:"lastLogin"`
	Wallet           model.Wallet    `json:"wallet,omitempty"`
	Cards            []model.Card    `json:"cards,omitempty"`
	Contacts         []model.Contact `json:"contacts,omitempty"`
}
