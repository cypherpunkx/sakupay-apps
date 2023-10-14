package repository

import (
	"github.com/sakupay-apps/internal/model"
	// "github.com/sakupay-apps/internal/model/dto"
	// "github.com/sakupay-apps/utils/common"

	// "github.com/sakupay-apps/internal/model/dto"
	"gorm.io/gorm"
)

type BillRepository interface {
	Create(payload *model.Bill) (*model.Bill,error)
	Get(id string) (*model.Bill, error)
	// GetBillDetailById(id string) (model.Bill, error)
}

type billRepository struct {
	db *gorm.DB
}

func (b *billRepository) Create(payload *model.Bill) (*model.Bill,error) {
	bill := model.Bill {
		ID : payload.ID,
		UserID: payload.UserID,
		BilldetailsID: payload.BilldetailsID,
		Total: payload.Total,
		DueDate: payload.DueDate,

	}

	b.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&bill).Error; err != nil {
			return err
		}
		return nil
	})
	return  &bill,nil
}

func (b *billRepository) Get(id string) (*model.Bill, error) {
	bill := model.Bill{}

	if err := b.db.Where("WHERE id = ? ", id).Find(&bill).Error; err != nil {
		return nil, err
	}

	return &bill, nil
}


// func (r *billRepository) List(requestPaging dto.PaginationParam, queries ...string) ([]*model.Bill, *dto.Paging, error) {
// 	bills := []*model.Bill{}

// 	paginationQuery := common.GetPaginationParams(requestPaging)

// 	var totalRows int64

// 	if err := r.db.Limit(paginationQuery.Take).Offset(paginationQuery.Skip).Find(&bills).Count(&totalRows).Error; err != nil {
// 		return nil, nil, err
// 	}

// 	var count int = int(totalRows)

// 	return bills, common.Paginate(paginationQuery.Take, paginationQuery.Page, count), nil
// }

func NewBillRepository(db *gorm.DB) BillRepository {
	return &billRepository{
		db: db,
	}
}
