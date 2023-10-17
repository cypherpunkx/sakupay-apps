package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sakupay-apps/config"
	"github.com/sakupay-apps/internal/app/delivery/controller"
	"github.com/sakupay-apps/internal/app/delivery/middleware"
	"github.com/sakupay-apps/internal/app/manager"
)

func SetupRouter(router *gin.Engine) error {

	infraManager := manager.NewInfraManager(config.Cfg)
	serviceManager := manager.NewRepoManager(infraManager)
	repoManager := manager.NewServiceManager(serviceManager)

	// User Controller
	userController := controller.NewUserController(repoManager.UserService(), repoManager.AuthService())
	// Transaction Controller
	transactionController := controller.NewTransactionController(repoManager.TransactionService())
	// Bill Controller
	billController := controller.NewBillController(repoManager.BillService(), repoManager.UserService())
	// Contact Controller
	contactController := controller.NewContactController(repoManager.ContactService())
        // Card Controller
	cardController := controller.NewCardController(repoManager.CardService())

	v1 := router.Group("/api/v1")
	{
		sakupay := v1.Group("/sakupay")
		{
			auth := sakupay.Group("/auth")
			{
				auth.POST("/register", userController.Registration)
				auth.POST("/login", userController.Login)
			}

			users := sakupay.Group("/users", middleware.AuthMiddleware())
			{
				users.GET("/", userController.FindUsers)
				users.GET("/:id", userController.FindUser)
				users.PUT("/:id", userController.UpdatingUser)
				users.DELETE("/:id", userController.DeletedUser)
				users.POST("/:id/transaction", transactionController.CreateDeposit)
				users.POST("/:id/bill", billController.CreateNewBill)
				users.POST("/:id/contact", contactController.AddContact)
				// users.GET("/:id/contact", contactController.ListHandler)
				// users.GET("/:user_id/contact/:contact_id", contactController.GetHandler)
				// users.PUT("/:id/contact", contactController.UpdateHandler)
				// users.DELETE("/:id/contact", contactController.DeleteHandler)
				users.POST("/:id/cards", cardController.AddCard)
				//users.GET("/", cardController.)
			}
		}
	}
	return router.Run()

}
