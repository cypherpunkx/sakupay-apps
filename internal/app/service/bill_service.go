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
	billRepo repository.BillRepository
	userRepo repository.UserRepository
}

func NewBillService(billRepo repository.BillRepository, userRepo repository.UserRepository) BillService {
	return &billService{
		billRepo: billRepo,
		userRepo: userRepo,
	}
}

func (b *billService) CreateNewBill(payload *model.Bill) (*dto.BillResponse, error) {

	user, err := b.userRepo.Get(payload.UserID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	bill, err := b.billRepo.Create(payload)

	if err != nil {
		return nil, exception.ErrFailedCreate
	}

	billResponse := dto.BillResponse{
		ID:          bill.ID,
		User:        *user,
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

	bills, err := s.billRepo.List(user.ID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	billResponses := []*dto.BillResponse{}

	for _, bill := range bills {
		if bill.UserID == user.ID {
			billResponses = append(billResponses, &dto.BillResponse{
				ID:          bill.ID,
				User:        *user,
				BillDetails: bill.Billdetails,
				Total:       bill.Total,
				DueDate:     bill.DueDate,
			})
		}
	}

	return billResponses, nil

}
