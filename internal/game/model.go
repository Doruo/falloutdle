package game

import (
	"time"
)

// Game represents a game state.
type Game struct {
	Date        time.Time `json:"date" gorm:"primaryKey;"`
	CharacterID uint      `json:"id" gorm:"characterID;"`
}

func NewGame(id uint) *Game {
	return &Game{
		Date:        time.Now(),
		CharacterID: id,
	}
}
