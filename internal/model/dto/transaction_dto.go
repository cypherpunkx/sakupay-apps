package dto

import (
	"time"

	"github.com/sakupay-apps/internal/model"
)

type TransactionResponse struct {
	ID              string     `json:"id"`
	User            model.User `json:"user"`
	TransactionType string     `json:"transactionType"`
	Amount          float64    `json:"amount"`
	Description     string     `json:"description"`
	Timestamp       time.Time  `json:"timestamp"`
}
