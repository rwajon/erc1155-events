package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rwajon/erc1155-events/api/controllers"
)

func Init() *gin.Engine {
	router := gin.Default()
	apiV1 := router.Group("/api").Group("/v1")

	router.GET("/ping", controllers.Ping)
	apiV1.GET("/ping", controllers.Ping)
	transactionRoutes(apiV1.Group("/transactions"))
	watchListRoutes(apiV1.Group("/watch-list"))

	return router
}
