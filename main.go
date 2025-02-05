package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Printf("Hello, World!")
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

