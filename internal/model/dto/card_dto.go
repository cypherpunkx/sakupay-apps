package dto

import (
	"time"

	"github.com/sakupay-apps/internal/model"
)

type CardResponse struct {
	ID             string     `json:"id"`
	User           model.User `json:"user"`
	CardNumber     string     `json:"cardNumber"`
	CardholderName string     `json:"cardholderName"`
	Balance        float64    `json:"balance"`
	ExpirationDate time.Time  `json:"expirationDate"`
	CVV            string     `json:"cvv"`
}
