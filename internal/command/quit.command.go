package command

import (
	"blue-goblin/internal/act"
	"blue-goblin/internal/console"
	"blue-goblin/internal/player"
)

func QuitCommand(p *player.Player, act *act.Act, params []string) {
	console.SayLine("Goodbye.")
	p.AddFlag(player.Quitting)
}
