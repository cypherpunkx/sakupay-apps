package service

import (
	"testing"

	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockContactRepository struct {
	mock.Mock
}

func (*mockContactRepository) Create(payload *model.Contact) (*model.Contact, error) {
	return &model.Contact{}, nil
}

func (*mockContactRepository) Delete(id string) (*model.Contact, error) {
	return nil, nil
}

func (*mockContactRepository) DeleteContactByID(userID string, contactID string) (*model.Contact, error) {
	return nil, nil
}

func (*mockContactRepository) Get(id string) (*model.Contact, error) {
	return nil, nil
}

func (*mockContactRepository) GetContactByID(userID string, contactID string) (*model.Contact, error) {
	return nil, nil
}

func (*mockContactRepository) ListContacts(id string) ([]*model.Contact, error) {
	return nil, nil
}

func (*mockContactRepository) Paging(requestPaging dto.PaginationParam, query ...string) ([]*model.Contact, *dto.Paging, error) {
	return nil, nil, nil
}

func (*mockContactRepository) Update(id string, payload *model.Contact) (*model.Contact, error) {
	return nil, nil
}

func (r *mockContactRepository) List() ([]*model.Contact, error) {
	return []*model.Contact{}, nil
}

func TestRegisterNewContact(t *testing.T) {
	mockContactRepo := &mockContactRepository{}
	mockUserRepo := &mockUserRepository{}
	contactService := NewContactService(mockContactRepo, mockUserRepo)

	payload := &model.Contact{}
	contactResponse, err := contactService.RegisterNewContact(payload)
	assert.Nil(t, err)
	assert.NotNil(t, contactResponse)
}
