package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fedjaw/pokedexcli/internal/pokecache"
)

// https://pokeapi.co/docs/v2#locations-section
// https://jsonlint.com
// https://mholt.github.io/json-to-go/
// https://pokeapi.co/api/v2/location-area
type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetMap(url string, cache *pokecache.Cache) (LocationAreaResponse, error) {
    val, ok := cache.Get(url)
    if ok {
        fmt.Println("from cache")
        locationAreaResponse := LocationAreaResponse{}
        err := json.Unmarshal(val, &locationAreaResponse)
        if err != nil {
            return LocationAreaResponse{}, err
        }
        return locationAreaResponse, nil
    }

    res, err := http.Get(url)
    if err != nil {
        return LocationAreaResponse{}, err
    }
    body, err := io.ReadAll(res.Body)
    res.Body.Close()

    if res.StatusCode > 200 {
        return LocationAreaResponse{}, fmt.Errorf("Response failed with status code: %d\n", res.StatusCode)
    }
    if err != nil {
        return LocationAreaResponse{}, err
    }

    locationAreaResponse := LocationAreaResponse{}
    err = json.Unmarshal(body, &locationAreaResponse)
    if err != nil {
        return LocationAreaResponse{}, err
    }

    cache.Add(url, body)

    return locationAreaResponse, nil
}
