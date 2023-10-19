package service

import (
	"testing"

	"github.com/sakupay-apps/internal/model"
	"github.com/stretchr/testify/assert"
)

type mockBillRepository struct{}

func (*mockBillRepository) Get(id string) (*model.Bill, error) {
	return nil, nil
}

type mockBillDetailsRepository struct{}

func (*mockBillDetailsRepository) Get(id string) (*model.BillDetails, error) {
	return nil, nil
}

func (m *mockBillRepository) Create(payload *model.Bill) (*model.Bill, error) {
	return &model.Bill{}, nil
}

func (m *mockBillRepository) List(userID string) ([]*model.Bill, error) {
	return []*model.Bill{}, nil
}

func TestCreateNewBill(t *testing.T) {
	mockBillRepo := &mockBillRepository{}
	mockUserRepo := &mockUserRepository{}
	mockBillDetailsRepo := &mockBillDetailsRepository{}
	billService := NewBillService(mockBillRepo, mockUserRepo, mockBillDetailsRepo)

	payload := &model.Bill{}

	billResponse, err := billService.CreateNewBill(payload)
	assert.Nil(t, err)
	assert.NotNil(t, billResponse)
}
