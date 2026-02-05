package main

// Config holds the application state for pagination
type config struct {
	nextLocationURL     *string
	previousLocationURL *string
	cache               *Cache
	caughtPokemon       map[string]PokemonDetail
}

// LocationAreasResponse represents the paginated response from PokeAPI location-area endpoint
type LocationAreasResponse struct {
	Count    int                  `json:"count"`
	Next     *string              `json:"next"`
	Previous *string              `json:"previous"`
	Results  []LocationAreaResult `json:"results"`
}

// LocationAreaResult represents a single location area in the results
type LocationAreaResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// LocationArea represents detailed information about a specific location area
type LocationArea struct {
	ID                int                `json:"id"`
	Name              string             `json:"name"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

// PokemonEncounter represents a Pokemon encounter in a location area
type PokemonEncounter struct {
	Pokemon Pokemon `json:"pokemon"`
}

// Pokemon represents basic Pokemon information
type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// PokemonDetail represents detailed Pokemon information from the Pokemon endpoint
type PokemonDetail struct {
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

// PokemonStat represents a Pokemon's individual stat
type PokemonStat struct {
	BaseStat int `json:"base_stat"`
	Stat     struct {
		Name string `json:"name"`
	} `json:"stat"`
}

// PokemonType represents a Pokemon's type
type PokemonType struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}
