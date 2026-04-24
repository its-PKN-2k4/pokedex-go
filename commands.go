package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
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

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("Needs at least 1 argument for this command")
	}
	pokemon := args[0]
	pokemonInfo, err := cfg.pokeapiClient.CatchPokemon(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonInfo.Name)
	seed := 4
	offRate := pokemonInfo.BaseExperience % 15
	success := rand.Intn(int(math.Max(float64(seed), float64(offRate))))

	if success%seed != 0 {
		fmt.Printf("%s escaped!\n", pokemonInfo.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemonInfo.Name)
	cfg.caughtPokemon[pokemonInfo.Name] = pokemonInfo
	fmt.Printf("now you can inspect %s's info!\n", pokemonInfo.Name)
	return nil
}

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("Needs at least 1 argument for this command")
	}

	pokemon := args[0]
	data, exist := cfg.caughtPokemon[pokemon]
	if !exist {
		fmt.Printf("you have not caught %s, no info found\n", pokemon)
		return nil
	}

	fmt.Printf("Name: %s\n", data.Name)
	fmt.Printf("ID: %v\n", data.ID)
	fmt.Printf("Height: %v\n", data.Height)
	fmt.Printf("Weight: %v\n", data.Weight)
	fmt.Printf("WeighBase Experience: %v\n", data.BaseExperience)

	fmt.Println("Stats: ")
	for _, stats := range data.Stats {
		fmt.Printf("	-%s: %v\n", stats.Stat.Name, stats.BaseStat)
	}

	fmt.Println("Types: ")
	for _, types := range data.Types {
		fmt.Printf("	-%s\n", types.Type.Name)
	}

	fmt.Println("Abilities: ")
	for _, abilities := range data.Abilities {
		fmt.Printf("	-%s (is_hidden: %v)\n", abilities.Ability.Name, abilities.IsHidden)
	}

	return nil
}
