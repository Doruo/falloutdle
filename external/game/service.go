package game

import (
	"github.com/doruo/falloutdle/internal/character"
)

// gameService implémentation concrète
type gameService struct {
	characterService character.Service
	currentGame      *Game
}

// /----- GET FUNCTIONS -----/

// GetTodayGame
func (gs *gameService) GetTodayGame() (*Game, error) {

	today := getTodayDate()
	// Already existing game today created
	if gs.currentGame != nil && gs.currentGame.Date.Equal(today) {
		return gs.currentGame, nil
	}
	return nil, nil
}

// GetGameState retourne l'état actuel du jeu
func (gs *gameService) GetGameState() *Game {
	return gs.currentGame
}
