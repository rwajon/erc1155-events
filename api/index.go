package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rwajon/erc1155-events/api/routes"
	_ "github.com/rwajon/erc1155-events/docs"
	"github.com/rwajon/erc1155-events/models"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run(app *models.App) {
	router := routes.Init(app)

	router.Use(gin.Recovery())
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to ERC1155-events!")
	})
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + app.Envs.Port)
}
