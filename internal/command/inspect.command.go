package command

import (
	"blue-goblin/internal/act"
	"blue-goblin/internal/console"
	"blue-goblin/internal/player"
)

func InspectCommand(player *player.Player, act *act.Act, params []string) {
	var location = act.GetLocation(player.CurrentLocation)
	var itemName = params[0]
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
