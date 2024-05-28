package commands

import (
	"errors"
	"fmt"
	"os"

	"github.com/fedjaw/pokedexcli/internals/api"
)

type Config struct {
	Next *string
	Prev *string
}

func NewConfig() Config {
	return Config{Next: nil, Prev: nil}
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

func GetCliCommandsMap() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			Name:        "map",
			Description: "Displays the names of the next 20 location areas in the Pokemon world",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "map",
			Description: "Displays the names of the previous 20 location areas in the Pokemon world",
			Callback:    commandMapb,
		},
	}
}

func commandHelp(cfg *Config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	fmt.Println()
	commands := GetCliCommandsMap()
	for k, v := range commands {
		fmt.Print(k)
		fmt.Print(": ")
		fmt.Println(v.Description)
	}
	fmt.Println()
	return nil
}

func commandExit(cfg *Config) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *Config) error {
	if cfg.Next == nil {
		startUrl := "https://pokeapi.co/api/v2/location-area"
		cfg.Next = &startUrl
	}

	locationAreas, err := api.GetMap(*cfg.Next)
	cfg.Next = locationAreas.Next
	cfg.Prev = locationAreas.Previous
	for _, locationArea := range locationAreas.Results {
		fmt.Println(locationArea.Name)
	}
	return err
}

func commandMapb(cfg *Config) error {
	if cfg.Prev == nil {
		return errors.New("no prev location")
	}

	locationAreas, err := api.GetMap(*cfg.Prev)
	cfg.Next = locationAreas.Next
	cfg.Prev = locationAreas.Previous
	for _, locationArea := range locationAreas.Results {
		fmt.Println(locationArea.Name)
	}
	return err
}
