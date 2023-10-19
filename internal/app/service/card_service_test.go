package service

import (
	"testing"

	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCardRepository struct {
	mock.Mock
}

// GetCardUserID implements repository.CardRepository.
func (*mockCardRepository) GetCardUserID(userID string, cardID string) (*model.Card, error) {
	panic("unimplemented")
}

func (*mockCardRepository) Delete(id string) (*model.Card, error) {
	return nil, nil
}

func (*mockCardRepository) DeleteCardID(userID string, cardID string) (*model.Card, error) {
	return nil, nil
}

func (*mockCardRepository) Get(id string) (*model.Card, error) {
	return nil, nil
}

func (*mockCardRepository) ListCards(id string) ([]*model.Card, error) {
	return nil, nil
}

func (*mockCardRepository) Paging(requestPaging dto.PaginationParam, queries ...string) ([]*model.Card, *dto.Paging, error) {
	return nil, nil, nil
}

func (*mockCardRepository) Create(card *model.Card) (*model.Card, error) {
	return &model.Card{}, nil
}

type mockUserRepository struct {
	mock.Mock
}

func (*mockUserRepository) Create(payload *model.User) (*model.User, error) {
	return nil, nil
}

func (*mockUserRepository) Delete(id string) (*model.User, error) {
	return nil, nil
}

func (*mockUserRepository) GetUsername(username string) (*model.User, error) {
	return nil, nil
}

func (*mockUserRepository) GetUsernamePassword(username string, password string) (*model.User, error) {
	return nil, nil
}

func (*mockUserRepository) List() ([]*model.User, error) {
	return nil, nil
}

func (*mockUserRepository) Paging(requestPaging dto.PaginationParam, query ...string) ([]*model.User, *dto.Paging, error) {
	return nil, nil, nil
}

func (*mockUserRepository) Update(id string, payload *model.User) (*model.User, error) {
	return nil, nil
}

func (*mockUserRepository) Get(id string) (*model.User, error) {
	return &model.User{}, nil
}

func TestRegisterNewCard(t *testing.T) {
	mockCardRepo := &mockCardRepository{}
	mockUserRepo := &mockUserRepository{}
	cardService := NewCardService(mockCardRepo, mockUserRepo)

	// Test case 1: Positive case
	payload := &model.Card{}
	cardResponse, err := cardService.RegisterNewCard(payload)
	assert.Nil(t, err)
	assert.NotNil(t, cardResponse)
}
