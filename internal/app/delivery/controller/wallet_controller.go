package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sakupay-apps/internal/app/service"
	"github.com/sakupay-apps/internal/model"
	"github.com/sakupay-apps/utils/common"
)

type WalletController struct {
	service service.WalletService
}

func (w *WalletController) CreateHandler(c *gin.Context) {
	var wallet model.Wallet
	wallet.ID = common.GenerateUUID()
	if err := c.ShouldBindJSON(&wallet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := w.service.RegisterNewWallet(wallet)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Success Create New Wallet",
		"data":    wallet,
	})
}

func NewWalletController(wlService service.WalletService) *WalletController {
	return &WalletController{
		service: wlService,
	}
}
