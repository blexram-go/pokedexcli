package main

import (
	"errors"
	"fmt"
	"os"
)

func commandHelp(cfg *config) error {
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

func commandExit(cfg *config) error {
	os.Exit(0)
	return nil
}

func commandMapF(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreasURL)
	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("-%s\n", area.Name)
	}

	cfg.nextLocationAreasURL = resp.Next
	cfg.prevLocationAreasURL = resp.Previous

	return nil
}

func commandMapB(cfg *config) error {
	if cfg.prevLocationAreasURL == nil {
		return errors.New("this is the first page")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreasURL)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("-%s\n", area.Name)
	}

	cfg.nextLocationAreasURL = resp.Next
	cfg.prevLocationAreasURL = resp.Previous

	return nil
}
