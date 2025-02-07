package main

import (
	"fmt"

	"github.com/fedjaw/pokedexcli/internal/api"
)

func commandMap(cfg *config) error {
    if cfg.Next == nil {
        startUrl := "https://pokeapi.co/api/v2/location-area"
        cfg.Next = &startUrl
    }

    locationAreaResponse, err := api.GetMap(*cfg.Next, cfg.Cache)
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
