package main

import (
	"net/http"

	"github.com/chuckpreslar/emission"
	"github.com/gin-gonic/gin"
	"github.com/rwajon/erc1155-events/api"
	"github.com/rwajon/erc1155-events/config"
	"github.com/rwajon/erc1155-events/db"
	"github.com/rwajon/erc1155-events/models"
	"github.com/rwajon/erc1155-events/services"
)

// @title ERC1155-events
// @version 1.0
// @description ERC1155 events.
// @termsOfService http://swagger.io/terms/

// @contact.name Jonathan Rwabahizi
// @contact.url http://www.swagger.io/support
// @contact.email jonathanrwabahizi@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api/v1
// @schemes http
// @schemes https
func main() {
	app := &models.App{
		Envs:         config.GetEnvs(),
		EventEmitter: emission.NewEmitter(),
	}

	if db.Init() == nil {
		r := gin.Default()
		r.Any("/*any", func(c *gin.Context) {
			c.String(http.StatusInternalServerError, "can not connect to database")
		})
		r.Run(":" + app.Envs.Port)
	}
	go func() {
		services.ListenToERC1155Events(*app)
	}()
	api.Run(app)
}
