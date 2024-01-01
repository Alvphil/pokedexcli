package main

import (
	"fmt"
)

func commandMap(cfg *config) error {
	locationAreas, _ := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocation)

	if locationAreas.Next != nil {
		*cfg.nextLocation = *locationAreas.Next
	}
	if locationAreas.Previous != nil {
		*cfg.previousLocation = *locationAreas.Previous
	}

	for _, result := range locationAreas.Results {
		fmt.Printf("%v\n", result.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	locationAreas, _ := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocation)

	if locationAreas.Previous != nil {
		*cfg.previousLocation = *locationAreas.Previous
	}
	*cfg.nextLocation = *locationAreas.Next

	for _, result := range locationAreas.Results {
		fmt.Printf("%v\n", result.Name)
	}
	return nil
}
