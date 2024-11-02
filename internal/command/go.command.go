package command

import (
	"blue-goblin/internal/act"
	"blue-goblin/internal/console"
	"blue-goblin/internal/player"
)

func GoCommand(player *player.Player, act *act.Act, params []string) {
	var location = act.GetLocation(player.CurrentLocation)
	var target = params[0]
	var exit = location.GetExit(target)
	if exit == nil {
		console.SayLine("You cannot go that way.")
		return
	}
	var destination = act.GetLocation(exit.Destination)
	console.SayLine("You go " + target + ".")
	player.MoveTo(destination)
	console.SayLine(destination.GetFullDescription())
}
