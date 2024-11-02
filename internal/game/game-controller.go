package game

import (
	"blue-goblin/internal/command"
	"blue-goblin/internal/console"
	"blue-goblin/internal/player"
	"blue-goblin/internal/world"
	"fmt"
)

func Start() {
	// Set up the game control objects
	var player = player.Player{
		CurrentLocation: "act-one:start",
	}
	var actOne = world.LoadWorld()
	// Initial narration
	console.SayLine("Welcome to Blue Goblin!")
	console.SayLine(actOne.GetLocation(player.CurrentLocation).GetFullDescription())
	for {
		// Get the player's input
		var input, error = console.Listen()
		if error != nil {
			fmt.Println("Error reading input: ", error)
			return
		}
		// Handle the command
		command.ExecuteCommand(&player, &actOne, input)
	}
}
