package api

import (
	"github.com/gin-gonic/gin"
	"github.com/rwajon/ERC1155-Events/api/routes"
	"github.com/rwajon/ERC1155-Events/config"
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
