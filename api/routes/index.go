package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rwajon/ERC1155-Events/api/controllers"
)

func Init() *gin.Engine {
	router := gin.Default()
	apiV1 := router.Group("/v1").Group("/api")

	router.GET("/v1/ping", controllers.Ping)
	apiV1.GET("/ping", controllers.Ping)

	return router
}
