package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/JulienQNN/boot-pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	prevLocationsURL *string
	nextLocationsURL *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}
		if cmd, ok := commandRegistry()[cleanedInput[0]]; ok {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Printf("Error executing command '%s': %s\n", cmd.name, err.Error())
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func commandRegistry() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the map of the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the map of the Pokemon world (backwards)",
			callback:    commandMapb,
		},
	}
}
