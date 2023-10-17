package dto

import (
	"time"

	"github.com/sakupay-apps/internal/model"
)

type BillResponse struct {
	ID          string            `json:"id"`
	User        model.User        `json:"user"`
	BillDetails model.BillDetails `json:"billDetails"`
	Total       float64           `json:"total"`
	DueDate     time.Time         `json:"dueDate"`
}
