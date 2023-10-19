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
	DepositMoney(payload *model.Transaction, cardID string) (*dto.TransactionResponse, error)
	TransactionHistory(id string) ([]*dto.TransactionResponse, error)
	FindTransactionByUser(userID, transactionID string) (*dto.TransactionResponse, error)
	SendMoney(payload *model.Transaction, friendID string) (*dto.TransactionResponse, error)
	WithdrawMoneyToCard(payload *model.Transaction, cardID string) (*dto.TransactionResponse, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
	userRepo        repository.UserRepository
	cardRepo        repository.CardRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository, userRepo repository.UserRepository, cardRepo repository.CardRepository) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
		userRepo:        userRepo,
		cardRepo:        cardRepo,
	}
}

func (s *transactionService) DepositMoney(payload *model.Transaction, cardID string) (*dto.TransactionResponse, error) {
	if payload.Amount <= 10000 {
		return nil, exception.ErrMinimalTransaction
	}

	user, err := s.userRepo.Get(payload.UserID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	card, err := s.cardRepo.Get(cardID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	userCard, err := s.cardRepo.GetCardUserID(user.ID, card.ID)

	if payload.Amount > userCard.Balance {
		return nil, exception.ErrNotEnoughBalance
	}

	transaction, err := s.transactionRepo.CreateDepositTransaction(payload, card.ID)

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	transactionResponse := dto.TransactionResponse{
		ID: transaction.ID,
		User: model.User{
			ID:       user.ID,
			Username: user.Username,
			Wallet: model.Wallet{
				Name:    user.Wallet.Name,
				Balance: user.Wallet.Balance,
			},
		},
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
				ID: transaction.ID,
				User: model.User{
					Username: user.Username,
					Wallet: model.Wallet{
						Name:    user.Wallet.Name,
						Balance: user.Wallet.Balance,
					},
				},
				TransactionType: transaction.TransactionType,
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
		ID: userTransaction.ID,
		User: model.User{
			Username: user.Username,
			Wallet: model.Wallet{
				Name: user.Wallet.Name,
			},
		},
		TransactionType: userTransaction.TransactionType,
		Amount:          userTransaction.Amount,
		Description:     userTransaction.Description,
		Timestamp:       userTransaction.Timestamp,
	}

	return &transactionResponse, err
}

func (s *transactionService) SendMoney(payload *model.Transaction, friendID string) (*dto.TransactionResponse, error) {
	if payload.Amount <= 10000 {
		return nil, exception.ErrMinimalTransaction
	}

	user, err := s.userRepo.Get(payload.UserID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	if payload.Amount > user.Wallet.Balance {
		return nil, exception.ErrNotEnoughBalance
	}

	friend, err := s.userRepo.Get(friendID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	transaction, err := s.transactionRepo.CreateSendTransaction(payload, friend.ID)

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	transactionResponse := dto.TransactionResponse{
		ID: transaction.ID,
		User: model.User{
			Username: user.Username,
			Wallet: model.Wallet{
				Name:    user.Wallet.Name,
				Balance: user.Wallet.Balance,
			},
		},
		TransactionType: transaction.TransactionType,
		Amount:          transaction.Amount,
		Description:     transaction.Description,
		Timestamp:       time.Now(),
	}

	return &transactionResponse, err
}

func (s *transactionService) WithdrawMoneyToCard(payload *model.Transaction, cardID string) (*dto.TransactionResponse, error) {
	if payload.Amount <= 10000 {
		return nil, exception.ErrMinimalTransaction
	}

	user, err := s.userRepo.Get(payload.UserID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	if payload.Amount > user.Wallet.Balance {
		return nil, exception.ErrNotEnoughBalance
	}

	card, err := s.cardRepo.Get(cardID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	transaction, err := s.transactionRepo.CreateWitdrawTransaction(payload, card.ID)

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	transactionResponse := dto.TransactionResponse{
		ID: transaction.ID,
		User: model.User{
			Username: user.Username,
			Wallet: model.Wallet{
				Name:    user.Wallet.Name,
				Balance: user.Wallet.Balance,
			},
		},
		TransactionType: transaction.TransactionType,
		Amount:          transaction.Amount,
		Description:     transaction.Description,
		Timestamp:       time.Now(),
	}

	return &transactionResponse, err
}
