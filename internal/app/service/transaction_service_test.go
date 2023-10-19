package service

import (
	"testing"

	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (r *MockUserRepository) Create(payload *model.User) (*model.User, error) {
	return nil, nil
}

func (r *MockUserRepository) Delete(id string) (*model.User, error) {
	return nil, nil
}

func (r *MockUserRepository) GetUsername(username string) (*model.User, error) {
	return nil, nil
}

func (r *MockUserRepository) GetUsernamePassword(username string, password string) (*model.User, error) {
	return nil, nil
}

func (r *MockUserRepository) List() ([]*model.User, error) {
	return []*model.User{}, nil
}

func (r *MockUserRepository) Paging(requestPaging dto.PaginationParam, query ...string) ([]*model.User, *dto.Paging, error) {
	return nil, nil, nil
}

func (r *MockUserRepository) Update(id string, payload *model.User) (*model.User, error) {
	return nil, nil
}

func (r *MockUserRepository) Get(id string) (*model.User, error) {
	args := r.Called(id)
	return args.Get(0).(*model.User), args.Error(1)
}

type MockTransactionRepository struct {
	mock.Mock
}

// Get implements repository.TransactionRepository.
func (*MockTransactionRepository) Get(id string) (*model.Transaction, error) {
	return nil, nil
}

// GetTransaction implements repository.TransactionRepository.
func (*MockTransactionRepository) GetTransaction(userID string, transactionID string) (*model.Transaction, error) {
	return nil, nil
}

// List implements repository.TransactionRepository.
func (*MockTransactionRepository) List() ([]*model.Transaction, error) {
	return nil, nil
}

// ListTransactions implements repository.TransactionRepository.
func (*MockTransactionRepository) ListTransactions(id string) ([]*model.Transaction, error) {
	return nil, nil
}

// CreateDepositTransaction implements repository.TransactionRepository.
func (*MockTransactionRepository) CreateDepositTransaction(payload *model.Transaction, cardID string) (*model.Transaction, error) {
	return nil, nil
}

// CreateSendTransaction implements repository.TransactionRepository.
func (*MockTransactionRepository) CreateSendTransaction(payload *model.Transaction, friendID string) (*model.Transaction, error) {
	return nil, nil
}

// CreateWitdrawTransaction implements repository.TransactionRepository.
func (*MockTransactionRepository) CreateWitdrawTransaction(payload *model.Transaction, cardID string) (*model.Transaction, error) {
	return nil, nil
}

type MockCardRepository struct {
	mock.Mock
}

func (*MockCardRepository) Create(payload *model.Card) (*model.Card, error) {
	return nil, nil
}

func (*MockCardRepository) ListCards(id string) ([]*model.Card, error) {
	return nil, nil
}

func (*MockCardRepository) Paging(requestPaging dto.PaginationParam, queries ...string) ([]*model.Card, *dto.Paging, error) {
	return nil, nil, nil
}

func (*MockCardRepository) Get(id string) (*model.Card, error) {
	return nil, nil
}

func (*MockCardRepository) GetCardUserID(userID, cardID string) (*model.Card, error) {
	return nil, nil
}

func (*MockCardRepository) Delete(id string) (*model.Card, error) {
	return nil, nil
}

func (*MockCardRepository) DeleteCardID(userID, cardID string) (*model.Card, error) {
	return nil, nil
}

func TestCreateNewTransaction(t *testing.T) {
	mockUserRepo := &MockUserRepository{}
	mockTransactionRepo := &MockTransactionRepository{}
	mockCardRepo := &MockCardRepository{}

	service := NewTransactionService(mockTransactionRepo, mockUserRepo, mockCardRepo)

	mockUser := &model.User{}
	mockUserRepo.On("Get", mock.Anything).Return(mockUser, nil)

	mockTransaction := &model.Transaction{}
	mockTransactionRepo.On("Create", mock.Anything).Return(mockTransaction, nil)
	mockCard := &model.Card{}
	mockTransactionRepo.On("Get", mock.Anything).Return(mockTransaction, nil)

	payload := &model.Transaction{}
	response, err := service.DepositMoney(payload, mockCard.ID)

	assert.NotNil(t, response)
	assert.Nil(t, err)

	mockUserRepo.AssertExpectations(t)
	mockTransactionRepo.AssertExpectations(t)
	mockCardRepo.AssertExpectations(t)
}
