package repository

import (
	"github.com/sakupay-apps/internal/model"
	"gorm.io/gorm"
)

type WalletRepository interface {
	Create(payload model.Wallet) error
}

type walletRepository struct {
	db *gorm.DB
}

func (w *walletRepository) Create(payload model.Wallet) error {
	var wallet = model.Wallet{ID: payload.ID, UserID: payload.UserID, Name: payload.Name, Balance: payload.Balance}
	if err := w.db.Create(&wallet).Error; err != nil {
		return err
	}
	return nil
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepository{
		db: db,
	}
}
