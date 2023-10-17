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
	FindCardById(id string) (*dto.CardResponse, error)
	// UpdateCard(id string, payload *model.Card) (*model.Card, error)
	DeleteCard(id string) (*dto.CardResponse, error)
	FindAllCard(requesPaging dto.PaginationParam, byNameEmpl string) ([]*model.Card, *dto.Paging, error)
}

type cardService struct {
	cardRepo repository.CardRepository
	userRepo    repository.UserRepository
}

func NewCardService(cardRepo repository.CardRepository, userRepo repository.UserRepository) CardService {
	return &cardService{
		cardRepo: cardRepo,
		userRepo:    userRepo,
		
	}
}

func (cs *cardService) RegisterNewCard(payload *model.Card) (*dto.CardResponse, error) {

	cards, err := cs.cardRepo.List()

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	for _, card := range cards {
		if card.CardNumber == payload.CardNumber {
			return nil, exception.ErrPhoneNumberAlreadyExist
		}
	}

	user, err := cs.userRepo.Get(payload.UserID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	card, err := cs.cardRepo.Create(payload)

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	cardResponses := dto.CardResponse{
		ID:           card.ID,
		User:         *user,
		CardNumber:  card.CardNumber,
		CardholderName: card.CardholderName,
		ExpirationDate:   card.ExpirationDate,
		CVV: card.CVV,
	}

	return &cardResponses, err

}

func (cs *cardService) FindAllCardList(id string) ([]*dto.CardResponse, error) {
	user, err := cs.userRepo.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	cards, err := cs.cardRepo.List()

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	cardResponses := []*dto.CardResponse{}

	for _, card := range cards {
		if card.UserID == user.ID {
			cardResponses = append(cardResponses, &dto.CardResponse{
				ID:           card.ID,
				User:         *user,
				CardNumber: card.CardNumber,
				CardholderName: card.CardholderName,
				ExpirationDate: card.ExpirationDate,
				CVV: card.CVV,

			})
		}
	}

	return cardResponses, err
}

func (cs *cardService) FindCardById(id string) (*dto.CardResponse, error) {

	user, err := cs.userRepo.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	card, err := cs.cardRepo.Get(user.ID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	cardResponse := dto.CardResponse{
		ID:           card.ID,
		User:         *user,
		CardNumber: card.CardNumber,
		CardholderName: card.CardholderName,
		ExpirationDate: card.ExpirationDate,
		CVV: card.CVV,
	}

	return &cardResponse, err
}


func (cs *cardService) DeleteCard(id string) (*dto.CardResponse, error) {
		card, err := cs.cardRepo.Get(id)
	
		if err != nil {
			return nil,gorm.ErrRecordNotFound
		}
	
		user, err := cs.userRepo.Get(id)

		if err != nil {
			return nil, gorm.ErrRecordNotFound
		}

			cardResponse := dto.CardResponse{
				ID:           card.ID,
				User:         *user,
				CardNumber: card.CardNumber,
				CardholderName: card.CardholderName,
				ExpirationDate: card.ExpirationDate,
				CVV: card.CVV,
			}
	
		return &cardResponse, err
	}
	
	func (cs *cardService) FindAllCard(requesPaging dto.PaginationParam, byNameEmpl string) ([]*model.Card, *dto.Paging, error) {
		return cs.cardRepo.Paging(requesPaging, byNameEmpl)
	}
	