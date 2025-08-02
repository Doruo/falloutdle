package game

import (
	"time"
)

// Game represents a game state.
type Game struct {
	Id          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Date        time.Time `json:"date" gorm:"date;"`
	CharacterId uint      `json:"character_id" gorm:"characterID,omitEmpty;"`
}

func NewGame(charId uint) *Game {
	return &Game{
		Date:        time.Now(),
		CharacterId: charId,
	}
}
