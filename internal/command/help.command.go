package command

import (
	"blue-goblin/internal/act"
	"blue-goblin/internal/console"
	"blue-goblin/internal/player"
)

func HelpCommand(player *player.Player, act *act.Act, params []string) {
	var helpMessage = "You are in a text-based adventure game. You can type commands to interact with the world. Here are some commands you can use:\n"
	helpMessage += "> look - Look around.\n"
	helpMessage += "> go {direction} - Move to a different location.\n"
	helpMessage += "> take {item} - Pick up an item.\n"
	helpMessage += "> inventory - Show your inventory.\n"
	helpMessage += "> inspect {item} - Inspect an item in detail.\n"
	helpMessage += "> unlock {exitName} with {keyName} - Unlock an exit.\n"
	helpMessage += "> help - Show this help message.\n"
	helpMessage += "> quit - Exit the game."
	console.SayLine(helpMessage)
}
