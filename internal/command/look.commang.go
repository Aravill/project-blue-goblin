package command

import (
	"blue-goblin/internal/act"
	"blue-goblin/internal/console"
	"blue-goblin/internal/player"
	"strings"
)

func LookCommand(aliasUsed string, player *player.Player, act *act.Act, params []string) {
	var location = act.GetLocation(player.CurrentLocation)
	console.SayLine(location.GetFullDescription())
}

func LookCommandAliases() []string {
	return []string{"look"}
}

func LookCommandHelp() string {
	return "[" + strings.Join(LookCommandAliases(), ", ") + "] - Look around."
}
