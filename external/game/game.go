package game

import (
	"time"

	"github.com/doruo/falloutdle/internal/character"
)

// GameState represents a current game state
type GameState struct {
	Date time.Time `json:"date"`
}

// gameService implémentation concrète
type gameService struct {
	characterService character.Service
	currentGame      *GameState
}

// GetTodayGame récupère ou initialise la partie du jour
func (gs *gameService) GetTodayGame() (*GameState, error) {

}

// GetGameState retourne l'état actuel du jeu
func (gs *gameService) GetGameState() *GameState {
	return gs.currentGame
}
