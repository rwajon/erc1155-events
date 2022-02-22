package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rwajon/erc1155-events/api/controllers"
	"github.com/rwajon/erc1155-events/models"
)

func Init(app *models.App) *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Set("app", app)
		c.Next()
	})

	apiV1 := router.Group("/api").Group("/v1")

	router.GET("/ping", controllers.Ping)
	apiV1.GET("/ping", controllers.Ping)
	transactionRoutes(apiV1.Group("/transactions"))
	watchListRoutes(apiV1.Group("/watch-list"))

	return router
}
