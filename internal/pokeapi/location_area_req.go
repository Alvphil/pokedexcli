package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//pageURL *string

func (c *Client) ListLocationAreas(nextURL *string) (locationAreas, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if nextURL != nil && *nextURL != "" {
		fullURL = *nextURL
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return locationAreas{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreas{}, err
	}

	if resp.StatusCode > 499 {
		return locationAreas{}, fmt.Errorf("bad status: %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	locations := locationAreas{}
	err = decoder.Decode(&locations)
	if err != nil {
		// an error will be thrown if the JSON is invalid or has the wrong types
		// any missing fields will simply have their values in the struct set to their zero value
		fmt.Printf("Error decoding parameters: %s", err)
	}


	if locations.Next != nil {
		*nextURL = *locations.Next
	}
	return locations, nil
}
