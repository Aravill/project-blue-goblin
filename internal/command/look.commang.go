package command

import (
	"blue-goblin/internal/act"
	"blue-goblin/internal/console"
	"blue-goblin/internal/player"
)

func LookCommand(player *player.Player, act *act.Act, params []string) {
	var location = act.GetLocation(player.CurrentLocation)
	console.SayLine(location.GetFullDescription())
}
