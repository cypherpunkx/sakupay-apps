package repository

import (
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"

	"github.com/sakupay-apps/utils/common"
	"github.com/sakupay-apps/utils/constants"
	"github.com/sakupay-apps/utils/exception"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ContactRepository interface {
	BaseRepository[model.Contact]
	BaseRepositoryPaging[model.Contact]
}

type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{
		db: db,
	}
}

func (r *contactRepository) Create(payload *model.Contact) (*model.Contact, error) {

	contact := model.Contact{
		ID:           payload.ID,
		UserID:       payload.UserID,
		PhoneNumber:  payload.PhoneNumber,
		Relationship: payload.Relationship,
		IsFavorite:   payload.IsFavorite,
	}

	if err := r.db.Create(&contact).Error; err != nil {
		return nil, exception.ErrFailedCreate
	}

	return &contact, nil
}

func (r *contactRepository) Paging(requestPaging dto.PaginationParam, queries ...string) ([]*model.Contact, *dto.Paging, error) {

	contacts := []*model.Contact{}

	paginationQuery := common.GetPaginationParams(requestPaging)

	var totalRows int64

	if err := r.db.Limit(paginationQuery.Take).Offset(paginationQuery.Skip).Preload("User").Find(&contacts).Count(&totalRows).Error; err != nil {
		return nil, nil, err
	}

	var count int = int(totalRows)

	return contacts, common.Paginate(paginationQuery.Take, paginationQuery.Page, count), nil

}

func (r *contactRepository) List() ([]*model.Contact, error) {

	var contacts []*model.Contact

	if err := r.db.Find(&contacts).Error; err != nil {
		return nil, err
	}

	return contacts, nil
}

func (r *contactRepository) Get(id string) (*model.Contact, error) {
	var contact model.Contact

	if err := r.db.Where(constants.WHERE_BY_ID, id).First(&contact).Error; err != nil {
		return nil, err
	}

	return &contact, nil
}

func (c *contactRepository) Update(id string, payload *model.Contact) (*model.Contact, error) {

	contact := model.Contact{}

	err := c.db.Model(&contact).Where(constants.WHERE_BY_ID, id).Clauses(clause.Returning{}).Updates(model.Contact{
		ID:           payload.ID,
		UserID:       payload.UserID,
		PhoneNumber:  payload.PhoneNumber,
		Relationship: payload.Relationship,
		IsFavorite:   payload.IsFavorite,
	}).Error

	if err != nil {
		return nil, err
	}

	return &contact, nil
}

func (c *contactRepository) Delete(id string) (*model.Contact, error) {
	contact := model.Contact{}

	if err := c.db.Clauses(clause.Returning{}).Where(constants.WHERE_BY_ID, id).Delete(&contact).Error; err != nil {
		return nil, err
	}

	return &contact, nil
}
