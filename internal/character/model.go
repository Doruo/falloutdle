package character

import (
	"fmt"
	"strings"
	"time"
)

var mainCharacters = []string{
	"Vault Dweller",
	"Chosen One",
	"Lone Wanderer",
	"Courier",
	"Sole Survivor",
	"Resident",
}

// Character represents a Fallout character with all relevant game information
type Character struct {
	ID          uint     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string   `json:"name" gorm:"size:255;index"`
	WikiTitle   string   `json:"wiki_title" gorm:"size:500;uniqueIndex"` // Original wiki page title
	Games       []string `json:"games" gorm:"serializer:json"`           // Main games where character appears
	Mentions    []string `json:"mentions" gorm:"serializer:json"`        // Games where character is mentioned
	Race        string   `json:"race" gorm:"size:100"`
	Gender      string   `json:"gender" gorm:"size:20"`
	Status      string   `json:"status" gorm:"size:50"` // Alive, Deceased, Unknown
	Affiliation []string `json:"affiliation" gorm:"serializer:json"`
	Role        string   `json:"role" gorm:"size:255"`
	Titles      []string `json:"titles" gorm:"serializer:json"`
	MainGame    string   `json:"main_game" gorm:"size:100;index"` // Primary game of origin
	ImageURL    string   `json:"image_url" gorm:"type:text"`

	PlayedAt *time.Time `json:"played_at,omitempty" gorm:"index"` // last played date
}

// NewCharacter creates a new Character instance
func NewCharacter(name, wikiTitle string) *Character {
	return &Character{
		Name:        name,
		WikiTitle:   wikiTitle,
		Games:       make([]string, 0),
		Mentions:    make([]string, 0),
		Affiliation: make([]string, 0),
		Titles:      make([]string, 0),
	}
}

// /----- GETTER FUNCTIONS -----/

// GetMainGame returns the primary game of origin (first in games list)
func (c *Character) GetMainGame() string {
	if len(c.Games) > 0 {
		return c.Games[0]
	}
	return ""
}

// IsMainCharacter determines if character is playable (like Vault Dweller)
func (c *Character) IsMainCharacter() bool {
	for _, playable := range mainCharacters {
		if strings.Contains(c.Name, playable) {
			return true
		}
	}
	return false
}

func (c *Character) IsPlayed() bool {
	return c.PlayedAt != nil
}

// /----- SETTER FUNCTIONS -----/

func (c *Character) UpdateAsPlayed() *Character {
	now := time.Now()
	c.PlayedAt = &now
	return c
}

// /----- STRING FUNCTIONS -----/

// String returns a compact version for logs or short displays.
// This method provides a readable text representation of the character
func (c *Character) String() string {

	if c == nil {
		return "<nil>"
	}

	var parts []string
	parts = append(parts, c.Name)
	parts = append(parts, fmt.Sprintf("ID:%d", c.ID))
	parts = append(parts, c.Race)
	parts = append(parts, c.MainGame)

	if len(parts) == 0 {
		return "Character{}"
	}
	return fmt.Sprintf("Character{%s}", strings.Join(parts, ", "))
}
