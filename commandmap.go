package main

type cliCommand struct {
	name        string
	description string
	callback    func([]string) error
	//this is so the commands can receive and update the previous/next state.
}

var commands map[string]cliCommand

// definiing and using init() is to prevent circular dependencies between here and main.
func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Database",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}
