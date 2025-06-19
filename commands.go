package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show usage information",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Show next 20 location areas in the Pokemon world",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Show previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List all of the Pokemon located in the provided location area",
			callback:    commandExplore,
		},
	}
}

func commandExit(cfg *config, _ *string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, _ *string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func commandMapf(cfg *config, _ *string) error {
	res, err := cfg.Client.ListLocations(cfg.Next)
	if err != nil {
		return err
	}

	print("help")
	cfg.Next = res.Next
	cfg.Previous = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, _ *string) error {
	if cfg.Previous == nil {
		return fmt.Errorf("you're on the first page")
	}

	res, err := cfg.Client.ListLocations(cfg.Previous)
	if err != nil {
		return err
	}

	cfg.Next = res.Next
	cfg.Previous = res.Previous

	for _, loc := range res.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(cfg *config, location *string) error {
	if location == nil {
		return fmt.Errorf("error exploring. no location area provided")
	}

	res, err := cfg.Client.GetLocation(location)
	if err != nil {
		return err
	}

	for _, encounter := range res.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}
