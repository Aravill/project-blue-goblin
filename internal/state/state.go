package state

import (
	"blue-goblin/internal/act"
	"blue-goblin/internal/player"
)

type State struct {
	Id     string        `json:"id"`
	Player player.Player `json:"player"`
	Act    act.Act       `json:"act"`
}

func NewState(saveName string, p player.Player, a act.Act) State {
	return State{
		Id:     saveName,
		Player: p,
		Act:    a,
	}
}
