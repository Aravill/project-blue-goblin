package command

func HelpCommandAliases() []string {
	return []string{"help"}
}

func HelpCommandHelp() string {
	return "[" + HelpCommandAliases()[0] + "] - Show help."
}
