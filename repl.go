package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fedjaw/pokedexcli/internal/pokecache"
)

type config struct {
    Next *string
    Previous *string
    Cache *pokecache.Cache
}

type cliCommand struct {
    name string
    description string
    callback func(*config) error
}

func getCommands() map[string]cliCommand {
    return map[string]cliCommand{
        "exit": {
            name: "exit",
            description: "Exit the Pokedex",
            callback: commandExit,
        },
        "help": {
            name: "help",
            description: "Displays a help message",
            callback: commandHelp,
        },
        "map": {
            name: "map",
			description: "Displays the names of the next 20 location areas in the Pokemon world",
            callback: commandMap,
        },
        "mapb": {
            name: "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
            callback: commandMapb,
        },
    }
}

func startRepl() {
    const PROMT = "Pokedex > " 

    cache := pokecache.NewCache(5 * time.Second)
    cfg := config{
        Next: nil,
        Previous: nil,
        Cache: cache,
    }

    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print(PROMT)
        scanner.Scan()
        text := scanner.Text()
        cleanedTextArray := cleanInput(text)
        command := cleanedTextArray[0]

        commands := getCommands()

        c, ok := commands[command]

        if !ok {
            fmt.Println("Unknown command")
        } else {
            err := c.callback(&cfg)
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
