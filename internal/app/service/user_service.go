package service

import (
	"github.com/sakupay-apps/internal/app/repository"
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"
	"github.com/sakupay-apps/utils/exception"
	"gorm.io/gorm"
)

type UserService interface {
	RegisterNewUser(payload *model.User) (*dto.UserResponse, error)
	FindUserByID(id string) (*dto.UserResponse, error)
	FindAllUser(requestPaging dto.PaginationParam, queries ...string) ([]*dto.UserResponse, *dto.Paging, error)
	UpdateUserByID(id string, payload *model.User) (*dto.UserResponse, error)
	RemoveUser(id string) (*dto.UserResponse, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterNewUser(payload *model.User) (*dto.UserResponse, error) {

	user, err := s.repo.Create(payload)

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	userResponse := dto.UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}

	return &userResponse, err
}

func (s *userService) FindUserByID(id string) (*dto.UserResponse, error) {

	user, err := s.repo.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	userResponse := dto.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		LastLogin: user.LastLogin,
	}

	return &userResponse, err
}

func (s *userService) FindAllUser(requestPaging dto.PaginationParam, queries ...string) ([]*dto.UserResponse, *dto.Paging, error) {

	users, paging, err := s.repo.List(requestPaging, queries...)

	if err != nil {
		return nil, nil, gorm.ErrRecordNotFound
	}

	var userResponses []*dto.UserResponse

	for _, user := range users {
		userResponse := dto.UserResponse{
			ID:          user.ID,
			Username:    user.Username,
			Email:       user.Email,
			Password:    user.Password,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			PhoneNumber: user.PhoneNumber,
			LastLogin:   user.LastLogin,
		}

		userResponses = append(userResponses, &userResponse)
	}

	return userResponses, paging, err
}

func (s *userService) RemoveUser(id string) (*dto.UserResponse, error) {

	user, err := s.repo.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	user, err = s.repo.Delete(user.ID)

	if err != nil {
		return nil, exception.ErrFailedDelete
	}

	userResponse := dto.UserResponse{
		ID:               user.ID,
		Username:         user.Username,
		Email:            user.Email,
		Password:         user.Password,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		PhoneNumber:      user.PhoneNumber,
		RegistrationDate: user.RegistrationDate,
	}

	return &userResponse, err
}

func (s *userService) UpdateUserByID(id string, payload *model.User) (*dto.UserResponse, error) {

	user, err := s.repo.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	user, err = s.repo.Update(user.ID, payload)

	if err != nil {
		return nil, exception.ErrFailedUpdate
	}

	userResponse := dto.UserResponse{
		ID:               user.ID,
		Username:         user.Username,
		Email:            user.Email,
		Password:         user.Password,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		PhoneNumber:      user.PhoneNumber,
		RegistrationDate: user.RegistrationDate,
	}

	return &userResponse, err
}
