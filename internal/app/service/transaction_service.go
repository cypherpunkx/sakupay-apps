package service

import (
	"errors"
	"time"

	"github.com/sakupay-apps/internal/app/repository"
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"
	"github.com/sakupay-apps/utils/exception"
	"gorm.io/gorm"
)

type TransactionService interface {
	Deposit(payload *model.Transaction) (*dto.TransactionResponse, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
	userRepo        repository.UserRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository, userRepo repository.UserRepository) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
		userRepo:        userRepo,
	}
}

func (s *transactionService) Deposit(payload *model.Transaction) (*dto.TransactionResponse, error) {
	user, err := s.userRepo.Get(payload.UserID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	if user.Wallet.Balance < 0 {
		return nil, errors.New("balance kurang")
	}

	transaction, err := s.transactionRepo.Create(payload)

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	transactionResponse := dto.TransactionResponse{
		ID:              transaction.ID,
		User:            *user,
		TransactionType: transaction.TransactionType,
		Amount:          transaction.Amount,
		Description:     transaction.Description,
		Timestamp:       time.Now(),
	}

	return &transactionResponse, err
}
