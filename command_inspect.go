package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, name string) error {
	if name == "" {
		return errors.New("you must provide a pokemon name")
	}
	pokemon, ok := cfg.Pokedex.Pokemons[name]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, v := range pokemon.Stats {
		fmt.Printf("  - %s%d\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Println("Types:")
	for _, v := range pokemon.Types {
		fmt.Printf("  - %s\n", v.Type.Name)
	}
	return nil
}
