package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// GetPokemon retrieves Pokemon details by name
func (c *Client) GetPokemon(name string) (PokemonDetails, error) {
	url := baseURL + "/pokemon/" + name

	cachedPokemon, ok := c.cache.Get(url)
	if ok {
		var pokemonResp PokemonDetails
		err := json.Unmarshal(cachedPokemon, &pokemonResp)
		if err != nil {
			return PokemonDetails{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonDetails{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonDetails{}, err
	}

	pokemonResp := PokemonDetails{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return PokemonDetails{}, err
	}

	c.cache.Add(url, dat)
	return pokemonResp, nil
}
