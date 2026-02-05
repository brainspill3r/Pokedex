package main

import (
	"strings"
)

// cleanInput splits the user's input into words based on whitespace,
// lowercases the input, and trims leading/trailing whitespace
func cleanInput(text string) []string {
	// Convert to lowercase and trim whitespace
	lowered := strings.ToLower(strings.TrimSpace(text))

	// If the trimmed string is empty, return empty slice
	if lowered == "" {
		return []string{}
	}

	// Split by whitespace and return
	return strings.Fields(lowered)
}
