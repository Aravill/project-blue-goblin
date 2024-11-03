package command

import (
	"blue-goblin/internal/act"
	"blue-goblin/internal/console"
	"blue-goblin/internal/player"
)

func TakeCommand(player *player.Player, act *act.Act, params []string) {
	var location = act.GetLocation(player.CurrentLocation)
	var itemName = params[0]
	var i = location.RemoveItem(itemName)
	if i == nil {
		console.SayLine("There is no " + itemName + " here.")
		return
	}
	console.SayLine("You take the " + itemName + ".")
	player.AddItem(i)
}
