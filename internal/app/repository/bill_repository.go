package repository

import (
	"fmt"

	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/utils/constants"
	"gorm.io/gorm"
)

type BillRepository interface {
	Create(payload *model.Bill) (*model.Bill, error)
	List(id string) ([]*model.Bill, error)
	Get(id string) (*model.Bill, error)
	// GetByUserID(id string) ([]*model.Bill, error)
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

	b.db.Transaction(func(tx *gorm.DB) error {
		wallet := model.Wallet{}
		if err := tx.Create(&bill).Error; err != nil {
			return err
		}

		if err := tx.Model(&wallet).Where(constants.WHERE_BY_USER_ID, bill.UserID).Select("balance").First(&wallet).Error; err != nil {
			return gorm.ErrInvalidTransaction
		}

		wallet.Balance -= bill.Total

		if err := tx.Model(&wallet).Where(constants.WHERE_BY_USER_ID, bill.UserID).Select("balance").Updates(&wallet).Error; err != nil {
			return gorm.ErrInvalidTransaction
		}
		return nil
	})

	return &bill, nil
}

func (b *billRepository) List(id string) ([]*model.Bill, error) {
	bills := []*model.Bill{}

	if err := b.db.Model(&model.Bill{}).Where(constants.WHERE_BY_USER_ID, id).Find(&bills).Error; err != nil {
		return nil, err
	}

	fmt.Println(bills)

	return bills, nil
}

func (b *billRepository) Get(id string) (*model.Bill, error) {
	bill := model.Bill{}

	if err := b.db.Where(constants.WHERE_BY_USER_ID, id).First(&bill).Error; err != nil {
		return nil, err
	}

	return &bill, nil
}

// func (b *billRepository) GetByUserID(id string) ([]*model.Bill, error) {
// 	bills := []*model.Bill{}

// 	if err := b.db.Model(&model.Bill{}).Where(constants.WHERE_BY_USER_ID, id).Preload("Bil").Find(&bills).Error; err != nil {
// 		return nil, err
// 	}
// 	return bills, nil
// }
