package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fedjaw/pokedexcli/commands"
)

func startRepl() {
	const PROMPT string = "pokedex > "
	commands := commands.GetCliCommandsMap()
	for {
		fmt.Print(PROMPT)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		if _, ok := commands[input]; !ok {
			fmt.Println("Invalid input. Enter help to show all options.")
			continue
		}
		commands[input].Callback()
	}
}
