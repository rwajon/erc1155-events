package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rwajon/erc1155-events/api/controllers"
)

func watchListRoutes(router *gin.RouterGroup) *gin.RouterGroup {
	router.POST("/", controllers.AddAddressInWatchList)
	router.GET("/", controllers.GetWatchList)
	router.GET("/:address", controllers.GetOneAddressWatchList)
	router.PUT("/:addressId", controllers.UpdateAddressInWatchList)
	router.DELETE("/:addressId", controllers.DeleteAddressInWatchList)
	return router
}
