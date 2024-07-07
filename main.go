package main

import (
	"log"
	"time"

	"github.com/gobash-blex/pokedexcli/internal/database"
	"github.com/gobash-blex/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient        pokeapi.Client
	nextLocationAreasURL *string
	prevLocationAreasURL *string
	caughtPokemon        map[string]pokeapi.PokemonData
	DB                   *database.DB
}

func main() {
	db, err := database.NewDB("database.json")
	if err != nil {
		log.Fatal(err)
	}

	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.PokemonData),
		DB:            db,
	}
	startRepl(&cfg)
}
