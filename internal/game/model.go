package game

import (
	"time"
)

// Game represents a game state.
type Game struct {
	Id          uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Date        time.Time `json:"date" gorm:"date;"`
	CharacterID uint      `json:"character_id" gorm:"characterID,omitEmpty;"`
}

func NewGame(id uint) *Game {
	return &Game{
		Id:   id,
		Date: time.Now(),
	}
}
