package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gobash-blex/pokedexcli/internal/pokeapi"
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

func commandMap() error {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("-%s\n", area.Name)
	}
	return nil
}
