package main

import "github.com/Alvphil/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocation     *string
	previousLocation *string
}

func main() {
	defaultNextLocation := ""
	defaultPreviousLocation := ""

	cfg := config{
		pokeapiClient:    pokeapi.NewClient(),
		nextLocation:     &defaultNextLocation,
		previousLocation: &defaultPreviousLocation,
	}
	startRepl(&cfg)
}
