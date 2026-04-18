package main

import (
	"time"

	"github.com/its-PKN-2k4/pokedex-go/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Minute, 5*time.Second)
	cfg := &config{
		pokeapiClient: &pokeClient,
	}

	startREPL(cfg)
}
