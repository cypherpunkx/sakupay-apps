package service

import (
	"github.com/sakupay-apps/internal/app/repository"
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"
	"github.com/sakupay-apps/utils/exception"
	"gorm.io/gorm"
)

type ContactService interface {
	RegisterNewContact(payload *model.Contact) (*dto.ContactResponse, error)
	FindAllContactList(id string) ([]*dto.ContactResponse, error)
	FindContactById(id string) (*dto.ContactResponse, error)
	// UpdateContact(id string, payload *model.Contact) (*model.Contact, error)
	// DeleteContact(id string) (*model.Contact, error)
	// FindAllContact(requesPaging dto.PaginationParam, byNameEmpl string) ([]*model.Contact, *dto.Paging, error)
}

type contactService struct {
	contactRepo repository.ContactRepository
	userRepo    repository.UserRepository
}

func NewContactService(contactRepo repository.ContactRepository, userRepo repository.UserRepository) ContactService {
	return &contactService{
		contactRepo: contactRepo,
		userRepo:    userRepo,
	}
}

func (s *contactService) RegisterNewContact(payload *model.Contact) (*dto.ContactResponse, error) {

	contacts, err := s.contactRepo.List()

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	for _, contact := range contacts {
		if contact.PhoneNumber == payload.PhoneNumber {
			return nil, exception.ErrPhoneNumberAlreadyExist
		}
	}

	user, err := s.userRepo.Get(payload.UserID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	contact, err := s.contactRepo.Create(payload)

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	contactResponses := dto.ContactResponse{
		ID:           contact.ID,
		User:         *user,
		PhoneNumber:  contact.PhoneNumber,
		Relationship: contact.Relationship,
		IsFavorite:   contact.IsFavorite,
	}

	return &contactResponses, err

}

func (s *contactService) FindAllContactList(id string) ([]*dto.ContactResponse, error) {
	user, err := s.userRepo.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	contacts, err := s.contactRepo.List()

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	contactResponses := []*dto.ContactResponse{}

	for _, contact := range contacts {
		if contact.UserID == user.ID {
			contactResponses = append(contactResponses, &dto.ContactResponse{
				ID:           contact.ID,
				User:         *user,
				PhoneNumber:  contact.PhoneNumber,
				Relationship: contact.Relationship,
				IsFavorite:   contact.IsFavorite,
			})
		}
	}

	return contactResponses, err
}

func (s *contactService) FindContactById(id string) (*dto.ContactResponse, error) {

	user, err := s.userRepo.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	contact, err := s.contactRepo.Get(user.ID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	contactResponse := dto.ContactResponse{
		ID:           contact.ID,
		User:         *user,
		PhoneNumber:  contact.PhoneNumber,
		Relationship: contact.PhoneNumber,
		IsFavorite:   contact.IsFavorite,
	}

	return &contactResponse, err
}

// func (s *contactService) DeleteContact(id string) (*dto.ContactResponse, error) {
// 	contact, err := s.contactRepo.Get(id)

// 	if err != nil {
// 		return nil,gorm.ErrRecordNotFound
// 	}

// 	user

// 	contactResponse := dto.ContactResponse{
// 		ID: contact.ID,
// 		User: ,
// 	}

// 	return
// }

// func (c *contactService) FindAllContact(requesPaging dto.PaginationParam, byNameEmpl string) ([]*model.Contact, *dto.Paging, error) {
// 	return c.repo.Paging(requesPaging, byNameEmpl)
// }
