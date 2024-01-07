package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	//check cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		//cache hit
		pokemon := Pokemon{}
		err := json.Unmarshal(dat, &pokemon)
		if err != nil {
			// an error will be thrown if the JSON is invalid or has the wrong types
			// any missing fields will simply have their values in the struct set to their zero value
			fmt.Printf("Error decoding parameters: %s", err)
		}

		return pokemon, nil
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	if resp.StatusCode > 499 {
		return Pokemon{}, fmt.Errorf("bad status: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	pokemon := Pokemon{}
	err = json.Unmarshal(dat, &pokemon)

	//decoder := json.NewDecoder(resp.Body)

	//err = decoder.Decode(&locations)
	if err != nil {
		// an error will be thrown if the JSON is invalid or has the wrong types
		// any missing fields will simply have their values in the struct set to their zero value
		fmt.Printf("Error decoding parameters: %s", err)
	}

	c.cache.Add(fullURL, dat)

	return pokemon, nil
}
