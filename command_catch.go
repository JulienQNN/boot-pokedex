package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	pokemonResp, err := cfg.pokeapiClient.GetPokemon(args[0])
	pokemonName := pokemonResp.Name
	if err != nil {
		return err
	}

	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")

	const threshold = 50
	res := rand.Intn(pokemonResp.BaseExperience)

	if res > threshold {
		fmt.Printf("%v escaped!\n", pokemonName)
		return nil
	} else {
		fmt.Printf("%v caught!\n", pokemonName)
		cfg.caughtPokemonNames[pokemonName] = pokemonResp
	}

	return nil
}
