package service

import (
	"github.com/sakupay-apps/internal/app/repository"
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"
	"github.com/sakupay-apps/utils/exception"
	"gorm.io/gorm"
)

type CardService interface {
	RegisterNewCard(payload *model.Card) (*dto.CardResponse, error)
	FindAllCardList(id string) ([]*dto.CardResponse, error)
	FindCardByID(id string) (*dto.CardResponse, error)
	DeleteCardByID(userID, cardID string) (*dto.CardResponse, error)
	FindAllCard(requesPaging dto.PaginationParam, byNameEmpl string) ([]*model.Card, *dto.Paging, error)
}

type cardService struct {
	cardRepo repository.CardRepository
	userRepo repository.UserRepository
}

func NewCardService(cardRepo repository.CardRepository, userRepo repository.UserRepository) CardService {
	return &cardService{
		cardRepo: cardRepo,
		userRepo: userRepo,
	}
}

func (cs *cardService) RegisterNewCard(payload *model.Card) (*dto.CardResponse, error) {
	user, err := cs.userRepo.Get(payload.UserID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	cards, err := cs.cardRepo.ListCards(user.ID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	for _, card := range cards {
		if card.CardNumber == payload.CardNumber {
			return nil, exception.ErrCardNumberAlreadyExist
		}
	}

	card, err := cs.cardRepo.Create(payload)

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	cardResponses := dto.CardResponse{
		ID:             card.ID,
		User:           *user,
		CardNumber:     card.CardNumber,
		CardholderName: card.CardholderName,
		Balance:        card.Balance,
		ExpirationDate: card.ExpirationDate,
		CVV:            card.CVV,
	}

	return &cardResponses, err

}

func (cs *cardService) FindAllCardList(id string) ([]*dto.CardResponse, error) {
	user, err := cs.userRepo.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	cards, err := cs.cardRepo.ListCards(user.ID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	cardResponses := []*dto.CardResponse{}

	for _, card := range cards {
		if card.UserID == user.ID {
			cardResponses = append(cardResponses, &dto.CardResponse{
				ID:             card.ID,
				User:           *user,
				CardNumber:     card.CardNumber,
				Balance:        card.Balance,
				CardholderName: card.CardholderName,
				ExpirationDate: card.ExpirationDate,
				CVV:            card.CVV,
			})
		}
	}

	return cardResponses, err
}

func (cs *cardService) FindCardByID(id string) (*dto.CardResponse, error) {

	user, err := cs.userRepo.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	card, err := cs.cardRepo.Get(user.ID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	cardResponse := dto.CardResponse{
		ID:             card.ID,
		User:           *user,
		CardNumber:     card.CardNumber,
		CardholderName: card.CardholderName,
		Balance:        card.Balance,
		ExpirationDate: card.ExpirationDate,
		CVV:            card.CVV,
	}

	return &cardResponse, err
}

func (cs *cardService) DeleteCardByID(userID, cardID string) (*dto.CardResponse, error) {
	user, err := cs.userRepo.Get(userID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	card, err := cs.cardRepo.Get(cardID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	userCard, err := cs.cardRepo.DeleteCardID(user.ID, card.ID)

	cardResponse := dto.CardResponse{
		ID:             userCard.ID,
		User:           *user,
		CardNumber:     userCard.CardNumber,
		CardholderName: userCard.CardholderName,
		ExpirationDate: userCard.ExpirationDate,
		Balance:        userCard.Balance,
		CVV:            userCard.CVV,
	}

	return &cardResponse, err
}

func (cs *cardService) FindAllCard(requesPaging dto.PaginationParam, byNameEmpl string) ([]*model.Card, *dto.Paging, error) {
	return cs.cardRepo.Paging(requesPaging, byNameEmpl)
}
