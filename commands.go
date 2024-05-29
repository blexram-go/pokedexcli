package main

import (
	"fmt"
	"os"
)

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedexcli!")
	fmt.Println("Here are your available commands:")

	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
