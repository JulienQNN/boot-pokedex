package main

import (
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	locationResp, err := cfg.pokeapiClient.GetLocation(cfg.nextLocationsURL, args[0])
	if err != nil {
		return err
	}

	fmt.Println("Exploring " + args[0] + "...")
	if len(locationResp.PokemonEncounters) != 0 {
		fmt.Println("Found Pokemon:")
	}
	for _, result := range locationResp.PokemonEncounters {
		fmt.Println("-", result.Pokemon.Name)
	}
	return nil
}
