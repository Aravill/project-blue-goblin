package command

import (
	"blue-goblin/internal/act"
	"blue-goblin/internal/audio"
	"blue-goblin/internal/console"
	"blue-goblin/internal/player"
	"strings"
)

func TakeCommand(player *player.Player, act *act.Act, params []string) {
	var location = act.GetLocation(player.CurrentLocation)
	var itemName = strings.Join(params, " ")
	var i = location.RemoveItem(itemName)
	if i == nil {
		console.SayLine("There is no " + itemName + " here.")
		return
	}
	go audio.PlaySound("pick-up-1")
	console.SayLine("You take the " + itemName + ".")
	player.AddItem(i)
}
