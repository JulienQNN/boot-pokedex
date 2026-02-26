package main

import (
	"time"

	"github.com/JulienQNN/boot-pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 10*time.Second)
	cfg := &config{
		pokeapiClient:      pokeClient,
		caughtPokemonNames: make(map[string]pokeapi.Pokemon),
	}

	startRepl(cfg)
}
