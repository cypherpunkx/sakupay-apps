package repository

import (
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"
	"github.com/sakupay-apps/utils/common"
	"github.com/sakupay-apps/utils/constants"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	BaseRepository[model.User]
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(payload *model.User) (*model.User, error) {
	user := model.User{
		ID:          payload.ID,
		Username:    payload.Username,
		Email:       payload.Email,
		Password:    payload.Password,
		FirstName:   payload.FirstName,
		LastName:    payload.LastName,
		PhoneNumber: payload.PhoneNumber,
	}

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) List(requestPaging dto.PaginationParam, queries ...string) ([]*model.User, *dto.Paging, error) {
	users := []*model.User{}

	paginationQuery := common.GetPaginationParams(requestPaging)

	var totalRows int64

	if err := r.db.Limit(paginationQuery.Take).Offset(paginationQuery.Skip).Find(&users).Count(&totalRows).Error; err != nil {
		return nil, nil, err
	}

	var count int = int(totalRows)

	return users, common.Paginate(paginationQuery.Take, paginationQuery.Page, count), nil
}

func (r *userRepository) Get(id string) (*model.User, error) {
	user := model.User{}

	if err := r.db.Where(constants.WHERE_BY_ID, id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Update(id string, payload *model.User) (*model.User, error) {
	user := model.User{}

	result := r.db.Model(&user).Where(constants.WHERE_BY_ID, id).Clauses(clause.Returning{}).Updates(model.User{
		ID:          id,
		Username:    payload.Username,
		Email:       payload.Email,
		Password:    payload.Password,
		FirstName:   payload.FirstName,
		LastName:    payload.LastName,
		PhoneNumber: payload.PhoneNumber,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *userRepository) Delete(id string) (*model.User, error) {
	user := model.User{}

	if err := r.db.Clauses(clause.Returning{}).Where(constants.WHERE_BY_ID, id).Delete(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
