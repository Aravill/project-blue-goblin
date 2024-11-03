package state

import (
	"blue-goblin/internal/act"
	"blue-goblin/internal/player"
)

type State struct {
	Id     string
	Player player.Player
	Act    act.Act
}

func NewState(saveName string, p player.Player, a act.Act) State {
	return State{
		Id:     saveName,
		Player: p,
		Act:    a,
	}
}
