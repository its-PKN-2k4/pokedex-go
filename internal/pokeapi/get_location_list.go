package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type RespOverviewLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocations(pageURL *string) (RespOverviewLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespOverviewLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespOverviewLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespOverviewLocations{}, err
	}

	locationsResp := RespOverviewLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespOverviewLocations{}, err
	}

	return locationsResp, nil
}
