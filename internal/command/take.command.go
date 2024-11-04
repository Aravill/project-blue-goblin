package command

import (
	"blue-goblin/internal/act"
	"blue-goblin/internal/audio"
	"blue-goblin/internal/console"
	"blue-goblin/internal/player"
	"strings"
)

func TakeCommand(aliasUsed string, player *player.Player, act *act.Act, params []string) {
	var location = act.GetLocation(player.CurrentLocation)
	var itemName = strings.Join(params, " ")
	if len(itemName) == 0 {
		console.SayLine(aliasUsed + " what?")
		return
	}
	var i = location.RemoveItem(itemName)
	if i == nil {
		console.SayLine("There is no " + itemName + " here.")
		return
	}
	go audio.PlaySound("pick-up-1")
	console.SayLine("You take the " + itemName + ".")
	player.AddItem(i)
}

func TakeCommandAliases() []string {
	return []string{"take", "get", "pick"}
}

func TakeCommandHelp() string {
	return "[" + strings.Join(TakeCommandAliases(), ", ") + "]" + " {item} - Take an item."
}
