package main

import (
	"time"

	"github.com/Alvphil/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocation     *string
	previousLocation *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func main() {
	defaultNextLocation := ""
	defaultPreviousLocation := ""

	cfg := config{
		pokeapiClient:    pokeapi.NewClient(time.Hour),
		nextLocation:     &defaultNextLocation,
		previousLocation: &defaultPreviousLocation,
		caughtPokemon:    make(map[string]pokeapi.Pokemon),
	}
	startRepl(&cfg)
}
