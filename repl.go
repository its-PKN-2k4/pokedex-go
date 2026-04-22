package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/its-PKN-2k4/pokedex-go/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

type config struct {
	pokeapiClient    *pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	running := true
	for running {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleanedInput := cleanInput(input)
		if len(cleanedInput) == 0 {
			continue
		}

		cmdName := cleanedInput[0]
		args := []string{}
		if len(cleanedInput) > 0 {
			args = cleanedInput[1:]
		}
		if cmd, exist := getCommands()[cmdName]; exist {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Printf("Error: %v", err)
			}
		} else {
			fmt.Println("Command not found")
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Show how to use the Pokedex",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Get all pokemon that can be encountered in a location area",
			callback:    commandExplore,
		},
	}
}
