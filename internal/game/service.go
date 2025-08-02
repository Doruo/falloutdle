package game

import (
	"fmt"
	"strings"
	"time"

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
func (s *Service) NewCurrentGame() (*Game, error) {

	// Retrieves random character from database
	character, error := s.getCharacterValidRandom()

	// Retrieves another character if not valid
	for !s.characterService.IsValidForGame(character) {

		if error != nil {
			return nil, error
		}

		character, error = s.getCharacterValidRandom()
	}

	// Marks character or update played date
	s.characterService.UpdateCharacterAsPlayed(character.ID)

	// Create and save game into database
	game := NewGame(character.ID)
	s.Add(game)

	return game, nil
}

// /----- CREATE LOGIC FUNCTIONS -----/

// Add a game into database, returns nil if no error
func (s *Service) Add(g *Game) error {
	if error := s.repository.Add(g); error != nil {
		return error
	}
	return nil
}

// /----- GET LOGIC FUNCTIONS -----/

func (s *Service) GetGames() ([]Game, error) {

	games, error := s.repository.GetAll(0, 0)

	if error != nil {
		return nil, fmt.Errorf("failed to get characters: %w", error)
	}

	return games, nil
}

func (s *Service) GetGameByID(id uint) (*Game, error) {

	if id <= 0 {
		return nil, fmt.Errorf("invalid ID")
	}

	game, error := s.repository.GetByID(id)
	if error != nil {
		return nil, fmt.Errorf("failed to get game from ID %d: %w", id, error)
	}

	return game, nil
}

func (s *Service) GetGameByDate(date time.Time) (*Game, error) {

	game, error := s.repository.GetByDate(date)
	if error != nil {
		return nil, fmt.Errorf("failed to get game from date %d: %w", date, error)
	}

	return game, nil
}

func (s *Service) getCharacterValidRandom() (*character.Character, error) {

	// Retrieves random character from database
	character, error := s.characterService.GetCharacterRandom()

	// Retrieves another character if not valid
	for !s.characterService.IsValidForGame(character) {

		if error != nil {
			return nil, error
		}
		character, error = s.characterService.GetCharacterRandom()
	}

	return character, nil
}

// GetCurrentCharacter returns today current character.
// Creates a new one if none found
func (s *Service) GetCurrentCharacter() (*character.Character, error) {

	game, error := s.GetGameCurrent()

	if error != nil {
		return nil, error
	}
	character, error := s.characterService.GetCharacterByID(game.CharacterId)

	if error != nil {
		return nil, error
	}

	return character, nil
}

// GetGameCurrent returns today current game.
// Creates a new one for if none found
func (s *Service) GetGameCurrent() (*Game, error) {

	// Creates a new one for today if none found
	if s.currentGame == nil {

		fmt.Println("LOG: no game found for today, creating new one...")

		var error error
		s.currentGame, error = s.NewCurrentGame()

		if error != nil {
			return nil, error
		}
	}

	return s.currentGame, nil
}

// /----- POST LOGIC FUNCTIONS -----/

func (s *Service) ProcessGuess(name string) (bool, error) {

	character, error := s.GetCurrentCharacter()

	if error != nil {
		return false, error
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
