package repository

import (
	"github.com/sakupay-apps/internal/model"
	"gorm.io/gorm"
)

type BillRepository interface {
	Create(payload *model.Bill) (*model.Bill,error)
}

type billRepository struct {
	db *gorm.DB
}

func (b *billRepository) Create(payload *model.Bill) (*model.Bill,error) {
	bill := model.Bill {
		ID : payload.ID,
		UserID: payload.UserID,
		Billdetails: payload.Billdetails,
		Total: payload.Total,
		DueDate: payload.DueDate,

	}

	if err := b.db.Create(&bill).Error; err != nil {
		return nil,err
	}
	return  &bill,nil
}

func (b *billRepository) Get(id string) (*model.Bill, error) {
	bill := model.Bill{}

	if err := b.db.Where("WHERE id = ? ", id).Find(&bill).Error; err != nil {
		return nil, err
	}

	return &bill, nil
}


func NewBillRepository(db *gorm.DB) BillRepository {
	return &billRepository{
		db: db,
	}
}
