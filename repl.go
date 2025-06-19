package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/brendenwelch/pokedex/internal/pokeapi"
)

type config struct {
	Client   pokeapi.Client
	Next     *string
	Previous *string
	Caught   []string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		userInput := cleanInput(scanner.Text())
		if len(userInput) == 0 {
			continue
		}

		commandName := userInput[0]
		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}

		var arg *string
		if len(userInput) >= 2 {
			arg = &userInput[1]
		}

		if err := command.callback(cfg, arg); err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	return strings.Fields(text)
}
