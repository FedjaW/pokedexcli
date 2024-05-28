package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fedjaw/pokedexcli/commands"
)

func startRepl() {
	const PROMPT string = "pokedex > "
	cmds := commands.GetCliCommandsMap()
	cfg := commands.NewConfig()
	for {
		fmt.Print(PROMPT)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		if _, ok := cmds[input]; !ok {
			fmt.Println("Invalid input. Enter help to show all options.")
			continue
		}
		err := cmds[input].Callback(&cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
