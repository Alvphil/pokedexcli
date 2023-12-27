package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	locationAreas, _ := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocation)

	for _, result := range locationAreas.Results {
		fmt.Printf("%v\n", result.Name)
	}

	return nil
}

func commandMapb() error {
	return nil
}
