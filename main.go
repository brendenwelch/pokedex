package main

import (
	"time"

	"github.com/brendenwelch/pokedex/internal/pokeapi"
)

func main() {
	client := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		Client: client,
	}

	startRepl(cfg)
}
