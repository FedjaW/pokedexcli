package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	commands := getCliCommandsMap()
	for {
		fmt.Print("pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if _, ok := commands[input]; !ok {
			fmt.Println("Invalid input. Enter help to show all options.")
			continue
		}
		error := commands[input].callback()
		if error != nil {
			fmt.Println(error.Error())
			break
		}
	}
}

func getCliCommandsMap() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	commands := getCliCommandsMap()
	for k, v := range commands {
		fmt.Print(k)
		fmt.Print(": ")
		fmt.Println(v.description)
	}
	return nil
}

func commandExit() error {
	return errors.New("Bye")
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
