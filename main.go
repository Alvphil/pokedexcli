package main

import (
	"time"

	"github.com/Alvphil/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocation     *string
	previousLocation *string
}

func main() {
	defaultNextLocation := ""
	defaultPreviousLocation := ""

	cfg := config{
		pokeapiClient:    pokeapi.NewClient(time.Hour),
		nextLocation:     &defaultNextLocation,
		previousLocation: &defaultPreviousLocation,
	}
	startRepl(&cfg)
}
