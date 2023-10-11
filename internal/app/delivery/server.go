package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/sakupay-apps/internal/app/delivery/routes"
)

type application struct {
	engine *gin.Engine
}

func (app *application) Run() {
	if err := routes.SetupRouter(app.engine); err != nil {
		panic("Aplication error")
	}
}

func Server() *application {
	router := gin.Default()

	return &application{
		engine: router,
	}

}
