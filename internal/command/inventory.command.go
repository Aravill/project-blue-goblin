package command

import (
	"blue-goblin/internal/act"
	"blue-goblin/internal/console"
	"blue-goblin/internal/player"
)

func InventoryCommand(player *player.Player, act *act.Act, params []string) {
	console.SayLine("You look into your backpack.")
	if len(player.Items) == 0 {
		console.SayLine("It is empty.")
		return
	}
	var fullDescription = "You have: "
	for _, i := range player.Items {
		if i == player.Items[len(player.Items)-1] {
			fullDescription += i.Name + "."
		} else {
			fullDescription += i.Name + ", "
		}
	}
	console.SayLine(fullDescription)
}
