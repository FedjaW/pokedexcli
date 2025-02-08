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
	Count    int     `json:"count"`
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
		return LocationAreaResponse{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
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

type LocationAreaDetailsResponse struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func GetMapDetails(url string, cache *pokecache.Cache) (LocationAreaDetailsResponse, error) {
	val, ok := cache.Get(url)
	if ok {
		locationAreaDetailsResponse := LocationAreaDetailsResponse{}
		err := json.Unmarshal(val, &locationAreaDetailsResponse)
		if err != nil {
			return LocationAreaDetailsResponse{}, err
		}
		return locationAreaDetailsResponse, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return LocationAreaDetailsResponse{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 200 {
		return LocationAreaDetailsResponse{}, fmt.Errorf("response failed with status code: %d", res.StatusCode)
	}
	if err != nil {
		return LocationAreaDetailsResponse{}, err
	}

	locationAreaDetailsResponse := LocationAreaDetailsResponse{}
	err = json.Unmarshal(body, &locationAreaDetailsResponse)
	if err != nil {
		return LocationAreaDetailsResponse{}, err
	}

	cache.Add(url, body)

	return locationAreaDetailsResponse, nil
}
