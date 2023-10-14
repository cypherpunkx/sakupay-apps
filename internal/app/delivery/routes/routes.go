package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sakupay-apps/config"
	"github.com/sakupay-apps/internal/app/delivery/controller"
	"github.com/sakupay-apps/internal/app/repository"
	"github.com/sakupay-apps/service"
)

func SetupRouter(router *gin.Engine) error {

	
	v1 := router.Group("/api/v1")
	{
		sakupay := v1.Group("/sakupay")
		{

			bills := sakupay.Group("/bill")
			repository := repository.NewBillRepository(config.DB)
			service := service.NewBillService(repository)
			controller := controller.NewBillController(service)
			{		
			bills.POST("/", controller.CreateNewBill) 
			bills.GET("/:id", controller.GetDetailBill) 
			}

		}
	}

	return router.Run()
}
