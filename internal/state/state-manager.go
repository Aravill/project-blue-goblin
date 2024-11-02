package state

import (
	"blue-goblin/internal/player"
	"encoding/json"
	"os"
)

const saveDir = "saves"

func SaveGame(saveName string, player player.Player) error {
	var state = NewState(saveName, player)
	data, err := json.MarshalIndent(state, "", "  ")
	if err != nil {
		return err
	}
	if err := os.MkdirAll(saveDir, 0755); err != nil {
		return err
	}
	return os.WriteFile(BuildFileName(saveName), data, 0644)
}

func SaveExists(saveName string) bool {
	_, err := os.Stat(BuildFileName(saveName))
	return err == nil
}

func LoadGame(saveName string) (State, error) {
	var state State
	data, err := os.ReadFile(BuildFileName(saveName))
	if err != nil {
		return state, err
	}
	err = json.Unmarshal(data, &state)
	return state, err
}

func BuildFileName(saveName string) string {
	return saveDir + "/" + saveName + ".save.json"
}
