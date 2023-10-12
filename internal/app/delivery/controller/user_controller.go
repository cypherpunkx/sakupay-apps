package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sakupay-apps/internal/app/service"
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/internal/model/dto"
	"github.com/sakupay-apps/utils/common"
	"github.com/sakupay-apps/utils/exception"
	"gorm.io/gorm"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}

func (ctr *UserController) Registration(c *gin.Context) {
	payload := model.User{}

	payload.ID = common.GenerateUUID()

	fmt.Println(payload)

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  exception.StatusBadRequest,
			"message": exception.FieldErrors(err),
		})
		return
	}

	data, err := ctr.service.RegisterNewUser(&payload)

	if err != nil {
		if errors.Is(err, exception.ErrFailedCreate) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServer,
				Message: exception.ErrFailedCreate,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
			Code:    http.StatusInternalServerError,
			Status:  exception.StatusInternalServer,
			Message: err.Error(),
		})
	}

	c.JSON(http.StatusCreated, dto.Response{
		Code:    http.StatusCreated,
		Status:  exception.StatusSuccess,
		Message: "Get All Users",
		Data:    data,
	})
}

func (ctr *UserController) FindUser(c *gin.Context) {
	id := c.Param("id")

	data, err := ctr.service.FindUserByID(id)

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
	}

	c.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Status:  exception.StatusSuccess,
		Message: "Get User By ID",
		Data:    data,
	})
}

func (ctr *UserController) FindUsers(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    http.StatusBadRequest,
			Status:  exception.StatusInternalServer,
			Message: exception.ErrInvalidPage.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "5"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResponse{
			Code:    http.StatusBadRequest,
			Status:  exception.StatusInternalServer,
			Message: exception.ErrInvalidPage.Error(),
		})
		return
	}

	paginationParam := dto.PaginationParam{
		Page:  page,
		Limit: limit,
	}

	data, paging, err := ctr.service.FindAllUser(paginationParam)

	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
	// 			Code:    http.StatusInternalServerError,
	// 			Status:  exception.StatusInternalServer,
	// 			Message: gorm.ErrRecordNotFound.Error(),
	// 		})
	// 		return
	// 	}

	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
	// 		Code:    http.StatusInternalServerError,
	// 		Status:  exception.StatusInternalServer,
	// 		Message: err.Error(),
	// 	})
	// }

	c.JSON(http.StatusOK, dto.Response{
		Code:    http.StatusOK,
		Status:  exception.StatusSuccess,
		Message: "Get User By ID",
		Data:    data,
		Paging:  *paging,
	})
}
