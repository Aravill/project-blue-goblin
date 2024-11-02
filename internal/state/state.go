package state

import "blue-goblin/internal/player"

type State struct {
	Id     string
	Player player.Player
}

func NewState(saveName string, player player.Player) State {
	return State{
		Id:     saveName,
		Player: player,
	}
}
