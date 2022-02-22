package models

import (
	"github.com/chuckpreslar/emission"
	"github.com/rwajon/erc1155-events/config"
)

type App struct {
	Envs         config.Env
	EventEmitter *emission.Emitter
}
