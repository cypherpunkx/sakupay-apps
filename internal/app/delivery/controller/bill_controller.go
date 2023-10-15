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
)

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type BillController struct {
	billservice service.BillService
	userService	service.UserService
}

func (b *BillController) CreateNewBill(c *gin.Context) {

	id := c.Param("id")

	var bill model.Bill
	bill.ID = common.GenerateUUID()
	bill.UserID = id

	if err := c.ShouldBindJSON(&bill); err != nil {

		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"code" 	  :http.StatusBadRequest,  
			"status"  :exception.StatusBadRequest,
			"message" :exception.FieldErrors(err),
		})
		return
	}

	data,err := b.billservice.CreateNewBill(&bill)
	
	if err != nil {
		if errors.Is(err, exception.ErrFailedCreate) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServer,
				Message: exception.ErrFailedCreate.Error(),
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
	
	c.JSON(http.StatusCreated, response{
		Status	: http.StatusCreated,
		Message	: "Success Create Bills",
		Data	: data,
	})
}

func NewBillController(billService service.BillService, userService service.UserService) *BillController{

	return  &BillController{
		billservice	  : billService,
		userService: userService,
		}
}