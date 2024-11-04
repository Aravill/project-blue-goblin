package command

import (
	"blue-goblin/internal/act"
	"blue-goblin/internal/console"
	"blue-goblin/internal/player"
	"strings"
)

func InspectCommand(aliasUsed string, player *player.Player, act *act.Act, params []string) {
	var location = act.GetLocation(player.CurrentLocation)
	var itemName = strings.Join(params, " ")
	if len(itemName) == 0 {
		console.SayLine(aliasUsed + " what?")
		return
	}
	var i = location.GetItem(itemName)
	if i == nil {
		i = player.GetItem(itemName)
		if i == nil {
			console.SayLine("There is no " + itemName + " here.")
			return
		}
	}
	console.SayLine(i.Description)
}

func InspectCommandAliases() []string {
	return []string{"inspect", "examine"}
}

func InspectCommandHelp() string {
	return "[" + strings.Join(InspectCommandAliases(), ", ") + "]" + " {item} - Examine an item."
}
