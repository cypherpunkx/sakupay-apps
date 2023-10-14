package service

import (
	// "fmt"

	"github.com/sakupay-apps/internal/app/repository"
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"
	"github.com/sakupay-apps/utils/exception"

	// "github.com/sakupay-apps/utils/common"
	// "github.com/sakupay-apps/utils/exception"
	"gorm.io/gorm"
)

type BillService interface {
	CreateNewBill(payload *model.Bill) (*dto.BillResponse, error)
	FindBillByID(id string) (*dto.BillResponse, error)
	// FindAllBill(requestPaging dto.PaginationParam, queries ...string) ([]dto.UserResponse, *dto.Paging, error)
}

type billService struct {
	billRepo repository.BillRepository
	userRepo repository.UserRepository
}

func (b *billService) CreateNewBill (payload *model.Bill) (*dto.BillResponse, error) {
	bill,err := b.repository.Create(payload) 

	if err != nil {
		return nil,exception.ErrFailedCreate
	}
	user,err := b.repository.Get(bill.UserID)

	if err != nil {
		return nil , exception.ErrNotFound
	}

		billResponse := dto.BillResponse {
			ID : bill.ID,
			User: model.User {
				ID: 	   user.ID,
				Username : user.Username,
			},
			Total : bill.Total,
			DueDate: bill.DueDate,
		}

	return &billResponse, nil
}

// func (n *billService) FindBillById(id string) (*model.Bill, error) {
// 	bill, err := n.repository.Get(id)
// 	if err != nil {
// 		return model.Bill{}, err
// 	}

// 	return bill, nil
// }

// func (b *billService) CreateBill(payload model.Bill) error {
// 	// user, err := b.userService.(payload.UserId)
// 	// if err != nil {
// 	// 	return fmt.Errorf("User with id %s is not found", payload.UserId)
// 	// }

// 	// newBillDetail := make([]model.BillDetails,0,len(payload.BillDetails))
// 	// for _, billDetail := range payload.BillDetails {
// 	// 	billDetail.ID = common.GenerateUUID()
// 	// 	billDetail.BillID = payload.ID
// 	// 	// billDetail.Name = billDetail.Name
// 	// 	// billDetail.Category = billDetail.Category
// 	// 	// billDetail.Description = billDetail.Description
// 	// 	// billDetail.Website = billDetail.Website
// 	// 	newBillDetail = append(newBillDetail, billDetail)
// 	// }

// 	// payload.ID = common.GenerateUUID()
// 	// // payload.UserID = user.ID
// 	// payload.BillDetails = newBillDetail
// 	// // payload.Total = payload.Total
// 	// payload.DueDate= time.Now()
// 	payload.ID = common.GenerateUUID()
// 	for _,billDetail := range payload.BillDetails {
// 		billDetail.BillID = common.GenerateUUID()
// 	}
// 	err := b.repository.Create()
// 	if err != nil {
// 		return fmt.Errorf("Failed to register new bill : %v",err)
// 	}
// 	return nil
// }

func (b *billService) FindBillByID(id string) (*dto.BillResponse, error) {

	bill, err := b.repository.Get(id)

	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	billResponse := dto.BillResponse {
		ID     		:bill.ID,
		User   		:model.User{},
		BillDetails :bill.BillDetails,
		Total		:bill.Total,
		DueDate		:bill.DueDate,
	}

	return &billResponse, err
}

// func (b *billService) FindAllBill(requestPaging dto.PaginationParam, queries ...string) ([]dto.UserResponse, *dto.Paging, error) {

// 	bills, paging, err := b.repository.List(requestPaging, queries...)

// 	if err != nil {
// 		return nil, nil, gorm.ErrRecordNotFound
// 	}

// 	var billResponses []dto.BillResponse

// 	for _, bill := range bills {
// 		billResponse := dto.BillResponse{
// 			ID:          bill.ID,
			
// 		}

// 		billResponses = append(billResponses, billResponse)
// 	}

// 	return []dto.UserResponse{}, paging, err
// }


func NewBillService(billRepo repository.BillRepository) BillService {
	return &billService{
		repository: billRepo,
	}
}
