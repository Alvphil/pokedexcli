package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//pageURL *string

func (c *Client) ListLocationAreas(nextURL *string) (LocationAreas, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if nextURL != nil && *nextURL != "" {
		fullURL = *nextURL
	}

	//chack cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		//cache hit
		locations := LocationAreas{}
		err := json.Unmarshal(dat, &locations)
		if err != nil {
			// an error will be thrown if the JSON is invalid or has the wrong types
			// any missing fields will simply have their values in the struct set to their zero value
			fmt.Printf("Error decoding parameters: %s", err)
		}

		return locations, nil
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreas{}, err
	}

	if resp.StatusCode > 499 {
		return LocationAreas{}, fmt.Errorf("bad status: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, err
	}
	locations := LocationAreas{}
	err = json.Unmarshal(dat, &locations)

	//decoder := json.NewDecoder(resp.Body)

	//err = decoder.Decode(&locations)
	if err != nil {
		// an error will be thrown if the JSON is invalid or has the wrong types
		// any missing fields will simply have their values in the struct set to their zero value
		fmt.Printf("Error decoding parameters: %s", err)
	}

	c.cache.Add(fullURL, dat)

	return locations, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	//chack cache
	dat, ok := c.cache.Get(fullURL)
	if ok {
		//cache hit
		location := LocationArea{}
		err := json.Unmarshal(dat, &location)
		if err != nil {
			// an error will be thrown if the JSON is invalid or has the wrong types
			// any missing fields will simply have their values in the struct set to their zero value
			fmt.Printf("Error decoding parameters: %s", err)
		}

		return location, nil
	}
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	if resp.StatusCode > 499 {
		return LocationArea{}, fmt.Errorf("bad status: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}
	location := LocationArea{}
	err = json.Unmarshal(dat, &location)

	//decoder := json.NewDecoder(resp.Body)

	//err = decoder.Decode(&locations)
	if err != nil {
		// an error will be thrown if the JSON is invalid or has the wrong types
		// any missing fields will simply have their values in the struct set to their zero value
		fmt.Printf("Error decoding parameters: %s", err)
	}

	c.cache.Add(fullURL, dat)

	return location, nil
}
