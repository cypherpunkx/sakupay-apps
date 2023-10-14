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
			uesrController := controller.NewUserController(userService)

			{
				users.GET("/", uesrController.FindUsers)
				users.POST("/", uesrController.Registration)
				users.GET("/:id", uesrController.FindUser)
				users.PUT("/:id", uesrController.UpdatingUser)
				users.DELETE("/:id", uesrController.DeletedUser)
			}
			wallet := sakupay.Group("/wallet")
			walletRepo := repository.NewWalletRepository(config.DB)
			walletService := service.NewWalletService(walletRepo, userRepo)
			walletController := controller.NewWalletController(walletService)
			{
				wallet.POST("/", walletController.CreateHandler)
			}
		}
	}

	return router.Run()
}
