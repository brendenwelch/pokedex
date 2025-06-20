package main

import (
	"fmt"
	"math/rand"
	"os"
	"slices"
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
		"catch": {
			name:        "catch",
			description: "Attempt to catch the given Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect the given Pokemon, if already caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all of the Pokemon in the Pokedex",
			callback:    commandPokedex,
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

func commandCatch(cfg *config, pokemon *string) error {
	if pokemon == nil {
		return fmt.Errorf("error catching. no pokemon name provided")
	}

	res, err := cfg.Client.GetPokemon(pokemon)
	if err != nil {
		fmt.Println("error catching. pokemon not found")
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", *pokemon)
	if rand.Intn(res.BaseExperience) > 25 {
		fmt.Printf("%s escaped!\n", *pokemon)
		return nil
	}
	fmt.Printf("%s was caught!\n", *pokemon)
	cfg.Caught = append(cfg.Caught, *pokemon)

	return nil
}

func commandInspect(cfg *config, pokemon *string) error {
	if pokemon == nil {
		return fmt.Errorf("error inspecting. no pokemon name provided")
	}

	if !slices.Contains(cfg.Caught, *pokemon) {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	res, err := cfg.Client.GetPokemon(pokemon)
	if err != nil {
		fmt.Println("error inspecting. pokemon not found")
		return err
	}

	fmt.Printf("Name: %v\n", res.Name)
	fmt.Printf("Height: %v\n", res.Height)
	fmt.Printf("Weight: %v\n", res.Weight)
	fmt.Println("Stats:")
	for _, val := range res.Stats {
		fmt.Printf("-%v: %v\n", val.Stat.Name, val.BaseStat)
	}
	fmt.Println("Types:")
	for _, val := range res.Types {
		fmt.Printf("- %v\n", val.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *config, _ *string) error {
	if len(cfg.Caught) == 0 {
		fmt.Println("There are no pokemon in your Pokedex")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for _, val := range cfg.Caught {
		fmt.Printf("- %v\n", val)
	}

	return nil
}
