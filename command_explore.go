package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, words ...string) error {
	if len(words) == 0 {
		return errors.New("no location provided")
	}
	
	locationsResp, err := cfg.pokeapiClient.ExploreLocation(words[0])
	if err != nil {
		return err
	}


	for _, loc := range locationsResp.PokemonEncounters {
		fmt.Println(loc.Pokemon.Name)
	}
	return nil
}
