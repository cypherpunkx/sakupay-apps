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

type CardRepository interface {
	Create(payload *model.Card) (*model.Card, error)
	Paging(requestPaging dto.PaginationParam, queries ...string) ([]*model.Card, *dto.Paging, error)
	ListCards(id string) ([]*model.Card, error)
	Get(id string) (*model.Card, error)
	Delete(id string) (*model.Card, error)
	DeleteCardID(userID, cardID string) (*model.Card, error)
}

type cardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) CardRepository {
	return &cardRepository{
		db: db,
	}
}

func (cr *cardRepository) Create(payload *model.Card) (*model.Card, error) {

	card := model.Card{
		ID:             payload.ID,
		UserID:         payload.UserID,
		CardNumber:     payload.CardNumber,
		CardholderName: payload.CardholderName,
		ExpirationDate: payload.ExpirationDate,
		Balance:        payload.Balance,
		CVV:            payload.CVV,
	}

	if err := cr.db.Create(&card).Error; err != nil {
		return nil, exception.ErrFailedCreate
	}

	return &card, nil
}

func (cr *cardRepository) ListCards(id string) ([]*model.Card, error) {
	cards := []*model.Card{}

	if err := cr.db.Model(&cards).Where(constants.WHERE_BY_USER_ID, id).Find(&cards).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return cards, nil
}

func (cr *cardRepository) Paging(requestPaging dto.PaginationParam, queries ...string) ([]*model.Card, *dto.Paging, error) {

	cards := []*model.Card{}

	paginationQuery := common.GetPaginationParams(requestPaging)

	var totalRows int64

	if err := cr.db.Limit(paginationQuery.Take).Offset(paginationQuery.Skip).Preload("Card").Find(&cards).Count(&totalRows).Error; err != nil {
		return nil, nil, err
	}

	var count int = int(totalRows)

	return cards, common.Paginate(paginationQuery.Take, paginationQuery.Page, count), nil

}

func (cr *cardRepository) Get(id string) (*model.Card, error) {
	var card model.Card

	if err := cr.db.Where(constants.WHERE_BY_ID, id).First(&card).Error; err != nil {
		return nil, err
	}

	return &card, nil
}

func (c *cardRepository) Delete(id string) (*model.Card, error) {
	card := model.Card{}

	if err := c.db.Clauses(clause.Returning{}).Where(constants.WHERE_BY_ID, id).Delete(&card).Error; err != nil {
		return nil, err
	}

	return &card, nil
}

func (c *cardRepository) DeleteCardID(userID, cardID string) (*model.Card, error) {
	card := model.Card{}

	if err := c.db.Clauses(clause.Returning{}).Where(constants.WHERE_BY_USER_ID_AND_CARD_ID, userID, cardID).Delete(&card).Error; err != nil {
		return nil, err
	}

	return &card, nil
}
