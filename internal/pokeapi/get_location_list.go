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

func (cli *Client) ListLocations(pageURL *string) (RespOverviewLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := cli.cache.Get(url); ok {
		locationsResp := RespOverviewLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespOverviewLocations{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespOverviewLocations{}, err
	}

	resp, err := cli.httpClient.Do(req)
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

	cli.cache.Add(url, dat)
	return locationsResp, nil
}

type RespPokemonByLocation struct {
	GameIndex         int    `json:"game_index"`
	ID                int    `json:"id"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (cli *Client) ListPokemonByLocation(identifier string) (RespPokemonByLocation, error) {
	url := baseURL + "/location-area/" + identifier

	if val, ok := cli.cache.Get(url); ok {
		localPokemonResp := RespPokemonByLocation{}
		err := json.Unmarshal(val, &localPokemonResp)
		if err != nil {
			return RespPokemonByLocation{}, err
		}
		return localPokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemonByLocation{}, err
	}

	resp, err := cli.httpClient.Do(req)
	if err != nil {
		return RespPokemonByLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemonByLocation{}, err
	}

	localPokemonResp := RespPokemonByLocation{}
	err = json.Unmarshal(dat, &localPokemonResp)
	if err != nil {
		return RespPokemonByLocation{}, err
	}

	cli.cache.Add(url, dat)
	return localPokemonResp, nil
}
