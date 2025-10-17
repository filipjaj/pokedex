package main

import (
	"time"

	"github.com/bootdotdev/pokedexcli/internal/pokeapi"
	"github.com/bootdotdev/pokedexcli/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(time.Minute * 10)
	pokeClient := pokeapi.NewClient(5 * time.Second, *cache)
	cfg := &config{
		pokeapiClient: pokeClient,
	
		pokedex:       make(map[string]pokeapi.PokemonDetails),
	}

	startRepl(cfg)
}
