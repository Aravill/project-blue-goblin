package command

import (
	"blue-goblin/internal/act"
	"blue-goblin/internal/audio"
	"blue-goblin/internal/console"
	"blue-goblin/internal/player"
	"strings"
)

func UnlockCommand(aliasUsed string, player *player.Player, act *act.Act, params []string) {
	// unlock {exitName} with {keyName}
	var location = act.GetLocation(player.CurrentLocation)
	var exit = location.GetExit(params[0])
	if exit == nil {
		console.SayLine("There is no " + params[0] + " here.")
		return
	}
	if !exit.Lock.Status {
		console.SayLine("The " + params[0] + " is not locked.")
		return
	}
	if len(params) < 3 {
		console.SayLine("Unlock what with what?")
		return
	}
	var exitName = params[0]
	var keyName = strings.Join(params[2:], " ")
	var i = player.GetItem(keyName)
	if i == nil {
		console.SayLine("You do not have " + keyName + ".")
		return
	}
	var result = exit.Lock.Unlock(i)
	if !result {
		console.SayLine("The " + exitName + " cannot be unlocked with " + keyName + ".")
		return
	} else {
		player.RemoveItem(i.Id)
		go audio.PlaySound("unlock-1")
		console.SayLine("You unlock the " + exitName + " with " + keyName + ". The " + keyName + " has been removed from your bag.")
	}
}

func UnlockCommandAliases() []string {
	return []string{"unlock"}
}

func UnlockCommandHelp() string {
	return "[" + strings.Join(UnlockCommandAliases(), ", ") + "]" + " {object} with {key} - Unlock an object with a key."
}
