package models

import (
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

// GameCode represents standardized game codes
type GameCode string

const (
	FO1    GameCode = "FO1"
	FO2    GameCode = "FO2"
	FO3    GameCode = "FO3"
	FNV    GameCode = "FNV"
	FO4    GameCode = "FO4"
	FO76   GameCode = "FO76"
	FOS    GameCode = "FOS"    // Fallout Shelter
	FOSO   GameCode = "FOSO"   // Fallout Shelter Online
	FOSBR  GameCode = "FOSBR"  // Fallout Shelter
	FO76SD GameCode = "FO76SD" // Fallout 76 Steel Dawn
	FO76SR GameCode = "FO76SR" // Fallout 76 Steel Reign
	FOT    GameCode = "FOT"    // Fallout Tactics
	FOBOS  GameCode = "FOBOS"  // Fallout: Brotherhood of Steel
	FBGNC  GameCode = "FBGNC"  // Fallout Board Game: New California
	FOWW   GameCode = "FOWW"   // Fallout: Wasteland Warfare
)

// GameFullName returns the full name of a game from its code
func (g GameCode) GameFullName() string {
	gameNames := map[GameCode]string{
		FO1:    "Fallout",
		FO2:    "Fallout 2",
		FO3:    "Fallout 3",
		FNV:    "Fallout: New Vegas",
		FO4:    "Fallout 4",
		FO76:   "Fallout 76",
		FOS:    "Fallout Shelter",
		FOSO:   "Fallout Shelter Online",
		FOSBR:  "Fallout Shelter",
		FO76SD: "Fallout 76: Steel Dawn",
		FO76SR: "Fallout 76: Steel Reign",
		FOT:    "Fallout Tactics",
		FOBOS:  "Fallout: Brotherhood of Steel",
		FBGNC:  "Fallout Board Game: New California",
		FOWW:   "Fallout: Wasteland Warfare",
	}

	if name, exists := gameNames[g]; exists {
		return name
	}
	return string(g)
}

// NormalizeGameCodes converts comma-separated game codes to slice
func NormalizeGameCodes(gamesStr string) []string {
	if gamesStr == "" {
		return []string{}
	}

	games := strings.Split(gamesStr, ",")
	var normalized []string

	for _, game := range games {
		game = strings.TrimSpace(game)
		if game != "" {
			normalized = append(normalized, game)
		}
	}

	return normalized
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

// CharacterFilter represents filters for character queries
type CharacterFilter struct {
	Game        string `json:"game"`
	Race        string `json:"race"`
	Gender      string `json:"gender"`
	Status      string `json:"status"`
	Affiliation string `json:"affiliation"`
	Playable    *bool  `json:"playable"`
	Limit       int    `json:"limit"`
	Offset      int    `json:"offset"`
}
