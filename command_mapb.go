package main

import (
	"fmt"

	"github.com/fedjaw/pokedexcli/internal/api"
)

func commandMapb(cfg *config, name string) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locationAreaResponse, err := api.GetMap(*cfg.Previous, cfg.Cache)
	if err != nil {
		return err
	}

	cfg.Next = locationAreaResponse.Next
	cfg.Previous = locationAreaResponse.Previous

	for _, locationArea := range locationAreaResponse.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}
