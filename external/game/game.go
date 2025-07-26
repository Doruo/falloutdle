package game

import (
	"fmt"
	"time"

	"github.com/doruo/falloutdle/internal/character"
	"github.com/doruo/falloutdle/internal/database"
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

// Game logic service
type GameService struct {
	characterService character.Service
	currentGame      *Game
}

func NewGameService() *GameService {

	db := database.GetInstance()
	repo := character.NewCharacterRepository(db)

	return &GameService{
		characterService: *character.NewCharacterService(repo),
		currentGame:      nil,
	}
}

var instance *GameService

func GetServiceInstance() *GameService {
	if instance == nil {
		instance = NewGameService()
	}
	return instance
}

// NewCurrentGame creates a new game for today from a RandomCharacter
func (gs *GameService) NewCurrentGame() (*Game, error) {

	// Retrieves random character from database
	character, error := gs.GetRandomValidCharacter()

	// Retrieves another character if not valid
	for !gs.characterService.IsValidForGame(character) {

		if error != nil {
			return nil, error
		}

		character, error = gs.GetRandomValidCharacter()
	}

	// Marks character or update played date
	gs.characterService.UpdateAsPlayed(character.ID)

	return NewGame(*character), nil
}

// /----- GET LOGIC FUNCTIONS -----/

func (gs *GameService) GetRandomCharacter() (*character.Character, error) {

	// Retrieves random character from database
	character, error := gs.characterService.GetRandomCharacter()

	if error != nil {
		return nil, error
	}

	return character, nil
}

func (gs *GameService) GetRandomValidCharacter() (*character.Character, error) {

	// Retrieves random character from database
	character, error := gs.GetRandomCharacter()

	// Retrieves another character if not valid
	for !gs.characterService.IsValidForGame(character) {

		if error != nil {
			return nil, error
		}
		character, error = gs.GetRandomCharacter()
	}

	return character, nil
}

// /----- GET FUNCTIONS -----/

// GetCurrentCharacter returns today current character.
// Creates a new one if none found
func (gs *GameService) GetCurrentCharacter() (*character.Character, error) {

	game, err := gs.GetCurrentGame()

	if err != nil {
		return nil, err
	}

	fmt.Println(game)

	return &game.CurrentCharacter, nil
}

// GetCurrentGame returns today current game.
// Creates a new one for if none found
func (gs *GameService) GetCurrentGame() (*Game, error) {

	// Creates a new one for today if none found
	if gs.currentGame == nil {

		fmt.Println("LOG: no game found for today, creating new one...")

		var err error
		gs.currentGame, err = gs.NewCurrentGame()

		if err != nil {
			return nil, err
		}
	}

	return gs.currentGame, nil
}

// /----- UTILITY FUNCTIONS -----/

// getCurrentDate returns today's date in 24h UTC format.
func getCurrentDate() time.Time {
	return time.Now().UTC().Truncate(24 * time.Hour)
}
