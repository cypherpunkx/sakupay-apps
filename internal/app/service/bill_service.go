package service

import (
	"github.com/sakupay-apps/internal/app/repository"
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"
	"github.com/sakupay-apps/utils/exception"
	"gorm.io/gorm"
)

type BillService interface {
	CreateNewBill(payload *model.Bill) (*dto.BillResponse, error)
	GetAllBills(id string) ([]*dto.BillResponse, error)
}

type billService struct {
	billRepo        repository.BillRepository
	billDetailsRepo repository.BillDetailsRepository
	userRepo        repository.UserRepository
}

func NewBillService(billRepo repository.BillRepository, userRepo repository.UserRepository, billDetailsRepo repository.BillDetailsRepository) BillService {
	return &billService{
		billRepo:        billRepo,
		userRepo:        userRepo,
		billDetailsRepo: billDetailsRepo,
	}
}

func (b *billService) CreateNewBill(payload *model.Bill) (*dto.BillResponse, error) {

	user, err := b.userRepo.Get(payload.UserID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	if payload.Total > user.Wallet.Balance {
		return nil, exception.ErrNotEnoughBalance
	}

	bill, err := b.billRepo.Create(payload)

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	billResponse := dto.BillResponse{
		ID: bill.ID,
		User: model.User{
			ID:       user.ID,
			Username: user.Username,
			Wallet: model.Wallet{
				Name:    user.Wallet.Name,
				Balance: user.Wallet.Balance,
			},
		},
		BillDetails: bill.Billdetails,
		Total:       bill.Total,
		DueDate:     bill.DueDate,
	}

	return &billResponse, nil
}

func (s *billService) GetAllBills(id string) ([]*dto.BillResponse, error) {
	user, err := s.userRepo.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	bill, err := s.billRepo.Get(user.ID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	billDetail, err := s.billDetailsRepo.Get(bill.ID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	bills, err := s.billRepo.List(user.ID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	billResponses := []*dto.BillResponse{}

	for _, bill := range bills {
		if bill.UserID == user.ID && bill.ID == billDetail.BillID {
			billResponses = append(billResponses, &dto.BillResponse{
				ID: bill.ID,
				User: model.User{
					ID:       user.ID,
					Username: user.Username,
					Wallet: model.Wallet{
						Name:    user.Wallet.Name,
						Balance: user.Wallet.Balance,
					},
				},
				BillDetails: *billDetail,
				Total:       bill.Total,
				DueDate:     bill.DueDate,
			})
		}
	}

	return billResponses, nil

}
