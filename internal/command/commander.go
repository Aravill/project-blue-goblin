package command

import (
	"blue-goblin/internal/act"
	"blue-goblin/internal/console"
	"blue-goblin/internal/player"
	"strings"
)

type CommandRecord struct {
	Aliases []string
	Handler func(string, *player.Player, *act.Act, []string)
	Doc     string
}

var commandDictionary = []CommandRecord{
	{Aliases: GoCommandAliases(), Handler: GoCommand, Doc: GoCommandHelp()},
	{Aliases: QuitCommandAliases(), Handler: QuitCommand, Doc: QuitCommandHelp()},
	{Aliases: TakeCommandAliases(), Handler: TakeCommand, Doc: TakeCommandHelp()},
	{Aliases: InventoryCommandAliases(), Handler: InventoryCommand, Doc: InventoryCommandHelp()},
	{Aliases: InspectCommandAliases(), Handler: InspectCommand, Doc: InspectCommandHelp()},
	{Aliases: LookCommandAliases(), Handler: LookCommand, Doc: LookCommandHelp()},
	{Aliases: UnlockCommandAliases(), Handler: UnlockCommand, Doc: UnlockCommandHelp()},
	{Aliases: HelpCommandAliases(), Handler: nil, Doc: HelpCommandHelp()},
}

var wordToCommand = make(map[string]CommandRecord)
var initialized = false

func PopulateMap() {
	if initialized {
		return
	}
	for _, record := range commandDictionary {
		for _, alias := range record.Aliases {
			wordToCommand[alias] = record
		}
	}
	initialized = true
}

func ExecuteCommand(player *player.Player, act *act.Act, input string) {
	PopulateMap()
	var inputParts = strings.Split(input, " ")
	var commandName = inputParts[0]
	var params = inputParts[1:]
	var handler = wordToCommand[commandName].Handler
	if handler == nil {
		// Workaround for the help command
		for _, alias := range HelpCommandAliases() {
			if commandName == alias {
				HelpCommand(commandName, player, act, params)
				return
			}
		}
		console.SayLine("You don't know how to do that.")
		return
	}
	handler(commandName, player, act, params)

}

// This is here to avoid circular dependencies
func HelpCommand(aliasUsed string, player *player.Player, act *act.Act, params []string) {
	var helpText = "You can use the following commands:\n"
	for _, command := range commandDictionary {
		helpText += command.Doc + "\n"
	}
	console.Say(helpText)
}
