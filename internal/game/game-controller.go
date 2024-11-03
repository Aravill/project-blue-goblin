package game

import (
	Act "blue-goblin/internal/act"
	"blue-goblin/internal/command"
	"blue-goblin/internal/console"
	Player "blue-goblin/internal/player"
	"blue-goblin/internal/state"
	"blue-goblin/internal/world"
	"fmt"
	"os"
)

const defaultSaveName = "slot-1"

func Start() {
	// Set up the game control objects
	var player = Player.Player{}
	// This should load the particular act or the whole world, rn it's just a workaround
	var act Act.Act
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
			InitGameFromState(&player, &act, state)
		} else {
			InitNewGame(&player, &act)
		}
	} else {
		InitNewGame(&player, &act)
	}

	console.SayLine(act.GetLocation(player.CurrentLocation).GetFullDescription())
	for {
		// Get the player's input
		var input, error = console.Listen()
		if error != nil {
			fmt.Println("Error reading input: ", error)
			return
		}
		// Handle the command
		command.ExecuteCommand(&player, &act, input)
		// Check if the player is quitting
		if player.HasFlag(Player.Quitting) {
			Quit(player, act)
		}
	}
}

func InitGameFromState(player *Player.Player, act *Act.Act, state state.State) {
	player.CurrentLocation = state.Player.CurrentLocation
	player.Flags = state.Player.Flags
	player.Items = state.Player.Items
	act.Id = state.Act.Id
	act.Locations = state.Act.Locations
	console.SayLine("Game loaded successfully.")
}

func InitNewGame(player *Player.Player, act *Act.Act) {
	player.CurrentLocation = "act-one:start"
	player.Flags = []Player.Flag{}
	var loaded = world.LoadWorld()
	act.Id = loaded.Id
	act.Locations = world.LoadWorld().Locations
	console.SayLine("Starting a new game.")
}

func Quit(player Player.Player, act Act.Act) {
	player.RemoveFlag(Player.Quitting)
	state.SaveGame(defaultSaveName, player, act)
	os.Exit(0)
}
