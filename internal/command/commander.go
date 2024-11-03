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
	case "quit":
		QuitCommand(player, act, params)
	case "take":
		TakeCommand(player, act, params)
	case "inventory":
		InventoryCommand(player, act, params)
	case "inspect":
		InspectCommand(player, act, params)
	case "look":
		LookCommand(player, act, params)
	default:
		console.SayLine("You don't know how to do that.")
	}
}
