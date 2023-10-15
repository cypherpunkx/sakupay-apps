package service

import (
	"time"

	"github.com/sakupay-apps/internal/app/repository"
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"
	"github.com/sakupay-apps/utils/exception"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	RegisterNewUser(payload *model.User) (*dto.UserResponse, error)
	FindUserByID(id string) (*dto.UserResponse, error)
	FindAllUser(requestPaging dto.PaginationParam, queries ...string) ([]*dto.UserResponse, *dto.Paging, error)
	UpdateUserByID(id string, payload *model.User) (*dto.UserResponse, error)
	RemoveUser(id string) (*dto.UserResponse, error)
	FindByUsername(username string) (*model.User, error)
	FindByUsernamePassword(username string, password string) (*model.User, error)
}
type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) RegisterNewUser(payload *model.User) (*dto.UserResponse, error) {

	users, _, err := s.repo.List(dto.PaginationParam{})

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	for _, user := range users {
		if user.Username == payload.Username {
			return nil, exception.ErrUsernameAlreadyExist
		}
		if user.Email == payload.Email {
			return nil, exception.ErrEmailAlreadyExist
		}
		if user.PhoneNumber == payload.PhoneNumber {
			return nil, exception.ErrPhoneNumberAlreadyExist
		}
	}

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	password := string(bytes)

	payload.Password = password

	user, err := s.repo.Create(payload)

	userResponse := dto.UserResponse{
		ID:               user.ID,
		Username:         user.Username,
		Email:            user.Email,
		Password:         user.Password,
		FirstName:        user.FirstName,
		LastName:         user.Username,
		PhoneNumber:      user.PhoneNumber,
		Wallet:           user.Wallet,
		RegistrationDate: user.RegistrationDate,
		ProfilePicture:   user.ProfilePicture,
		LastLogin:        time.Now(),
	}

	return &userResponse, err
}

func (s *userService) FindUserByID(id string) (*dto.UserResponse, error) {

	user, err := s.repo.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	userResponse := dto.UserResponse{
		ID:               user.ID,
		Username:         user.Username,
		Email:            user.Email,
		Password:         user.Password,
		FirstName:        user.FirstName,
		LastName:         user.Username,
		PhoneNumber:      user.PhoneNumber,
		Wallet:           user.Wallet,
		RegistrationDate: user.RegistrationDate,
		ProfilePicture:   user.ProfilePicture,
		LastLogin:        user.LastLogin,
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
			ID:               user.ID,
			Username:         user.Username,
			Email:            user.Email,
			Password:         user.Password,
			FirstName:        user.FirstName,
			LastName:         user.Username,
			PhoneNumber:      user.PhoneNumber,
			Wallet:           user.Wallet,
			RegistrationDate: user.RegistrationDate,
			ProfilePicture:   user.ProfilePicture,
			LastLogin:        user.LastLogin,
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
		LastName:         user.Username,
		PhoneNumber:      user.PhoneNumber,
		Wallet:           user.Wallet,
		RegistrationDate: user.RegistrationDate,
		ProfilePicture:   user.ProfilePicture,
		LastLogin:        user.LastLogin,
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
		LastName:         user.Username,
		PhoneNumber:      user.PhoneNumber,
		Wallet:           user.Wallet,
		RegistrationDate: user.RegistrationDate,
		ProfilePicture:   user.ProfilePicture,
		LastLogin:        user.LastLogin,
	}

	return &userResponse, err
}

func (s *userService) FindByUsername(username string) (*model.User, error) {
	return s.repo.GetUsername(username)
}

func (s *userService) FindByUsernamePassword(username string, password string) (*model.User, error) {
	return s.repo.GetUsernamePassword(username, password)
}
