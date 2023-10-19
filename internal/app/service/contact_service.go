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
	FindContactByUser(userID, contactID string) (*dto.ContactResponse, error)
	DeleteContactByUser(userID, contactID string) (*dto.ContactResponse, error)
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
		return nil, gorm.ErrRecordNotFound
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
		ID: contact.ID,
		User: model.User{
			ID:       user.ID,
			Username: user.Username,
			Wallet: model.Wallet{
				Name:    user.Wallet.Name,
				Balance: user.Wallet.Balance,
			},
		},
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

	contacts, err := s.contactRepo.ListContacts(user.ID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	contactResponses := []*dto.ContactResponse{}

	for _, contact := range contacts {
		if contact.UserID == user.ID {
			contactResponses = append(contactResponses, &dto.ContactResponse{
				ID: contact.ID,
				User: model.User{
					ID:       user.ID,
					Username: user.Username,
					Wallet: model.Wallet{
						Name:    user.Wallet.Name,
						Balance: user.Wallet.Balance,
					},
				},
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
		ID: contact.ID,
		User: model.User{
			ID:       user.ID,
			Username: user.Username,
			Wallet: model.Wallet{
				Name:    user.Wallet.Name,
				Balance: user.Wallet.Balance,
			},
		},
		PhoneNumber:  contact.PhoneNumber,
		Relationship: contact.PhoneNumber,
		IsFavorite:   contact.IsFavorite,
	}

	return &contactResponse, err
}

func (s *contactService) FindContactByUser(userID, contactID string) (*dto.ContactResponse, error) {

	user, err := s.userRepo.Get(userID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	contact, err := s.contactRepo.Get(contactID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	userContact, err := s.contactRepo.GetContactByID(user.ID, contact.ID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	contactResponse := dto.ContactResponse{
		ID: userContact.ID,
		User: model.User{
			ID:       user.ID,
			Username: user.Username,
			Wallet: model.Wallet{
				Name:    user.Wallet.Name,
				Balance: user.Wallet.Balance,
			},
		},
		PhoneNumber:  userContact.PhoneNumber,
		Relationship: userContact.PhoneNumber,
		IsFavorite:   userContact.IsFavorite,
	}

	return &contactResponse, err
}

func (s *contactService) DeleteContactByUser(userID, contactID string) (*dto.ContactResponse, error) {

	user, err := s.userRepo.Get(userID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	contact, err := s.contactRepo.Get(contactID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	userContact, err := s.contactRepo.DeleteContactByID(user.ID, contact.ID)

	if err != nil {
		return nil, exception.ErrFailedDelete
	}

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	contactResponse := dto.ContactResponse{
		ID: userContact.ID,
		User: model.User{
			ID:       user.ID,
			Username: user.Username,
			Wallet: model.Wallet{
				Name:    user.Wallet.Name,
				Balance: user.Wallet.Balance,
			},
		},
		PhoneNumber:  userContact.PhoneNumber,
		Relationship: userContact.PhoneNumber,
		IsFavorite:   userContact.IsFavorite,
	}

	return &contactResponse, err
}
