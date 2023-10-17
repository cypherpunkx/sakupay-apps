package service

import (
	"fmt"
	"mime/multipart"

	"github.com/sakupay-apps/internal/app/repository"
	"github.com/sakupay-apps/internal/model"
)

type UserPictureService interface {
	UploadUserPicture(userPicture model.UserPicture, file *multipart.File, extFile string) error
	FindUserPictureById(id string) (model.UserPicture, error)
}

type userPictureService struct {
	userRepoPic repository.UserPictureRepository
	fileRepo    repository.FileRepository
	userService UserService
}

func NewUserPictureService(userRepoPic repository.UserPictureRepository, fileRepo repository.FileRepository, userService UserService) UserPictureService {
	return &userPictureService{
		fileRepo:    fileRepo,
		userService: userService,
		userRepoPic: userRepoPic,
	}
}

func (u *userPictureService) UploadUserPicture(userPicture model.UserPicture, file *multipart.File, extFile string) error {
	userCred, err := u.userService.FindUserByID(userPicture.UserID)
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s-%s%s", userCred.Username, userPicture.UserID, extFile)

	filePath, err := u.fileRepo.Save(fileName, file)
	if err != nil {
		return err
	}
	userPicture.FileLocation = filePath
	err = u.userRepoPic.Create(userPicture)
	if err != nil {
		return fmt.Errorf("Failed Upload : %s", err.Error())
	}
	return nil
}

func (u *userPictureService) FindUserPictureById(id string) (model.UserPicture, error) {
	userPicture, err := u.userRepoPic.Get(id)
	if err != nil {
		return model.UserPicture{}, fmt.Errorf("Failed Get Picture By Id : %s", err.Error())
	}
	return userPicture, nil
}
