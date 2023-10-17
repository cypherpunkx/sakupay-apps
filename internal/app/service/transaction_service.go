package service

import (
	"time"

	"github.com/sakupay-apps/internal/app/repository"
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"
	"github.com/sakupay-apps/utils/exception"
	"gorm.io/gorm"
)

type TransactionService interface {
	CreateNewTransaction(payload *model.Transaction) (*dto.TransactionResponse, error)
	TransactionHistory(id string) ([]*dto.TransactionResponse, error)
	FindTransactionByUser(userID, transactionID string) (*dto.TransactionResponse, error)
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

func (s *transactionService) CreateNewTransaction(payload *model.Transaction) (*dto.TransactionResponse, error) {
	user, err := s.userRepo.Get(payload.UserID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	if payload.TransactionType == "send" && payload.Amount > user.Wallet.Balance {
		return nil, exception.ErrNotEnoughBalance
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

func (s *transactionService) TransactionHistory(id string) ([]*dto.TransactionResponse, error) {

	user, err := s.userRepo.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	transactions, err := s.transactionRepo.ListTransactions(user.ID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	transactionResponses := []*dto.TransactionResponse{}

	for _, transaction := range transactions {
		if transaction.UserID == user.ID {
			transactionResponses = append(transactionResponses, &dto.TransactionResponse{
				ID:              transaction.ID,
				TransactionType: transaction.TransactionType,
				User:            *user,
				Amount:          transaction.Amount,
				Description:     transaction.Description,
				Timestamp:       transaction.Timestamp,
			})
		}
	}

	return transactionResponses, err
}

func (s *transactionService) FindTransactionByUser(userID, transactionID string) (*dto.TransactionResponse, error) {

	user, err := s.userRepo.Get(userID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	transaction, err := s.transactionRepo.Get(transactionID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	userTransaction, err := s.transactionRepo.GetTransaction(user.ID, transaction.ID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	transactionResponse := dto.TransactionResponse{
		ID:              userTransaction.ID,
		User:            *user,
		TransactionType: userTransaction.TransactionType,
		Amount:          userTransaction.Amount,
		Description:     userTransaction.Description,
		Timestamp:       userTransaction.Timestamp,
	}

	return &transactionResponse, err
}
