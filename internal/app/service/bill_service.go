package service

import (
	"github.com/sakupay-apps/internal/app/repository"
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"
	"github.com/sakupay-apps/utils/common"
	"github.com/sakupay-apps/utils/exception"
	"gorm.io/gorm"
)

type BillService interface {
	CreateNewBill(payload *model.Bill) (*dto.BillResponse, error)
	FindBillByID(user_ID, bill_ID string) (*dto.BillResponse, error)
}

type billService struct {
	billRepo        repository.BillRepository
	userRepo        repository.UserRepository
	billDetailsRepo repository.BillDetailsRepository
}

func (b *billService) CreateNewBill(payload *model.Bill) (*dto.BillResponse, error) {

	user, err := b.userRepo.Get(payload.UserID)

	if err != nil {
		return nil, exception.ErrNotFound
	}

	billdetails := []model.BillDetails{}

	for _, billdetail := range payload.Billdetails {

		billdetail.ID = common.GenerateUUID()
		billdetail.BillID = payload.ID

		billdetails = append(billdetails, billdetail)
	}

	payload.Billdetails = billdetails

	bill, err := b.billRepo.Create(payload)

	if err != nil {
		return nil, exception.ErrFailedCreate
	}
	billResponse := dto.BillResponse{

		ID: bill.ID,
		User: model.User{
			ID:          user.ID,
			Username:    user.Username,
			Email:       user.Email,
			Password:    user.Password,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			PhoneNumber: user.PhoneNumber,
		},
		BillDetails: bill.Billdetails,
		Total:       bill.Total,
		DueDate:     bill.DueDate,
	}

	return &billResponse, nil
}

func (b *billService) FindBillByID(user_ID, bill_ID string) (*dto.BillResponse, error) {

	user, err := b.userRepo.Get(user_ID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	bill, err := b.billRepo.GetBillIdAndUserId(user.ID, bill_ID)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	billResponse := dto.BillResponse{

		ID:          bill.ID,
		User:        *user,
		BillDetails: bill.Billdetails,
		Total:       bill.Total,
		DueDate:     bill.DueDate,
	}

	return &billResponse, err
}

func NewBillService(billRepo repository.BillRepository, userRepo repository.UserRepository) BillService {
	return &billService{
		billRepo: billRepo,
		userRepo: userRepo,
	}
}
