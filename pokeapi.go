package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// fetchLocationAreas makes a GET request to the specified URL and returns location areas
// Uses caching to improve performance on repeated requests
func fetchLocationAreas(cfg *config, pageURL *string) (LocationAreasResponse, error) {
	url := "https://pokeapi.co/api/v2/location-area/"
	if pageURL != nil {
		url = *pageURL
	}

	// Check cache first
	if cachedData, found := cfg.cache.Get(url); found {
		fmt.Println("Cache hit! Using cached data.")
		var locationAreas LocationAreasResponse
		err := json.Unmarshal(cachedData, &locationAreas)
		if err != nil {
			return LocationAreasResponse{}, fmt.Errorf("error parsing cached JSON: %w", err)
		}
		return locationAreas, nil
	}

	// Cache miss - make HTTP request
	fmt.Println("Cache miss! Making HTTP request.")
	resp, err := http.Get(url)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return LocationAreasResponse{}, fmt.Errorf("bad response: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("error reading response: %w", err)
	}

	// Cache the response for future use
	cfg.cache.Add(url, body)

	var locationAreas LocationAreasResponse
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return LocationAreasResponse{}, fmt.Errorf("error parsing JSON: %w", err)
	}

	return locationAreas, nil
}

// fetchLocationArea makes a GET request for a specific location area and returns detailed info
func fetchLocationArea(cfg *config, locationName string) (LocationArea, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", locationName)

	// Check cache first
	if cachedData, found := cfg.cache.Get(url); found {
		fmt.Printf("Cache hit! Using cached data for %s.\n", locationName)
		var locationArea LocationArea
		err := json.Unmarshal(cachedData, &locationArea)
		if err != nil {
			return LocationArea{}, fmt.Errorf("error parsing cached JSON: %w", err)
		}
		return locationArea, nil
	}

	// Cache miss - make HTTP request
	fmt.Printf("Cache miss! Making HTTP request for %s.\n", locationName)
	resp, err := http.Get(url)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return LocationArea{}, fmt.Errorf("bad response: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error reading response: %w", err)
	}

	// Cache the response for future use
	cfg.cache.Add(url, body)

	var locationArea LocationArea
	err = json.Unmarshal(body, &locationArea)
	if err != nil {
		return LocationArea{}, fmt.Errorf("error parsing JSON: %w", err)
	}

	return locationArea, nil
}

// fetchPokemon makes a GET request for a specific Pokemon and returns detailed info
func fetchPokemon(cfg *config, pokemonName string) (PokemonDetail, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemonName)

	// Check cache first
	if cachedData, found := cfg.cache.Get(url); found {
		fmt.Printf("Cache hit! Using cached data for %s.\n", pokemonName)
		var pokemon PokemonDetail
		err := json.Unmarshal(cachedData, &pokemon)
		if err != nil {
			return PokemonDetail{}, fmt.Errorf("error parsing cached JSON: %w", err)
		}
		return pokemon, nil
	}

	// Cache miss - make HTTP request
	fmt.Printf("Cache miss! Making HTTP request for %s.\n", pokemonName)
	resp, err := http.Get(url)
	if err != nil {
		return PokemonDetail{}, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return PokemonDetail{}, fmt.Errorf("bad response: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonDetail{}, fmt.Errorf("error reading response: %w", err)
	}

	// Cache the response for future use
	cfg.cache.Add(url, body)

	var pokemon PokemonDetail
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return PokemonDetail{}, fmt.Errorf("error parsing JSON: %w", err)
	}

	return pokemon, nil
}
