package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rwajon/erc1155-events/api/routes"
	"github.com/rwajon/erc1155-events/config"
	_ "github.com/rwajon/erc1155-events/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() {
	envs := config.GetEnvs()
	router := routes.Init()

	router.Use(gin.Recovery())
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if envs.Port == "" {
		router.Run()
	} else {
		router.Run(":" + envs.Port)
	}
}
