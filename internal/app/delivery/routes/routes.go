package routes

import (
	"github.com/gin-gonic/gin"
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
		}
	}

	return router.Run()
}
