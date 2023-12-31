package repository

import (
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/utils/constants"
	"gorm.io/gorm"
)

type BillDetailsRepository interface {
	Get(id string) (*model.BillDetails, error)
}

type billDetailsRepository struct {
	db *gorm.DB
}

func NewBillDetailsRepository(db *gorm.DB) BillDetailsRepository {
	return &billDetailsRepository{
		db: db,
	}
}

func (b *billDetailsRepository) Get(id string) (*model.BillDetails, error) {
	billDetails := model.BillDetails{}

	if err := b.db.Where(constants.WHERE_BY_BILL_ID, id).First(&billDetails).Error; err != nil {
		return nil, err
	}

	return &billDetails, nil
}
