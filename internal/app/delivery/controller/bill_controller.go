package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sakupay-apps/internal/app/service"
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"
	"github.com/sakupay-apps/utils/common"
	"github.com/sakupay-apps/utils/exception"
	"gorm.io/gorm"
)

type billController struct {
	billService service.BillService
}

func NewBillController(billService service.BillService) *billController {
	return &billController{
		billService: billService,
	}
}

func (b *billController) CreateBill(c *gin.Context) {
	id := c.Param("id")

	payload := model.Bill{}

	payload.ID = common.GenerateUUID()
	payload.UserID = id
	payload.Billdetails.ID = common.GenerateUUID()
	payload.Billdetails.BillID = payload.ID

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  exception.StatusBadRequest,
			"message": exception.FieldErrors(err),
		})
		return
	}

	data, err := b.billService.CreateNewBill(&payload)

	if err != nil {
		if errors.Is(err, exception.ErrFailedCreate) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServer,
				Message: exception.ErrFailedCreate.Error(),
			})
			return
		}

		if errors.Is(err, gorm.ErrInvalidTransaction) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServer,
				Message: gorm.ErrInvalidTransaction.Error(),
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Status:  exception.StatusInternalServer,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Code:    http.StatusCreated,
		Status:  exception.StatusSuccess,
		Message: "Create Bill",
		Data:    data,
	})
}

func (b *billController) FindAllBills(c *gin.Context) {
	id := c.Param("id")

	data, err := b.billService.GetAllBills(id)

	if err != nil {
		if errors.Is(err, exception.ErrFailedCreate) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:   http.StatusInternalServerError,
				Status: exception.StatusInternalServer,

				Message: exception.ErrFailedCreate.Error(),
			})
			return
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:   http.StatusInternalServerError,
				Status: exception.StatusInternalServer,

				Message: gorm.ErrRecordNotFound.Error(),
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Status:  exception.StatusInternalServer,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Status:  exception.StatusSuccess,
		Message: "Get All Bills",
		Data:    data,
	})
}
