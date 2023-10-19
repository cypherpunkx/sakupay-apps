package repository

import (
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/utils/constants"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreateDepositTransaction(payload *model.Transaction, cardID string) (*model.Transaction, error)
	Get(id string) (*model.Transaction, error)
	List() ([]*model.Transaction, error)
	ListTransactions(id string) ([]*model.Transaction, error)
	GetTransaction(userID, transactionID string) (*model.Transaction, error)
	CreateSendTransaction(payload *model.Transaction, friendID string) (*model.Transaction, error)
	CreateWitdrawTransaction(payload *model.Transaction, cardID string) (*model.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) CreateDepositTransaction(payload *model.Transaction, cardID string) (*model.Transaction, error) {
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
		card := model.Card{}

		if transaction.TransactionType == "deposit" {
			if err := tx.Model(&wallet).Where(constants.WHERE_BY_USER_ID, transaction.UserID).Select("balance").First(&wallet).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			wallet.Balance += transaction.Amount

			if err := tx.Model(&wallet).Where(constants.WHERE_BY_USER_ID, transaction.UserID).Select("balance").Updates(&wallet).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			if err := tx.Model(&card).Where(constants.WHERE_BY_USER_ID_AND_CARD_ID, transaction.UserID, cardID).Select("balance").First(&card).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			card.Balance -= transaction.Amount

			if err := tx.Model(&card).Where(constants.WHERE_BY_USER_ID_AND_CARD_ID, transaction.UserID, cardID).Select("balance").Updates(&card).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			if err := tx.Create(&transaction).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

		} else {
			return gorm.ErrInvalidTransaction
		}

		return nil
	})

	return &transaction, nil
}

func (r *transactionRepository) Get(id string) (*model.Transaction, error) {
	transaction := model.Transaction{}

	if err := r.db.Where(constants.WHERE_BY_ID, id).First(&transaction).Error; err != nil {
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

func (r *transactionRepository) CreateSendTransaction(payload *model.Transaction, friendID string) (*model.Transaction, error) {
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

		if transaction.TransactionType == "send" {
			if err := tx.Model(&wallet).Where(constants.WHERE_BY_USER_ID, transaction.UserID).Select("balance").First(&wallet).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			wallet.Balance -= transaction.Amount

			if err := tx.Model(&wallet).Where(constants.WHERE_BY_USER_ID, transaction.UserID).Select("balance").Updates(&wallet).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			if err := tx.Model(&wallet).Where(constants.WHERE_BY_USER_ID, friendID).Select("balance").First(&wallet).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			wallet.Balance += transaction.Amount

			if err := tx.Model(&wallet).Where(constants.WHERE_BY_USER_ID, friendID).Select("balance").Updates(&wallet).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			if err := tx.Create(&transaction).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}
		} else {
			return gorm.ErrInvalidTransaction
		}
		return nil
	})

	return &transaction, nil
}

func (r *transactionRepository) CreateWitdrawTransaction(payload *model.Transaction, cardID string) (*model.Transaction, error) {
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
		card := model.Card{}

		if transaction.TransactionType == "withdrawal" {
			if err := tx.Model(&wallet).Where(constants.WHERE_BY_USER_ID, transaction.UserID).Select("balance").First(&wallet).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			wallet.Balance -= transaction.Amount

			if err := tx.Model(&wallet).Where(constants.WHERE_BY_USER_ID, transaction.UserID).Select("balance").Updates(&wallet).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			if err := tx.Model(&card).Where(constants.WHERE_BY_USER_ID_AND_CARD_ID, transaction.UserID, cardID).Select("balance").First(&card).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			card.Balance += transaction.Amount

			if err := tx.Model(&card).Where(constants.WHERE_BY_USER_ID_AND_CARD_ID, transaction.UserID, cardID).Select("balance").Updates(&card).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			if err := tx.Create(&transaction).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}
		} else {
			return gorm.ErrInvalidTransaction
		}

		return nil
	})

	return &transaction, nil
}
