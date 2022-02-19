package main

import (
	"github.com/rwajon/erc1155-events/api"
	"github.com/rwajon/erc1155-events/db"
	"github.com/rwajon/erc1155-events/services"
)

// @title Swagger Example API
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
func main() {
	go db.Init()
	go services.ListenToERC1155Events()
	api.Run()
}
