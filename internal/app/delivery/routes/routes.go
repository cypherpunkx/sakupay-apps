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
			contacts := sakupay.Group("/contacts")
			
				repository := repository.NewContactRepository(config.DB)
				service := service.NewContactService(repository)
				controller := controller.NewContactController(service)

			{

				contacts.POST("/", controller.CreateHandler)
				contacts.GET("/", controller.ListHandler)
				contacts.GET("/:id", controller.GetHandler)
				contacts.POST("/:id", controller.UpdateHandler)
				contacts.DELETE("/:id", controller.DeleteHandler)
			}
		}

	}


	return router.Run()
}
