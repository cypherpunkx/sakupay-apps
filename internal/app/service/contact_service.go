package service

import (
	"github.com/sakupay-apps/internal/app/repository"
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"
)

type ContactService interface {
	RegisterNewContact(payload model.Contact) error
	FindAllContactList() ([]model.Contact, error)
	FindContactById(id string) (*model.Contact,error)
	UpdateContact(id string, payload *model.Contact) (*model.Contact, error)
	DeleteContact(id string) (*model.Contact, error)
        FindAllContact(requesPaging dto.PaginationParam, byNameEmpl string) ([]model.Contact, dto.Paging, error)

}

type contactService struct {
	repo repository.ContactRepository

}

func (c *contactService) RegisterNewContact(payload model.Contact) error {

   return c.repo.Create(payload)
	
}

func  (c *contactService) FindAllContactList() ([]model.Contact, error) {
	return c.repo.List() 
}

func (c *contactService) FindContactById(id string) (*model.Contact,error) {
	return c.repo.Get(id)
}

func (c *contactService) UpdateContact(id string, payload *model.Contact) (*model.Contact, error) {	
	_, err := c.FindContactById(payload.ID)
	
	if err != nil {
		return nil,err
	}
	return c.repo.Update(id, payload)
}

func (c *contactService) DeleteContact(id string) (*model.Contact, error) {
	_, err := c.FindContactById(id)
	if err != nil {
		return nil,err
	}
	return c.repo.Delete(id)
}

func (c *contactService) FindAllContact(requesPaging dto.PaginationParam, byNameEmpl string) ([]model.Contact, dto.Paging, error){
	return c.repo.Paging(requesPaging,byNameEmpl)
}


func NewContactService(contRepo repository.ContactRepository) ContactService {
	return &contactService{
		repo: contRepo,
	}
}

