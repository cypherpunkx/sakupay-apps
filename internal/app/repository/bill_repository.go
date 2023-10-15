package repository

import (
	"github.com/sakupay-apps/internal/model"
	"gorm.io/gorm"
)

type BillRepository interface {
	Create(payload *model.Bill) (*model.Bill, error)
}

type billRepository struct {
	db *gorm.DB
}

func NewBillRepository(db *gorm.DB) BillRepository {
	return &billRepository{db: db}
}

func (r *billRepository) Create(payload *model.Bill) (*model.Bill, error) {
	bill := model.Bill{
		ID:            payload.ID,
		UserID:        payload.UserID,
		BilldetailsID: payload.BilldetailsID,
		Total:         payload.Total,
		DueDate:       payload.DueDate,
	}

	r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&bill).Error; err != nil {
			return err
		}
		return nil
	})

	return &bill, nil
}
