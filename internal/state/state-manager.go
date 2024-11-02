package state

import (
	"blue-goblin/internal/player"
	"encoding/json"
	"os"
)

func SaveGame(saveName string, player player.Player) error {
	var state = NewState(saveName, player)
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll("saves", 0755); err != nil {
		return err
	}
	return os.WriteFile("saves/"+saveName, data, 0644)
}

func LoadGame(saveName string) (State, error) {
	var state State
	data, err := os.ReadFile("saves/" + saveName)
	if err != nil {
		return state, err
	}
	err = json.Unmarshal(data, &state)
	return state, err
}
