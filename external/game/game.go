package game

import (
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
	cs := character.NewCharacterService(repo)

	return &GameService{
		characterService: *cs,
		currentGame:      nil,
	}
}

// /----- LOGIC FUNCTIONS -----/

// CreateCurrentGame creates a new game for today from a RandomCharacter
func (gs *GameService) CreateCurrentGame() (*Game, error) {

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

// GetCurrentCharacter returns today current character.
// Creates a new one if none found
func (gs *GameService) GetCurrentCharacter() (*character.Character, error) {

	game, err := gs.GetCurrentGame()

	if err != nil {
		return nil, err
	}

	return &game.CurrentCharacter, nil
}

// GetCurrentGame returns today current game.
// Creates a new one if none found
func (gs *GameService) GetCurrentGame() (*Game, error) {

	today := getCurrentDate()

	//fmt.Println(gs.currentGame)
	//fmt.Println("meme date ?: ", gs.currentGame.Date.Equal(today))

	// Already existing game created today
	if gs.currentGame != nil && gs.currentGame.Date.Equal(today) {
		return gs.currentGame, nil
	}

	return gs.CreateCurrentGame()
}

// /----- UTILITY FUNCTIONS -----/

// getCurrentDate returns today's date in 24h UTC format.
func getCurrentDate() time.Time {
	return time.Now().UTC().Truncate(24 * time.Hour)
}
