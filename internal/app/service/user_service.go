package service

import (
	"github.com/sakupay-apps/internal/app/repository"
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"
	"github.com/sakupay-apps/utils/exception"
)

type UserService interface {
	RegisterNewUser(payload *model.User) (*dto.UserResponse, error)
	FindUserByID(id string) (*dto.UserResponse, error)
	FindAllUser(requestPaging dto.PaginationParam, queries ...string) ([]*dto.UserResponse, *dto.Paging, error)
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

	return &userResponse, nil
}

func (s *userService) FindUserByID(id string) (*dto.UserResponse, error) {

	user, err := s.repo.Get(id)

	if err != nil {
		return nil, err
	}

	userResponse := dto.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		LastLogin: user.LastLogin,
	}

	return &userResponse, nil
}

func (s *userService) FindAllUser(requestPaging dto.PaginationParam, queries ...string) ([]*dto.UserResponse, *dto.Paging, error) {

	users, paging, err := s.repo.List(requestPaging, queries...)

	if err != nil {
		return nil, nil, err
	}

	var userResponses []*dto.UserResponse

	for _, user := range users {
		userResponse := dto.UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			LastLogin: user.LastLogin,
		}

		userResponses = append(userResponses, &userResponse)
	}

	return userResponses, paging, nil
}
