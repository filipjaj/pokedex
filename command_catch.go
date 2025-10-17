package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, words ...string) error {
	if len(words) == 0 {
		return errors.New("no pokemon provided")
	}

	pokemonName := words[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}


	const maxThreshold = 500
	catchThreshold := maxThreshold - pokemon.BaseExperience
	if catchThreshold < 0 {
		catchThreshold = 0
	}

	randomValue := rand.Intn(maxThreshold)

	if randomValue < catchThreshold {
		// Caught!
		cfg.pokedex[pokemon.Name] = pokemon
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		// Escaped
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}
