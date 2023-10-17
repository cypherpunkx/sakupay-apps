package repository

import (
	"github.com/sakupay-apps/internal/model"
	"gorm.io/gorm"
)

type UserPictureRepository interface {
	Create(userPicture *model.UserPicture) error
	Get(id string) (*model.UserPicture, error)
}

type userPicture struct {
	db *gorm.DB
}

func NewUserPictureRepository(db *gorm.DB) UserPictureRepository {
	return &userPicture{
		db: db,
	}
}

func (u *userPicture) Create(userPicture *model.UserPicture) error {
	userpc := model.UserPicture{}

	if err := u.db.Create(&userpc).Error; err != nil {
		return err
	}
	return nil
}

func (u *userPicture) Get(id string) (*model.UserPicture, error) {
	userPicure := model.UserPicture{}

	if err := u.db.Where("id = ?", id).First(&userPicure).Error; err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	return &userPicure, nil
}
