package repository

import (
	//"fmt"

	// "fmt"

	"fmt"

	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"

	"github.com/sakupay-apps/utils/common"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ContactRepository interface {
	Create(payload model.Contact) error
	List() ([]model.Contact, error)
	Get(id string) (*model.Contact, error)
	Update(id string, payload *model.Contact) (*model.Contact, error)
	Delete(id string) (*model.Contact, error)
	Paging(requestPaging dto.PaginationParam, query ...string) ([]model.Contact, dto.Paging, error)
}

type contactRepository struct {
	db *gorm.DB
}

func (c *contactRepository) Create(payload model.Contact) error {

	contactCreate := &model.Contact{
		ID:           payload.ID,
		UserID:       payload.UserID,
		PhoneNumber:  payload.PhoneNumber,
		Relationship: payload.Relationship,
		IsFavorite:   payload.IsFavorite,
	}

	err := c.db.Create(&contactCreate).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *contactRepository) Paging(requestPaging dto.PaginationParam, query ...string) ([]model.Contact, dto.Paging, error) {

	var paginationQuery dto.PaginationQuery
	var contact []model.Contact
	paginationQuery = common.GetPaginationParams(requestPaging)
	//querySelect := c.db.Limit(1).Offset(2).Find(&contact)

	querySelect:= c.db.Limit(paginationQuery.Take).Offset(paginationQuery.Skip).Find(&contact)
	if querySelect.Error != nil {
		return nil, dto.Paging{}, querySelect.Error
	}

	// Menghitung total jumlah data tanpa menggunakan LIMIT dan OFFSET
	var totalRows int64
	Count := c.db.Model(&model.Contact{}).Count(&totalRows)
	if Count.Error != nil {
		return nil, dto.Paging{}, Count.Error
	}

	totalPage := common.CountTotalPage(int(totalRows), paginationQuery.Take)

	page := dto.Paging{
		Page:        paginationQuery.Page,
		RowsPerPage: paginationQuery.Take,
		TotalPages:  totalPage,
		TotalRows: int(totalRows),
	}

	return contact, page, nil

}

func (c *contactRepository) List() ([]model.Contact, error) {

	var contactList []model.Contact

	if err := c.db.Find(&contactList).Error; err != nil {
		return nil, err
	}

	return contactList, nil
}

func (c *contactRepository) Get(id string) (*model.Contact, error) {
	var contact model.Contact
	err := c.db.Where("id = ?", id).First(&contact).Error 

	if err != nil {
		return nil, err
	}
	fmt.Println(contact)
	return &contact, nil
}

func (c *contactRepository) Update(id string, payload *model.Contact) (*model.Contact, error) {
     
	contact := model.Contact{}

	contacts := c.db.Model(&contact).Where("id = ?", id).Clauses(clause.Returning{}).Updates(model.Contact{

		ID: payload.ID,
		UserID:  payload.UserID,
		PhoneNumber: payload.PhoneNumber,
		Relationship:    payload.Relationship,
		IsFavorite:   payload.IsFavorite,
		

	})
	if contacts.Error != nil {
		return nil, contacts.Error
	}
	return &contact, nil
}

func (c *contactRepository) Delete(id string) (*model.Contact, error) {
	contact := model.Contact{}

	if err := c.db.Clauses(clause.Returning{}).Where("id = ?", id).Delete(&contact).Error; err != nil {
		return nil, err
	}

	return &contact, nil
}


func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{
		db: db,
	}
}
