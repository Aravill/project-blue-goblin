package command

import (
	"blue-goblin/internal/act"
	"blue-goblin/internal/console"
	"blue-goblin/internal/player"
	"strings"
)

func QuitCommand(aliasUsed string, p *player.Player, act *act.Act, params []string) {
	console.SayLine("Goodbye.")
	p.AddFlag(player.Quitting)
}

func QuitCommandAliases() []string {
	return []string{"quit", "exit"}
}

func QuitCommandHelp() string {
	return "[" + strings.Join(QuitCommandAliases(), ", ") + "] - Exit the game."
}
