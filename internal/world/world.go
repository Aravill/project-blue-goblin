package world

import (
	"blue-goblin/internal/act"
)

func LoadWorld() act.Act {
	// Returning just the first act for now
	return act.LoadActOne()
}
