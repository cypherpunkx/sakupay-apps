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
			userRepo := repository.NewUserRepository(config.DB)
			userService := service.NewUserService(userRepo)
			controller := controller.NewUserController(userService)

			{
				users.GET("/", controller.FindUsers)
				users.POST("/", controller.Registration)
				users.GET("/:id", controller.FindUser)
				users.PUT("/:id", controller.UpdatingUser)
				users.DELETE("/:id", controller.DeletedUser)
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
