package game

import (
	"blue-goblin/internal/command"
	"blue-goblin/internal/console"
	"blue-goblin/internal/player"
	"blue-goblin/internal/state"
	"blue-goblin/internal/world"
	"fmt"
	"os"
)

const defaultSaveName = "slot-1"

func Start() {
	// Set up the game control objects
	var p = player.Player{}
	// This should load the particular act or the whole world, rn it's just a workaround
	var act = world.LoadWorld()
	// Initial narration
	console.SayLine("Welcome to Blue Goblin!")

	if state.SaveExists(defaultSaveName) {
		console.SayLine("Would you like to continue your saved game [y/n]?")
		var input, error = console.Listen()
		if error != nil {
			fmt.Println("Error reading input: ", error)
			return
		}
		if len(input) > 0 && (input[0] == 'y' || input[0] == 'Y') {
			state, err := state.LoadGame(defaultSaveName)
			if err != nil {
				fmt.Println("Failed to load game:", err)
				return
			}
			InitGameFromState(&p, state)
		} else {
			InitNewGame(&p)
		}
	} else {
		InitNewGame(&p)
	}

	console.SayLine(act.GetLocation(p.CurrentLocation).GetFullDescription())
	for {
		// Get the player's input
		var input, error = console.Listen()
		if error != nil {
			fmt.Println("Error reading input: ", error)
			return
		}
		// Handle the command
		command.ExecuteCommand(&p, &act, input)
		// Check if the player is quitting
		if p.HasFlag(player.Quitting) {
			Quit(p)
		}
	}
}

func InitGameFromState(p *player.Player, state state.State) {
	p.CurrentLocation = state.Player.CurrentLocation
	p.Flags = state.Player.Flags
	console.SayLine("Game loaded successfully.")
}

func InitNewGame(p *player.Player) {
	p.CurrentLocation = "act-one:start"
	p.Flags = []player.Flag{}
	console.SayLine("Starting a new game.")
}

func Quit(p player.Player) {
	p.RemoveFlag(player.Quitting)
	state.SaveGame(defaultSaveName, p)
	os.Exit(0)
}
