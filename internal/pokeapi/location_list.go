package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationArea, error) {
	fmt.Println("Trying get cache ..")
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	cache, ok := c.cache.Get(url)
	if ok {
		fmt.Println("--- Cache hit! ---")
		locationsResp := LocationArea{}
		err := json.Unmarshal(cache, &locationsResp)
		if err != nil {
			return LocationArea{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	locationsResp := LocationArea{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(url, data)

	return locationsResp, nil
}
