package repository

import (
	"github.com/sakupay-apps/internal/model"
	"gorm.io/gorm"
)

type BillDetailsRepository interface {
	Get(id string) (*model.BillDetails, error)
}

type billDetailsRepository struct {
	db *gorm.DB
}

func (b *billRepository) GetBillDetailsById(id string) (*model.BillDetails, error) {
	billDetails := model.BillDetails{}

	if err := b.db.Where("WHERE user_id = ? ", id).Find(&billDetails).Error; err != nil {
		return nil, err
	}

	return &billDetails, nil
}

func NewBillDetailsRepository(db *gorm.DB) BillRepository {
	return &billRepository{
		db: db,
	}
}
