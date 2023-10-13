package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sakupay-apps/config"
	"github.com/sakupay-apps/internal/app/delivery/controller"
	"github.com/sakupay-apps/internal/app/repository"
	"github.com/sakupay-apps/internal/app/service"
)

func SetupRouter(router *gin.Engine) error {

	v1 := router.Group("/api/v1")
	{
		sakupay := v1.Group("/sakupay")
		{
			users := sakupay.Group("/users")
			{
				users.GET("/", func(c *gin.Context) {
					c.String(200, "ok")
				})
			}
			wallet := sakupay.Group("/wallet")
			walletRepo := repository.NewWalletRepository(config.DB)
			walletService := service.NewWalletService(walletRepo)
			controller := controller.NewWalletController(walletService)
			{
				wallet.POST("/", controller.CreateHandler)
			}
		}
	}

	return router.Run()
}
