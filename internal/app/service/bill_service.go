package service

import (
	"github.com/sakupay-apps/internal/app/repository"
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"
	"github.com/sakupay-apps/utils/exception"
)

type BillService interface {
	CreateNewBill(payload *model.Bill) (*dto.BillResponse, error)
}

type billService struct {
	billRepo repository.BillRepository
	userRepo repository.UserRepository
}

func NewBillRepository(billRepo repository.BillRepository, userRepo repository.UserRepository) BillService {
	return &billService{
		billRepo: billRepo,
		userRepo: userRepo,
	}
}

func (s *billService) CreateNewBill(payload *model.Bill) (*dto.BillResponse, error) {

	bill, err := s.billRepo.Create(payload)

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	user, err := s.userRepo.Get(bill.UserID)

	if err != nil {
		return nil, exception.ErrNotFound
	}

	billResponse := dto.BillResponse{
		ID: bill.ID,
		User: model.User{
			ID:       user.ID,
			Username: user.Username,
		},
		Total:   bill.Total,
		DueDate: bill.DueDate,
	}

	return &billResponse, nil
}
