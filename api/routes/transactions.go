package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rwajon/erc1155-events/api/controllers"
)

func transactionRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	router.GET("/", controllers.GetTransactions)
	return router
}
