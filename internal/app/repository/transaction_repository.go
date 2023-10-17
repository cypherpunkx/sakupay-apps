package repository

import (
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/utils/constants"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(payload *model.Transaction) (*model.Transaction, error)
	Get(id string) (*model.Transaction, error)
	List() ([]*model.Transaction, error)
	ListTransactions(id string) ([]*model.Transaction, error)
	GetTransaction(userID, transactionID string) (*model.Transaction, error)
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

			if err := tx.Model(&wallet).Where(constants.WHERE_BY_USER_ID, transaction.UserID).Select("balance").First(&wallet).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			wallet.Balance += transaction.Amount

			if err := tx.Model(&wallet).Where(constants.WHERE_BY_USER_ID, transaction.UserID).Select("balance").Updates(&wallet).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			return nil
		}

		if transaction.TransactionType == "send" {

			if err := tx.Create(&transaction).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			if err := tx.Model(&wallet).Where(constants.WHERE_BY_USER_ID, transaction.UserID).Select("balance").First(&wallet).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			wallet.Balance -= transaction.Amount

			if err := tx.Model(&wallet).Where(constants.WHERE_BY_USER_ID, transaction.UserID).Select("balance").Updates(&wallet).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			return nil
		}
		return nil
	})

	return &transaction, nil
}

func (r *transactionRepository) Get(id string) (*model.Transaction, error) {
	transaction := model.Transaction{}

	if err := r.db.Where(constants.WHERE_BY_ID, id).Preload("User").First(&transaction).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return &transaction, nil
}

func (r *transactionRepository) List() ([]*model.Transaction, error) {
	transactions := []*model.Transaction{}

	if err := r.db.Find(&transactions).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return transactions, nil
}

func (r *transactionRepository) ListTransactions(id string) ([]*model.Transaction, error) {
	transactions := []*model.Transaction{}

	if err := r.db.Where(constants.WHERE_BY_USER_ID, id).Find(&transactions).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return transactions, nil
}

func (r *transactionRepository) GetTransaction(userID, transactionID string) (*model.Transaction, error) {
	transaction := model.Transaction{}

	if err := r.db.Where("user_id = ? AND id = ?", userID, transactionID).First(&transaction).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return &transaction, nil
}
