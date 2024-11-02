package command

import (
	"blue-goblin/internal/act"
	"blue-goblin/internal/console"
	"blue-goblin/internal/player"
	"strings"
)

func ExecuteCommand(player *player.Player, act *act.Act, input string) {
	var inputParts = strings.Split(input, " ")
	var commandName = inputParts[0]
	var params = inputParts[1:]
	switch commandName {
	case "go":
		GoCommand(player, act, params)
	default:
		console.SayLine("I don't know how to do that.")
	}
}
