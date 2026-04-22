package main

import (
	"errors"
	"fmt"
	"os"
)

func commandExit(cfg *config, args ...string) error {
	if len(args) != 0 {
		return errors.New("No argument is supported for this command")
	}
	print("Closing the Pokedex... Goodbye!\n")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args ...string) error {
	if len(args) != 0 {
		return errors.New("No argument is supported for this command")
	}
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

func commandMapf(cfg *config, args ...string) error {
	if len(args) != 0 {
		return errors.New("No argument is supported for this command")
	}
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if len(args) != 0 {
		return errors.New("No argument is supported for this command")
	}
	if cfg.prevLocationsURL == nil {
		return errors.New("You're on the first page\n")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("Needs at least 1 argument for this command")
	}

	name := args[0]
	localPokemonResp, err := cfg.pokeapiClient.ListPokemonByLocation(name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", localPokemonResp.Name)
	fmt.Println("Found Pokemon:")

	for _, encounter := range localPokemonResp.PokemonEncounters {
		fmt.Printf(" - %v\n", encounter.Pokemon.Name)
	}
	return nil
}
