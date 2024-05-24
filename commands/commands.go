package commands

import (
	"errors"
	"fmt"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
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
	}
}

func commandHelp() error {
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

func commandExit() error {
	return errors.New("")
}
