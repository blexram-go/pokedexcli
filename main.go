package main

import (
	"time"

	"github.com/gobash-blex/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationAreasURL *string
	prevLocationAreasURL *string
	caughtPokemon        map[string]pokeapi.PokemonData
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.PokemonData),
	}
	startRepl(&cfg)
}
