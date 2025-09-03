package main

import (
	"fmt"
	"os"
)

func commandExit(_ []string) error {
	fmt.Println("Closing the Database... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ []string) error {
	fmt.Println("Welcome to the Music Libary!")
	fmt.Printf("Usage:\n\n")
	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
