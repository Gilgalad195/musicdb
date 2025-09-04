package main

import "database/sql"

type cliCommand struct {
	name        string
	description string
	callback    func(*sql.DB) error
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
		"add": {
			name:        "add",
			description: "adds a song to the database",
			callback:    commandAdd,
		},
		"list": {
			name:        "list",
			description: "lists the songs in the database",
			callback:    commandList,
		},
	}
}
