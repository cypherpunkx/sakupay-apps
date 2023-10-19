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

type transactionController struct {
	service service.TransactionService
}

func NewTransactionController(service service.TransactionService) *transactionController {
	return &transactionController{service: service}
}

func (ctr *transactionController) Deposit(c *gin.Context) {
	id := c.Param("id")
	cardID := c.Params.ByName("cardID")

	payload := model.Transaction{}

	payload.ID = common.GenerateUUID()
	payload.UserID = id

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  exception.StatusBadRequest,
			"message": exception.FieldErrors(err),
		})
		return
	}

	data, err := ctr.service.DepositMoney(&payload, cardID)

	if err != nil {
		if errors.Is(err, exception.ErrFailedCreate) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServer,
				Message: exception.ErrFailedCreate.Error(),
			})
			return
		}

		if errors.Is(err, exception.ErrMinimalTransaction) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServer,
				Message: exception.ErrMinimalTransaction.Error(),
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
		Message: "Create Transaction",
		Data:    data,
	})
}

func (ctr *transactionController) FindAllTransactions(c *gin.Context) {
	id := c.Param("id")

	data, err := ctr.service.TransactionHistory(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServer,
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
		Message: "Get All Transactions",
		Data:    data,
	})
}

func (ctr *transactionController) FindTransaction(c *gin.Context) {
	id := c.Params.ByName("id")
	transactionID := c.Params.ByName("transactionID")

	data, err := ctr.service.FindTransactionByUser(id, transactionID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServer,
				Message: gorm.ErrRecordNotFound.Error(),
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

	c.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Status:  exception.StatusSuccess,
		Message: "Get Transaction By ID ",
		Data:    data,
	})
}

func (ctr *transactionController) SendMoney(c *gin.Context) {
	id := c.Params.ByName("id")
	friendID := c.Params.ByName("userID")

	payload := model.Transaction{}

	payload.ID = common.GenerateUUID()
	payload.UserID = id

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  exception.StatusBadRequest,
			"message": exception.FieldErrors(err),
		})
		return
	}

	data, err := ctr.service.SendMoney(&payload, friendID)

	if err != nil {
		if errors.Is(err, exception.ErrFailedCreate) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServer,
				Message: exception.ErrFailedCreate.Error(),
			})
			return
		}

		if errors.Is(err, exception.ErrMinimalTransaction) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServer,
				Message: exception.ErrMinimalTransaction.Error(),
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
		Message: "Create Transaction",
		Data:    data,
	})
}

func (ctr *transactionController) WithdrawMoney(c *gin.Context) {
	id := c.Params.ByName("id")
	cardID := c.Params.ByName("cardID")

	payload := model.Transaction{}

	payload.ID = common.GenerateUUID()
	payload.UserID = id

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  exception.StatusBadRequest,
			"message": exception.FieldErrors(err),
		})
		return
	}

	data, err := ctr.service.WithdrawMoneyToCard(&payload, cardID)

	if err != nil {
		if errors.Is(err, exception.ErrFailedCreate) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServer,
				Message: exception.ErrFailedCreate.Error(),
			})
			return
		}

		if errors.Is(err, exception.ErrMinimalTransaction) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServer,
				Message: exception.ErrMinimalTransaction.Error(),
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
		Message: "Create Transaction",
		Data:    data,
	})
}
