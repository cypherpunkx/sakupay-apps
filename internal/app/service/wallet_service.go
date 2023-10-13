package service

import (
	"github.com/sakupay-apps/internal/app/repository"
	"github.com/sakupay-apps/internal/model"
)

type WalletService interface {
	RegisterNewWallet(payload model.Wallet) error
}

type walletService struct {
	repo repository.WalletRepository
}

func (w *walletService) RegisterNewWallet(payload model.Wallet) error {
	return w.repo.Create(payload)
}

func NewWalletService(wlRepo repository.WalletRepository) WalletService {
	return &walletService{
		repo: wlRepo,
	}
}
