package service

import (
	"github.com/sakupay-apps/internal/app/repository"
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"
	"github.com/sakupay-apps/utils/exception"
)

type WalletService interface {
	RegisterNewWallet(payload *model.Wallet) (*dto.WalletResponse, error)
}

type walletService struct {
	walletRepo repository.WalletRepository
	userRepo   repository.UserRepository
}

func (w *walletService) RegisterNewWallet(payload *model.Wallet) (*dto.WalletResponse, error) {
	wallet, err := w.walletRepo.Create(payload)

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	user, err := w.userRepo.Get(wallet.UserID)
	if err != nil {
		return nil, err
	}
	walletResponse := dto.WalletResponse{
		ID:   wallet.ID,
		Name: wallet.Name,
		User: model.User{
			ID:          user.ID,
			Username:    user.Username,
			Email:       user.Email,
			Password:    user.Password,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			PhoneNumber: user.PhoneNumber,
		},
		Balance: wallet.Balance,
	}

	return &walletResponse, err
}

func NewWalletService(wlRepo repository.WalletRepository, usrRepo repository.UserRepository) WalletService {
	return &walletService{
		walletRepo: wlRepo,
		userRepo:   usrRepo,
	}
}
