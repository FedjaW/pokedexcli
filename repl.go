package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fedjaw/pokedexcli/internal/api"
	"github.com/fedjaw/pokedexcli/internal/pokecache"
)

type config struct {
	Next     *string
	Previous *string
	Cache    *pokecache.Cache
	Pokedex  *pokedex
}

type pokedex struct {
	Pokemons map[string]api.PokemonResponse
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Displays the names of the pokemons for the given location areas in the Pokemon world",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Inspect your pokedex",
			callback:    commandPokedex,
		},
	}
}

func startRepl() {
	const PROMT = "Pokedex > "

	cache := pokecache.NewCache(5 * time.Second)
	cfg := config{
		Next:     nil,
		Previous: nil,
		Cache:    cache,
		Pokedex: &pokedex{
			Pokemons: make(map[string]api.PokemonResponse),
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(PROMT)
		scanner.Scan()
		text := scanner.Text()
		cleanedTextArray := cleanInput(text)
		command := cleanedTextArray[0]

		parameter := ""
		if len(cleanedTextArray) > 1 {
			parameter = cleanedTextArray[1]
		}

		commands := getCommands()
		c, ok := commands[command]

		if !ok {
			fmt.Println("Unknown command")
		} else {
			err := c.callback(&cfg, parameter)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
