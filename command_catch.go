package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name area provided")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	const threshold = 50
	randnum := rand.Intn(pokemon.BaseExperience)
	if randnum > threshold {
		return fmt.Errorf("failed to catch %s!", pokemonName)
	}

	fmt.Printf("You caught %s\n", pokemonName)
	return nil
}