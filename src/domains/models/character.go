package models

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Character represents a Fallout character with all relevant game information
type Character struct {
	ID          int      `json:"id" db:"id"`
	Name        string   `json:"name" db:"name"`
	WikiTitle   string   `json:"wiki_title" db:"wiki_title"` // Original wiki page title
	Games       []string `json:"games" db:"games"`           // Main games where character appears
	Mentions    []string `json:"mentions" db:"mentions"`     // Games where character is mentioned
	Race        string   `json:"race" db:"race"`
	Gender      string   `json:"gender" db:"gender"`
	Status      string   `json:"status" db:"status"` // Alive, Deceased, Unknown
	Affiliation []string `json:"affiliation" db:"affiliation"`
	Role        string   `json:"role" db:"role"`
	Titles      []string `json:"titles" db:"titles"`
	MainGame    string   `json:"main_game" db:"main_game"` // Primary game of origin
	ImageURL    string   `json:"image_url" db:"image_url"`
}

// GetMainGame returns the primary game of origin (first in games list)
func (c *Character) GetMainGame() string {
	if len(c.Games) > 0 {
		return c.Games[0]
	}
	return ""
}

// IsPlayable determines if character is playable (like Vault Dweller)
func (c *Character) IsPlayable() bool {
	playableCharacters := []string{
		"Vault Dweller",
		"Chosen One",
		"Lone Wanderer",
		"Courier",
		"Sole Survivor",
		"Resident",
	}

	for _, playable := range playableCharacters {
		if strings.Contains(c.Name, playable) {
			return true
		}
	}

	return false
}

// CleanRace removes wiki markup from race field
func (c *Character) CleanRace() string {
	race := c.Race
	// Remove wiki links [[Human]] -> Human
	race = strings.ReplaceAll(race, "[[", "")
	race = strings.ReplaceAll(race, "]]", "")

	// Handle complex race descriptions
	if strings.Contains(race, "<br") {
		// Take first race if multiple are listed
		parts := strings.Split(race, "<br")
		if len(parts) > 0 {
			race = strings.TrimSpace(parts[0])
		}
	}

	return race
}

// String implements the fmt.Stringer interface for Character
// This method provides a readable text representation of the character
func (c *Character) String() string {
	if c == nil {
		return "<nil Character>"
	}

	var builder strings.Builder
	builder.Grow(512) // Pre-allocate memory to avoid reallocations

	// Header with name and ID
	builder.WriteString("Character")
	if c.ID != 0 {
		builder.WriteString(fmt.Sprintf(" #%T", c.ID))
	}
	if c.Name != "" {
		builder.WriteString(fmt.Sprintf(": %s", c.Name))
	}
	builder.WriteString("\n")

	// Basic information
	if c.Race != "" {
		builder.WriteString(fmt.Sprintf("  Race: %s\n", c.Race))
	}
	if c.Gender != "" {
		builder.WriteString(fmt.Sprintf("  Gender: %s\n", c.Gender))
	}
	if c.Status != "" {
		builder.WriteString(fmt.Sprintf("  Status: %s\n", c.Status))
	}
	if c.Role != "" {
		builder.WriteString(fmt.Sprintf("  Role: %s\n", c.Role))
	}

	// Main game
	if c.MainGame != "" {
		builder.WriteString(fmt.Sprintf("  Main Game: %s\n", c.MainGame))
	}

	// Games where the character appears
	if len(c.Games) > 0 {
		builder.WriteString("  Appears in: ")
		gameNames := make([]string, 0, len(c.Games))
		for _, game := range c.Games {
			if game != "" {
				// Try to convert to GameCode to get full name
				gameCode := GameCode(game)
				fullName := gameCode.GameFullName()
				if fullName != "" {
					gameNames = append(gameNames, fullName)
				} else {
					gameNames = append(gameNames, game)
				}
			}
		}
		builder.WriteString(strings.Join(gameNames, ", "))
		builder.WriteString("\n")
	}

	// Mentions
	if len(c.Mentions) > 0 {
		builder.WriteString("  Mentioned in: ")
		mentionNames := make([]string, 0, len(c.Mentions))
		for _, mention := range c.Mentions {
			if mention != "" {
				gameCode := GameCode(mention)
				fullName := gameCode.GameFullName()
				if fullName != "" {
					mentionNames = append(mentionNames, fullName)
				} else {
					mentionNames = append(mentionNames, mention)
				}
			}
		}
		builder.WriteString(strings.Join(mentionNames, ", "))
		builder.WriteString("\n")
	}

	// Affiliations
	if len(c.Affiliation) > 0 {
		builder.WriteString("  Affiliations: ")
		affiliations := make([]string, 0, len(c.Affiliation))
		for _, affiliation := range c.Affiliation {
			if affiliation != "" {
				affiliations = append(affiliations, affiliation)
			}
		}
		builder.WriteString(strings.Join(affiliations, ", "))
		builder.WriteString("\n")
	}

	// Titles
	if len(c.Titles) > 0 {
		builder.WriteString("  Titles: ")
		titles := make([]string, 0, len(c.Titles))
		for _, title := range c.Titles {
			if title != "" {
				titles = append(titles, title)
			}
		}
		builder.WriteString(strings.Join(titles, ", "))
		builder.WriteString("\n")
	}

	// Wiki information and image
	if c.WikiTitle != "" {
		builder.WriteString(fmt.Sprintf("  Wiki: %s\n", c.WikiTitle))
	}
	if c.ImageURL != "" {
		builder.WriteString(fmt.Sprintf("  Image: %s\n", c.ImageURL))
	}

	return builder.String()
}

// StringCompact returns a compact version for logs or short displays
func (c *Character) StringCompact() string {
	if c == nil {
		return "<nil>"
	}

	var parts []string

	if c.Name != "" {
		parts = append(parts, c.Name)
	}
	if c.ID != 0 {
		parts = append(parts, fmt.Sprintf("ID:%d", c.ID))
	}
	if c.Race != "" {
		parts = append(parts, c.Race)
	}
	if c.MainGame != "" {
		parts = append(parts, c.MainGame)
	}

	if len(parts) == 0 {
		return "Character{}"
	}

	return fmt.Sprintf("Character{%s}", strings.Join(parts, ", "))
}

// StringJSON returns a formatted JSON representation (useful for debugging)
func (c *Character) StringJSON() string {
	if c == nil {
		return "null"
	}

	jsonBytes, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Sprintf("Character{JSON marshal error: %v}", err)
	}
	return string(jsonBytes)
}
