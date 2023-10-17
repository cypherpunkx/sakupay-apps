package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sakupay-apps/config"
	"github.com/sakupay-apps/internal/app/delivery/controller"
	"github.com/sakupay-apps/internal/app/delivery/middleware"
	"github.com/sakupay-apps/internal/app/manager"
)

func SetupRouter(router *gin.Engine) error {

	// router.Use(middleware.LogRequestMiddleware(logrus.New()))

	infraManager := manager.NewInfraManager(config.Cfg)
	serviceManager := manager.NewRepoManager(infraManager)
	repoManager := manager.NewServiceManager(serviceManager)

	// User Controller
	userController := controller.NewUserController(repoManager.UserService(), repoManager.AuthService(), repoManager.UserPictureService())
	// Transaction Controller
	transactionController := controller.NewTransactionController(repoManager.TransactionService())
	// Bill Controller
	billController := controller.NewBillController(repoManager.BillService())
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
				users.GET("/", userController.FindAllUsers)
				users.GET("/:id", userController.FindUser)
				users.PUT("/:id", userController.UpdateUser)
				users.POST("/:id/upload", userController.UploadPicture)
				users.GET("/:id/download", userController.DownloadPicture)
				users.DELETE("/:id", userController.DeleteUser)
				// Contact
				users.POST("/:id/contacts", contactController.AddContact)
				users.GET("/:id/contacts", contactController.FindAllContacts)
				users.GET("/:id/contacts/:contactID", contactController.FindContact)
				users.DELETE("/:id/contacts/:contactID", contactController.DeleteContact)
				// Transaction
				users.POST("/:id/transactions", transactionController.CreateTransaction)
				users.GET("/:id/transactions", transactionController.FindAllTransactions)
				users.GET("/:id/transactions/:transactionID", transactionController.FindTransaction)
				// Bill
				users.POST("/:id/bills", billController.CreateBill)
				users.GET("/:id/bills", billController.FindAllBills)
				// Card
				users.GET("/:id/cards", cardController.FindAllCards)
				users.POST("/:id/cards", cardController.AddCard)
				users.DELETE("/:id/cards/:cardID", cardController.DeleteCard)
			}

			// go func() {
			// 	for {
			// 		time.Sleep(10 * time.Second) // Cek setiap 24 jam

			// 		var bills []model.Bill
			// 		config.DB.Where("due_date BETWEEN ? AND ?", time.Now(), time.Now().Add(5*time.Minute)).Find(&bills)

			// 		for _, bill := range bills {
			// 			if !bill.Notified {
			// 				fmt.Printf("Mengirim pemberitahuan ke %s untuk tagihan %d\n", bill.DueDate, bill.Status)
			// 				// Di sini, Anda dapat mengirim pemberitahuan melalui email, SMS, atau media komunikasi lainnya.
			// 				// Atur `bill.Notified = true` setelah pemberitahuan dikirim agar pemberitahuan hanya dikirim sekali.
			// 			}
			// 		}
			// 		fmt.Println("OEKEEKEKEKKEKE")
			// 	}
			// }()
		}
	}

	return router.Run()

}
