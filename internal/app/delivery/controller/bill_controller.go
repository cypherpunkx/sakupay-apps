package controller

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/service"
	"github.com/sakupay-apps/utils/exception"
	// "github.com/sakupay-apps/utils/common"
	// "github.com/sakupay-apps/utils/exception"
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
	service service.BillService
}

func (b *BillController) CreateNewBill(c *gin.Context) {
	var bill model.Bill
	// bill.ID = common.GenerateUUID()
	// bill.UserID = common.GenerateUUID()
	// bill.BilldetailsID = common.GenerateUUID()
	// bill.DueDate = time.Now()
	if err := c.ShouldBindJSON(&bill); err != nil {

		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"code" 	  :http.StatusBadRequest,  
			"status"  :http.StatusBadRequest,
			"message" :err.Error(),
		})
		return
	}

	_,err := b.service.CreateNewBill(&bill)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse{
			Status  :  http.StatusInternalServerError,
			Message : err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusCreated, response{
		Status	: http.StatusCreated,
		Message	: "Success Create Bills",
		Data	: bill,
	})
}

func (b *BillController) GetDetailBill(c *gin.Context) {
	id := c.Param("id")

	bill, err := b.service.FindBillByID(id)
	if err != nil {
		if strings.Contains(err.Error(), exception.ErrNotFound.Error()) {
			c.JSON(http.StatusNotFound, errorResponse{
				Status:  http.StatusNotFound,
				Message: err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response{
		Status:  http.StatusOK,
		Message: "Success Get Detail Bill",
		Data:    bill,
	})
}

func NewBillController(service service.BillService) *BillController{
	controller := &BillController{
	service	   : service,
	}

	return controller
}