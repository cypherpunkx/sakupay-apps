package service

import (
	"github.com/sakupay-apps/utils/exception"
	"github.com/sakupay-apps/utils/security"
)

type AuthService interface {
	Login(username string, password string) (string, error)
}

type authService struct {
	service UserService
}

func NewAuthService(service UserService) AuthService {
	return &authService{service: service}
}

func (s *authService) Login(username string, password string) (string, error) {
	user, err := s.service.FindByUsernamePassword(username, password)

	if err != nil {
		return "", err
	}

	token, err := security.CreateAccessToken(user)

	if err != nil {
		return "", exception.ErrFailedCreateToken
	}

	return token, nil
}
