package command

import (
	"blue-goblin/internal/act"
	"blue-goblin/internal/console"
	"blue-goblin/internal/player"
	"strings"
)

func GoCommand(aliasUsed string, player *player.Player, act *act.Act, params []string) {
	var location = act.GetLocation(player.CurrentLocation)
	var target = params[0]
	if len(target) == 0 {
		console.SayLine(aliasUsed + " where?")
		return
	}
	var exit = location.GetExit(target)
	if exit == nil {
		console.SayLine("You cannot go that way.")
		return
	}
	if exit.Lock.Status {
		console.SayLine(exit.Lock.Description)
		return
	}
	var destination = act.GetLocation(exit.Destination)
	console.SayLine("You go " + target + ".")
	player.MoveTo(destination)
	console.SayLine(destination.GetFullDescription())
}

func GoCommandAliases() []string {
	return []string{"go", "head", "walk", "run", "move"}
}

func GoCommandHelp() string {
	return "[" + strings.Join(GoCommandAliases(), ", ") + "]" + " {direction} - Move in a direction."
}
