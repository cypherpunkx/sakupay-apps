package dto

import "github.com/sakupay-apps/internal/model"

type WalletResponse struct {
	ID      string     `json:"id"`
	User    model.User `json:"user"`
	Name    string     `json:"name"`
	Balance float64    `json:"balance"`
}
