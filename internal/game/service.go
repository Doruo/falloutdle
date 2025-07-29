package game

import (
	"fmt"
	"strings"

	"github.com/doruo/falloutdle/internal/character"
	"github.com/doruo/falloutdle/pkg/strutils"
)

// Game logic service
type Service struct {
	characterService character.Service
	repository       Repository
	currentGame      *Game
}

func NewGameService(cs *character.Service, repo *Repository) *Service {
	return &Service{
		characterService: *cs,
		repository:       *repo,
	}
}

// NewCurrentGame creates a new game for today from a RandomCharacter
func (gs *Service) NewCurrentGame() (*Game, error) {

	// Retrieves random character from database
	character, error := gs.getCharacterValidRandom()

	// Retrieves another character if not valid
	for !gs.characterService.IsValidForGame(character) {

		if error != nil {
			return nil, error
		}

		character, error = gs.getCharacterValidRandom()
	}

	// Marks character or update played date
	gs.characterService.UpdateCharacterAsPlayed(character.ID)

	// Create and save game into database
	game := NewGame(character.ID)
	gs.Add(game)

	return game, nil
}

// /----- CREATE LOGIC FUNCTIONS -----/

// Add a game into database, returns nil if no error
func (gs *Service) Add(g *Game) error {
	if err := gs.repository.Add(g); err != nil {
		return err
	}
	return nil
}

// /----- GET LOGIC FUNCTIONS -----/

func (gs *Service) GetGames() ([]Game, error) {

	games, error := gs.repository.GetAll(0, 0)

	if error != nil {
		return nil, error
	}

	return games, nil
}

func (gs *Service) getCharacterValidRandom() (*character.Character, error) {

	// Retrieves random character from database
	character, error := gs.characterService.GetCharacterRandom()

	// Retrieves another character if not valid
	for !gs.characterService.IsValidForGame(character) {

		if error != nil {
			return nil, error
		}
		character, error = gs.characterService.GetCharacterRandom()
	}

	return character, nil
}

// GetCurrentCharacter returns today current character.
// Creates a new one if none found
func (gs *Service) GetCurrentCharacter() (*character.Character, error) {

	game, err := gs.GetGameCurrent()

	if err != nil {
		return nil, err
	}

	character, err := gs.characterService.GetCharacterByID(game.CharacterID)

	if err != nil {
		return nil, err
	}

	return character, nil
}

// GetGameCurrent returns today current game.
// Creates a new one for if none found
func (gs *Service) GetGameCurrent() (*Game, error) {

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

// /----- POST LOGIC FUNCTIONS -----/

func (gs *Service) ProcessGuess(name string) (bool, error) {

	character, err := gs.GetCurrentCharacter()

	if err != nil {
		return false, err
	}

	return isGuessCorrect(name, character.Name), nil
}

// isGuessCorrect returns true if name guessed corresponds to the correct character name
func isGuessCorrect(guessName string, correctName string) bool {

	if guessName == correctName {
		return true
	}

	guess := strings.ToLower(strings.TrimSpace(guessName))
	correct := strings.ToLower(strings.TrimSpace(correctName))

	if guess == correct {
		return true
	}

	guess = strutils.NormalizeString(guessName)
	correct = strutils.NormalizeString(correctName)

	if len(guess) >= 4 && len(correct) >= 4 {
		return strings.Contains(correct, guess) || strings.Contains(guess, correct)
	}

	return false
}
