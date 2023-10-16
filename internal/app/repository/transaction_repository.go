package repository

import (
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"
	"github.com/sakupay-apps/utils/common"
	"github.com/sakupay-apps/utils/constants"

	// "github.com/sakupay-apps/utils/constants"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	BaseRepositoryPaging[model.Transaction]
	Create(payload *model.Transaction) (*model.Transaction, error)
	Get(id string) (*model.Transaction, error)
	ListByUserId(id string) ([]*model.Transaction, error)
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
		if transaction.TransactionType == "deposit" || transaction.TransactionType == "receive" {
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

		if transaction.TransactionType == "send" {
			if err := tx.Create(&transaction).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			if err := tx.Model(&wallet).Where("user_id = ?", transaction.UserID).Select("balance").First(&wallet).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			wallet.Balance -= transaction.Amount

			if err := tx.Model(&wallet).Where("user_id = ?", transaction.UserID).Select("balance").Updates(&wallet).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			return nil
		}

		// card := model.Card{

		// }
		if transaction.TransactionType == "withdraw" {
			if err := tx.Create(&transaction).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			if err := tx.Model(&wallet).Where("user_id = ?", transaction.UserID).Select("balance").First(&wallet).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			wallet.Balance -= transaction.Amount

			if err := tx.Model(&wallet).Where("user_id = ?", transaction.UserID).Select("balance").Updates(&wallet).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			if err := tx.Model(&wallet).Where("user_id = ?", transaction.UserID).Select("balance").Updates(&wallet).Error; err != nil {
				return gorm.ErrInvalidTransaction
			}

			return nil
		}
		return nil
	})
	return &transaction, nil
}

func (r *transactionRepository) ListByUserId(id string) ([]*model.Transaction, error) {
	var transaction []*model.Transaction

	if err := r.db.Where(constants.WHERE_BY_USER_ID, id).Find(&transaction).Error; err != nil {
		return nil, err
	}

	return transaction, nil
}

func (r *transactionRepository) GetTransaction(userID, transactionID string) (*model.Transaction, error) {
	transaction := model.Transaction{}

	if err := r.db.Where("user_id = ? AND id = ?", userID, transactionID).First(&transaction).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return &transaction, nil
}

func (r *transactionRepository) Get(id string) (*model.Transaction, error) {
	transaction := model.Transaction{}

	if err := r.db.Where(constants.WHERE_BY_ID, id).First(&transaction).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return &transaction, nil
}

func (r *transactionRepository) Paging(requestPaging dto.PaginationParam, queries ...string) ([]*model.Transaction, *dto.Paging, error) {
	transactions := []*model.Transaction{}

	paginationQuery := common.GetPaginationParams(requestPaging)

	var totalRows int64

	if err := r.db.Limit(paginationQuery.Take).Offset(paginationQuery.Skip).Find(&transactions).Count(&totalRows).Error; err != nil {
		return nil, nil, err
	}

	var count int = int(totalRows)

	return transactions, common.Paginate(paginationQuery.Take, paginationQuery.Page, count), nil
}
