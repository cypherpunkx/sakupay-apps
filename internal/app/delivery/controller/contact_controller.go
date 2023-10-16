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

type ContactController struct {
	service service.ContactService
}

func NewContactController(service service.ContactService) *ContactController {
	return &ContactController{
		service: service,
	}
}

func (cc *ContactController) AddContact(c *gin.Context) {
	var payload model.Contact

	payload.ID = common.GenerateUUID()
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"status":  exception.StatusBadRequest,
			"message": exception.FieldErrors(err),
		})
		return
	}
	data, err := cc.service.RegisterNewContact(&payload)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusInternalServerError, dto.ErrorResponse{
				Code:    http.StatusInternalServerError,
				Status:  exception.StatusInternalServer,
				Message: gorm.ErrRecordNotFound.Error(),
			})
			return
		}

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

	c.JSON(http.StatusCreated, dto.Response{
		Code:    http.StatusCreated,
		Status:  exception.StatusSuccess,
		Message: "Get All Contacts",
		Data:    data,
	})
}

// func (cc *ContactController) ListHandler(c *gin.Context) {

// 	page, _ := strconv.Atoi(c.Query("page"))
// 	limit, _ := strconv.Atoi(c.Query("limit"))
// 	name := c.Query("name")

// 	paginationParam := dto.PaginationParam{
// 		Page:  page,
// 		Limit: limit,
// 	}

// 	contact, paging, err := cc.service.FindAllContactList(paginationParam, name)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	status := map[string]any{
// 		"code":        200,
// 		"description": "Get All Data Successfully",
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"status": status,
// 		"data":   contact,
// 		"paging": paging,
// 	})
// }

// func (cc *ContactController) GetHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	data, err := cc.service.FindContactById(id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  http.StatusOK,
// 		"message": "Success Get Contact by Id",
// 		"data":    data,
// 	})
// 	return

// }

// func (cc *ContactController) DeleteHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	data, err := cc.service.DeleteContact(id)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  http.StatusOK,
// 		"message": "Success Delete Contact",
// 		"data":    data,
// 	})
// }

// func (cc *ContactController) UpdateHandler(c *gin.Context) {
// 	id := c.Param("id")
// 	payload := model.Contact{}

// 	if err := c.ShouldBindJSON(&payload); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	data, err := cc.service.UpdateContact(id, &payload)

// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusCreated, gin.H{
// 		"status":  http.StatusCreated,
// 		"message": "Success Updated Contact",
// 		"data":    data,
// 	})
// }
