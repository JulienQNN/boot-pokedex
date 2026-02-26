package main

import (
	"fmt"
)

func commandInspect(cfg *config, args []string) error {
	pokemonResp, _ := cfg.caughtPokemonNames[args[0]]

	if pokemonResp.Name == "" {
		fmt.Printf("Pokemon %s not found in caught Pokemon list\n", args[0])
		return nil
	}
	fmt.Printf("Name: %s\n", pokemonResp.Name)
	fmt.Printf("Height: %d\n", pokemonResp.Height)
	fmt.Printf("Weight: %d\n", pokemonResp.Weight)

	fmt.Println("Stats:")
	fmt.Printf("  - hp: %d\n", pokemonResp.Stats[0].BaseStat)
	fmt.Printf("  - attack: %d\n", pokemonResp.Stats[1].BaseStat)
	fmt.Printf("  - defense: %d\n", pokemonResp.Stats[2].BaseStat)
	fmt.Printf("  - special-attack: %d\n", pokemonResp.Stats[3].BaseStat)
	fmt.Printf("  - special-defense: %d\n", pokemonResp.Stats[4].BaseStat)
	fmt.Printf("  - speed: %d\n", pokemonResp.Stats[5].BaseStat)
	fmt.Println("Types:")
	for _, t := range pokemonResp.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}
