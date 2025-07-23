package character

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// characterService implements Repository using CharacterRepository
type CharacterService struct {
	repo *CharacterRepository
}

// NewCharacterService creates a new character service
func NewCharacterService(repo *CharacterRepository) *CharacterService {
	return &CharacterService{repo: repo}
}

// /----- GET FUNCTIONS -----/

// GetByID retrieves a character by ID
func (s *CharacterService) GetByID(id int) (*Character, error) {

	if id <= 0 {
		return nil, errors.New("invalid ID")
	}

	char, err := s.repo.GetByID(uint(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get character ID %d: %w", id, err)
	}

	return char, nil
}

// GetAllCharacters retrieves all valid characters for the game
func (s *CharacterService) GetAllCharacters() ([]Character, error) {

	characters, err := s.repo.GetAll(0, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to get characters: %w", err)
	}

	// Filter valid characters for the game
	var gameCharacters []Character
	for _, char := range characters {
		if s.isValidForGame(&char) {
			gameCharacters = append(gameCharacters, char)
		}
	}

	return gameCharacters, nil
}

// GetRandomCharacter selects a random character
func (s *CharacterService) GetRandomCharacter() (*Character, error) {

	characters, err := s.GetAllCharacters()
	if err != nil {
		return nil, fmt.Errorf("failed to get characters: %w", err)
	}

	if len(characters) == 0 {
		return nil, errors.New("no characters available")
	}

	randomIndex := newRandom().Intn(len(characters))

	return &characters[randomIndex], nil
}

// /----- UTILITY FUNCTIONS -----/

// MarkAsPlayed marks a character as played
func (s *CharacterService) MarkAsPlayed(characterID int) error {

	if characterID <= 0 {
		return errors.New("invalid character ID")
	}

	char, err := s.repo.GetByID(uint(characterID))
	if err != nil {
		return fmt.Errorf("character not found: %w", err)
	}

	char.UpdateAsPlayed()

	err = s.repo.Update(char)
	if err != nil {
		return fmt.Errorf("failed to update character: %w", err)
	}

	return nil
}

// isValidForGame checks if a character is valid for the game
func (s *CharacterService) isValidForGame(char *Character) bool {

	if char.Name == "" || char.Race == "" || char.Gender == "" {
		return false
	}

	if len(char.Games) == 0 && char.MainGame == "" {
		return false
	}

	// Exclude playable characters (they might be too obvious)
	if char.IsPlayable() {
		return false
	}

	return true
}

// newRandom returns a new random using seed based on today date
func newRandom() *rand.Rand {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	return rand.New(source)
}
