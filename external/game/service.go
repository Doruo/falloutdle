package game

import (
	"github.com/doruo/falloutdle/internal/character"
)

// Game logic service
type GameService struct {
	characterService character.Service
	currentGame      *Game
}

func NewGameService(cs character.Service) *GameService {
	return &GameService{
		characterService: cs,
		currentGame:      nil,
	}
}

// /----- LOGIC FUNCTIONS -----/

// CreateTodayGame creates a new game for today from a RandomCharacter
func (gs *GameService) CreateTodayGame() (*Game, error) {

	// Retrieves random character from database
	character, error := gs.characterService.GetRandomCharacter()

	// Retrieves another character if not valid
	for !gs.characterService.IsValidForGame(character) {

		if error != nil {
			return nil, error
		}
		character, error = gs.characterService.GetRandomCharacter()
	}

	// Marks character or update played date
	gs.characterService.UpdateAsPlayed(character.ID)
	return NewGame(*character), nil
}

// /----- GET FUNCTIONS -----/

// GetTodayGame
func (gs *GameService) GetTodayGame() (*Game, error) {

	today := getTodayDate()
	// Already existing game today created
	if gs.currentGame != nil && gs.currentGame.Date.Equal(today) {
		return gs.currentGame, nil
	}
	return gs.CreateTodayGame()
}
