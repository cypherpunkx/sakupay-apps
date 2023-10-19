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

func (r *MockTransactionRepository) Get(id string) (*model.Transaction, error) {
	args := r.Called(id)
	return args.Get(0).(*model.Transaction), args.Error(1)
}

func (r *MockTransactionRepository) GetTransaction(userID string, transactionID string) (*model.Transaction, error) {
	args := r.Called(userID, transactionID)
	return args.Get(0).(*model.Transaction), args.Error(1)
}

func (*MockTransactionRepository) List() ([]*model.Transaction, error) {
	return nil, nil
}

func (r *MockTransactionRepository) ListTransactions(id string) ([]*model.Transaction, error) {
	args := r.Called(id)
	return args.Get(0).([]*model.Transaction), args.Error(1)
}

func (r *MockTransactionRepository) Create(payload *model.Transaction) (*model.Transaction, error) {
	args := r.Called(payload)
	return args.Get(0).(*model.Transaction), args.Error(1)
}

func TestCreateNewTransaction(t *testing.T) {
	mockUserRepo := &MockUserRepository{}
	mockTransactionRepo := &MockTransactionRepository{}
	service := NewTransactionService(mockTransactionRepo, mockUserRepo)

	mockUser := &model.User{}
	mockUserRepo.On("Get", mock.Anything).Return(mockUser, nil)

	mockTransaction := &model.Transaction{}
	mockTransactionRepo.On("Create", mock.Anything).Return(mockTransaction, nil)

	payload := &model.Transaction{}
	response, err := service.CreateNewTransaction(payload)

	assert.NotNil(t, response)
	assert.Nil(t, err)

	mockUserRepo.AssertExpectations(t)
	mockTransactionRepo.AssertExpectations(t)
}
