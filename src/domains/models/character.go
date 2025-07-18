package models

// Character represents a Fallout character with all game-relevant properties
type Character struct {
	ID          int      `json:"id" db:"id"`
	Name        string   `json:"name" db:"name"`
	Games       []string `json:"games" db:"games"`       // Games where character appears
	Mentions    []string `json:"mentions" db:"mentions"` // Games where character is mentioned
	Race        string   `json:"race" db:"race"`
	Gender      string   `json:"gender" db:"gender"`
	Status      string   `json:"status" db:"status"` // Alive, Deceased, Unknown
	Affiliation []string `json:"affiliation" db:"affiliation"`
	Role        string   `json:"role" db:"role"`
	Titles      []string `json:"titles" db:"titles"`
	MainGame    string   `json:"main_game" db:"main_game"` // Primary game of appearance
	ImageURL    string   `json:"image" db:"image"`
}

// GameCode represents the mapping between wiki game codes and full names
type GameCode struct {
	Code     string `json:"code"`
	FullName string `json:"full_name"`
}

// Common game codes used in Fallout wiki
var GameCodes = map[string]string{
	"FO1":    "Fallout",
	"FO2":    "Fallout 2",
	"FO3":    "Fallout 3",
	"FO4":    "Fallout 4",
	"FNV":    "Fallout: New Vegas",
	"FO76":   "Fallout 76",
	"FOS":    "Fallout Shelter",
	"FOSO":   "Fallout Shelter Online",
	"FO76SD": "Fallout 76: Steel Dawn",
	"FO76SR": "Fallout 76: Steel Reign",
	"FOT":    "Fallout Tactics",
	"FOBOS":  "Fallout: Brotherhood of Steel",
	"FBGNC":  "Fallout Board Game: New California",
	"FOWW":   "Fallout: Wasteland Warfare",
}

// GetMainGameFullName returns the full name of the main game
func (c *Character) GetMainGameFullName() string {
	if fullName, exists := GameCodes[c.MainGame]; exists {
		return fullName
	}
	return c.MainGame
}

// GetGamesFullNames returns full names for all games
func (c *Character) GetGamesFullNames() []string {
	var fullNames []string
	for _, code := range c.Games {
		if fullName, exists := GameCodes[code]; exists {
			fullNames = append(fullNames, fullName)
		} else {
			fullNames = append(fullNames, code)
		}
	}
	return fullNames
}

// IsMainCharacter checks if character appears in main games (not mentions only)
func (c *Character) IsMainCharacter() bool {
	return len(c.Games) > 0
}

// HasAffiliation checks if character is affiliated with a specific faction
func (c *Character) HasAffiliation(faction string) bool {
	for _, aff := range c.Affiliation {
		if aff == faction {
			return true
		}
	}
	return false
}

// CharacterFilter represents filters for character queries
type CharacterFilter struct {
	Game        string   `json:"game"`
	Race        string   `json:"race"`
	Gender      string   `json:"gender"`
	Status      string   `json:"status"`
	Affiliation []string `json:"affiliation"`
	Limit       int      `json:"limit"`
	Offset      int      `json:"offset"`
}
