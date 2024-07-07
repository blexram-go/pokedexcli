package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func commandHelp(cfg *config, args ...string) error {
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

func commandExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}

func commandMapF(cfg *config, args ...string) error {
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

func commandMapB(cfg *config, args ...string) error {
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

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no location area provided")
	}

	locationAreaName := args[0]

	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}
	fmt.Printf("Pokemon in %s:\n", locationArea.Name)
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" -%s\n", pokemon.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.RequestPokemonByName(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemonName)
	}
	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s was caught!\n", pokemonName)
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.caughtPokemon[pokemonName]
	if !ok {
		return errors.New("you haven't caught this pokemon yet")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Moves:")
	for _, move := range pokemon.Moves {
		fmt.Printf(" - %s", move.Move.Name)
	}
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Pokemon type:")
	for _, typ := range pokemon.Types {
		fmt.Printf(" - %s\n", typ.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Pokemon in Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemon.Name)
	}
	return nil
}
