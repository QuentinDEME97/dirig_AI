package state

import (
	"agents/world"
	"agents/worker"
)

type State struct {
	Rulers []worker.Ruler
	Map world.Map
}