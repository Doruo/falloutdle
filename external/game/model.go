package game

import (
	"time"

	"github.com/doruo/falloutdle/internal/character"
)

// Game represents a current game state
type Game struct {
	CurrentCharacter character.Character
	Date             time.Time `json:"date"`
}

func NewGame(c character.Character) *Game {
	return &Game{
		CurrentCharacter: c,
		Date:             time.Now(),
	}
}

// /----- UTILITY FUNCTIONS -----/

// getTodayDate returns today's date in 24h UTC format.
func getTodayDate() time.Time {
	return time.Now().UTC().Truncate(24 * time.Hour)
}
