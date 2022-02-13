package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rwajon/erc1155-events/api/routes"
	"github.com/rwajon/erc1155-events/config"
)

func Run() {
	envs := config.GetEnvs()
	router := routes.Init()

	router.Use(gin.Recovery())

	if envs.Port == "" {
		router.Run()
	} else {
		router.Run(":" + envs.Port)
	}
}
