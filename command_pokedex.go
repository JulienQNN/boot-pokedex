package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {
	if len(cfg.caughtPokemonNames) == 0 {
		fmt.Println("No Pokemon caught yet")
	}
	for name := range cfg.caughtPokemonNames {
		fmt.Printf("- %s\n", name)
	}
	return nil
}
