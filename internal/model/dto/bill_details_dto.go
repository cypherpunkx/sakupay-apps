package dto

import "github.com/sakupay-apps/internal/model"


type BillDetailsResponse struct {
	ID          string     `json:"id"`
	Bill        model.Bill `json:"bill"`
	Name        string     `json:"name"`
	Category    string     `json:"category"`
	Description string     `json:"description"`
	Website     string     `json:"website"`
}

