package repository

import (
	"github.com/sakupay-apps/internal/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(payload *model.Transaction) (*model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(payload *model.Transaction) (*model.Transaction, error) {
	transaction := model.Transaction{
		ID:              payload.ID,
		UserID:          payload.UserID,
		TransactionType: payload.TransactionType,
		Amount:          payload.Amount,
		Description:     payload.Description,
		Timestamp:       payload.Timestamp,
	}

	r.db.Transaction(func(tx *gorm.DB) error {
		wallet := model.Wallet{}
		if transaction.TransactionType == "deposit" {
			if err := tx.Create(&transaction).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			if err := tx.Model(&wallet).Where("user_id = ?", transaction.UserID).Select("balance").First(&wallet).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			wallet.Balance += transaction.Amount

			if err := tx.Model(&wallet).Where("user_id = ?", transaction.UserID).Select("balance").Updates(&wallet).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			return nil
		}
		return nil
	})

	return &transaction, nil
}
