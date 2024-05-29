package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		commands := getCommands()
		command, ok := commands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		command.callback()
	}
}

type commandCLI struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]commandCLI {
	return map[string]commandCLI{
		"help": {
			name:        "help",
			description: "Displays a help menu",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Turns off the pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(str string) []string {
	output := strings.ToLower(str)
	words := strings.Fields(output)
	return words
}
