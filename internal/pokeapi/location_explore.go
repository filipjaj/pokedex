package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ExploreLocation -
func (c *Client) ExploreLocation(name string) (ExploreLocation, error) {
	url := baseURL + "/location-area/" + name
	fmt.Println(url)

	cachedExploredLocation, ok := c.cache.Get(url)
	if ok {
		var locationsResp ExploreLocation
		err := json.Unmarshal(cachedExploredLocation, &locationsResp)
		if err != nil {
			return ExploreLocation{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ExploreLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ExploreLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return ExploreLocation{}, err
	}

	locationsResp := ExploreLocation{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return ExploreLocation{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
