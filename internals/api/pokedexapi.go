package api

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

// https://pokeapi.co/docs/v2#locations-section
// https://jsonlint.com
// https://mholt.github.io/json-to-go/
// https://pokeapi.co/api/v2/location-area
type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetMap(url string) (LocationAreasResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return LocationAreasResponse{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
		return LocationAreasResponse{}, errors.New("error")
	}
	if err != nil {
		log.Fatal(err)
	}
	locationAreas := LocationAreasResponse{}
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		log.Fatal(err)
		return LocationAreasResponse{}, err
	}
	return locationAreas, nil
}
