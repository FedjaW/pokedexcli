package main

import (
	"fmt"
)

func commandPokedex(cfg *config, name string) error {
	for pokemonName := range cfg.Pokedex.Pokemons {
		fmt.Printf(" - %s\n", pokemonName)
	}

	return nil
}
