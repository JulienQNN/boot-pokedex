package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(pageURL *string, location string) (DescribeLocation, error) {
	fmt.Println("Trying get cache ..")
	url := baseURL + "/location-area/" + location
	if pageURL != nil {
		url = *pageURL
	}
	cache, ok := c.cache.Get(url)
	if ok {
		fmt.Println("--- Cache hit! ---")
		locationsResp := DescribeLocation{}
		err := json.Unmarshal(cache, &locationsResp)
		if err != nil {
			return DescribeLocation{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return DescribeLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return DescribeLocation{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return DescribeLocation{}, err
	}

	locationsResp := DescribeLocation{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return DescribeLocation{}, err
	}

	c.cache.Add(url, data)

	return locationsResp, nil
}
