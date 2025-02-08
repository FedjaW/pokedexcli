package main

import (
	"fmt"

	"github.com/fedjaw/pokedexcli/internal/api"
)

func commandExplore(cfg *config, name string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", name)
	locationAreaDetailsResponse, err := api.GetMapDetails(url, cfg.Cache)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationAreaDetailsResponse.PokemonEncounters {
		fmt.Println("- " + pokemon.Pokemon.Name)
	}

	return nil
}
