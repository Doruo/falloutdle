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
