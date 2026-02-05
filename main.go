package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func commandExit(cfg *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config, args []string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	// Get the commands registry
	commands := getCommands()

	// Dynamically generate help from registry
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()

	return nil
}

func commandMap(cfg *config, args []string) error {
	locationAreas, err := fetchLocationAreas(cfg, cfg.nextLocationURL)
	if err != nil {
		return err
	}

	// Update config with pagination URLs
	cfg.nextLocationURL = locationAreas.Next
	cfg.previousLocationURL = locationAreas.Previous

	// Print location area names
	for _, locationArea := range locationAreas.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}

func commandMapb(cfg *config, args []string) error {
	if cfg.previousLocationURL == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	locationAreas, err := fetchLocationAreas(cfg, cfg.previousLocationURL)
	if err != nil {
		return err
	}

	// Update config with pagination URLs
	cfg.nextLocationURL = locationAreas.Next
	cfg.previousLocationURL = locationAreas.Previous

	// Print location area names
	for _, locationArea := range locationAreas.Results {
		fmt.Println(locationArea.Name)
	}

	return nil
}

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("explore requires a location area name")
	}

	locationName := args[0]
	fmt.Printf("Exploring %s...\n", locationName)

	locationArea, err := fetchLocationArea(cfg, locationName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	if len(locationArea.PokemonEncounters) == 0 {
		fmt.Println("No Pokemon found in this area.")
	} else {
		for _, encounter := range locationArea.PokemonEncounters {
			fmt.Printf(" - %s\n", encounter.Pokemon.Name)
		}
	}

	return nil
}

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("catch requires a Pokemon name")
	}

	pokemonName := args[0]
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemon, err := fetchPokemon(cfg, pokemonName)
	if err != nil {
		return err
	}

	// Calculate catch probability based on base experience
	// Higher base experience = harder to catch
	// Formula: chance decreases as base experience increases
	// Base experience typically ranges from 50-300+
	catchChance := 50 // Base 50% chance
	if pokemon.BaseExperience > 0 {
		// Reduce chance based on base experience (more experienced = harder to catch)
		catchChance = max(10, 70-pokemon.BaseExperience/5) // Minimum 10% chance, max starts at 70%
	}

	// Generate random number 1-100
	rollResult := rand.Intn(100) + 1

	if rollResult <= catchChance {
		// Pokemon was caught!
		fmt.Printf("%s was caught!\n", pokemonName)
		// Add to caught Pokemon
		cfg.caughtPokemon[pokemonName] = pokemon
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		// Pokemon escaped
		fmt.Printf("%s escaped!\n", pokemonName)
	}

	return nil
}

func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("inspect requires a Pokemon name")
	}

	pokemonName := args[0]
	pokemon, exists := cfg.caughtPokemon[pokemonName]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	// Display stats
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	// Display types
	fmt.Println("Types:")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokemonType.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *config, args []string) error {
	fmt.Println("Your Pokedex:")
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You haven't caught any Pokemon yet!")
		return nil
	}

	for pokemonName := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", pokemonName)
	}

	return nil
}

// Helper function for max value (Go doesn't have built-in max for int)
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area to see Pokemon. Usage: explore <area_name>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a Pokemon. Usage: catch <pokemon_name>",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "View details of a caught Pokemon. Usage: inspect <pokemon_name>",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View all caught Pokemon",
			callback:    commandPokedex,
		},
	}
}

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Initialize config with cache and caught Pokemon storage
	cfg := &config{
		cache:         NewCache(5 * time.Minute),
		caughtPokemon: make(map[string]PokemonDetail),
	}

	// Create a scanner that reads from stdin
	scanner := bufio.NewScanner(os.Stdin)

	// Get the commands registry
	commands := getCommands()

	// Start infinite loop for REPL
	for {
		// Print prompt without newline
		fmt.Print("Pokedex > ")

		// Scan for user input (blocks until user presses enter)
		scanner.Scan()

		// Get the input text
		text := scanner.Text()

		// Clean the user's input using our cleanInput function
		cleaned := cleanInput(text)

		// If we have at least one word, process the command
		if len(cleaned) > 0 {
			commandName := cleaned[0]
			args := cleaned[1:] // Everything after the command name

			// Look up the command in the registry
			if cmd, exists := commands[commandName]; exists {
				// Call the callback function with arguments
				err := cmd.callback(cfg, args)
				if err != nil {
					fmt.Printf("Error: %v\n", err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
