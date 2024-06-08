package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (PokemonData, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	dat, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Println("cache hit!")
		pokemonData := PokemonData{}
		err := json.Unmarshal(dat, &pokemonData)
		if err != nil {
			return PokemonData{}, err
		}
		return pokemonData, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonData{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return PokemonData{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return PokemonData{}, err
	}

	pokemonData := PokemonData{}
	err = json.Unmarshal(dat, &pokemonData)
	if err != nil {
		return PokemonData{}, err
	}

	c.cache.Add(fullURL, dat)
	return pokemonData, nil
}
