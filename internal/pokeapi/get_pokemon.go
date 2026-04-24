package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type RespPokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	}
}

func (cli *Client) CatchPokemon(pokemon string) (RespPokemon, error) {
	url := baseURL + "/pokemon/" + pokemon

	if val, ok := cli.cache.Get(url); ok {
		pokemonResp := RespPokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return RespPokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	resp, err := cli.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	pokemonResp := RespPokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokemon{}, err
	}

	cli.cache.Add(url, dat)
	return pokemonResp, nil
}
