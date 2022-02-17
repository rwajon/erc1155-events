package main

import (
	"github.com/rwajon/erc1155-events/api"
	"github.com/rwajon/erc1155-events/db"
	"github.com/rwajon/erc1155-events/services"
)

func main() {
	go db.Init()
	go services.ListenToERC1155Events()
	api.Run()
}
