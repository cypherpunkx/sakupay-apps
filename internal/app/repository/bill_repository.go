package repository

import (
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/utils/constants"
	"gorm.io/gorm"
)

type BillRepository interface {
	Create(payload *model.Bill) (*model.Bill, error)
	List(id string) ([]*model.Bill, error)
}

type billRepository struct {
	db *gorm.DB
}

func NewBillRepository(db *gorm.DB) BillRepository {
	return &billRepository{
		db: db,
	}
}

func (b *billRepository) Create(payload *model.Bill) (*model.Bill, error) {
	bill := model.Bill{
		ID:          payload.ID,
		UserID:      payload.UserID,
		Billdetails: payload.Billdetails,
		Total:       payload.Total,
		DueDate:     payload.DueDate,
		Status:      payload.Status,
		Notified:    payload.Notified,
	}

	if err := b.db.Create(&bill).Error; err != nil {
		return nil, err
	}
	return &bill, nil
}

func (b *billRepository) List(id string) ([]*model.Bill, error) {
	bills := []*model.Bill{}

	if err := b.db.Model(&model.Bill{}).Where(constants.WHERE_BY_USER_ID, id).Preload("BillDetails").Find(&bills).Error; err != nil {
		return nil, err
	}
	return bills, nil
}
