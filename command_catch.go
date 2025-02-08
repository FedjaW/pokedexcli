package main

import (
	"fmt"
	"math/rand"

	"github.com/fedjaw/pokedexcli/internal/api"
)

func commandCatch(cfg *config, name string) error {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
	pokemonResponse, err := api.GetPokemon(url, cfg.Cache)
	if err != nil {
		return err
	}

	fmt.Println(pokemonResponse.BaseExperience)
	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	random := rand.Intn(100)
	if random < 50 {
		fmt.Printf("%s escaped!\n", name)
	} else {
		fmt.Printf("%s was caught!\n", name)
		cfg.Pokedex.Pokemons[name] = pokemonResponse
	}

	return nil
}
