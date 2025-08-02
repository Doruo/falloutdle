package character

import (
	"errors"
	"fmt"

	"github.com/doruo/falloutdle/pkg/random"
)

// characterService implements Repository using CharacterRepository
type Service struct {
	repository *Repository
}

// NewCharacterService creates a new character service
func NewCharacterService(repo *Repository) *Service {
	return &Service{repository: repo}
}

// /----- GET FUNCTIONS -----/

func (s *Service) GetCharacters() ([]Character, error) {

	characters, error := s.repository.GetAll(0, 0)

	if error != nil {
		return nil, fmt.Errorf("failed to get characters: %w", error)
	}

	return characters, nil
}

// GetCharactersValid retrieves all valid characters for the game
func (s *Service) GetCharactersValid() ([]Character, error) {

	characters, error := s.GetCharacters()

	if error != nil {
		return nil, fmt.Errorf("failed to get characters: %w", error)
	}

	// Filter valid characters for the game
	var validCharacters []Character
	for _, char := range characters {
		if s.IsValidForGame(&char) {
			validCharacters = append(validCharacters, char)
		}
	}

	return validCharacters, nil
}

// GetCharacterByID retrieves a character by ID
func (s *Service) GetCharacterByID(id uint) (*Character, error) {

	if id <= 0 {
		return nil, errors.New("invalid ID")
	}

	char, error := s.repository.GetByID(id)
	if error != nil {
		return nil, fmt.Errorf("failed to get character from ID %d: %w", id, error)
	}

	return char, nil
}

func (s *Service) GetByWikiTitle(title string) (*Character, error) {

	if title == "" {
		return nil, errors.New("invalid title")
	}

	char, error := s.repository.GetByWikiTitle(title)
	if error != nil {
		return nil, fmt.Errorf("failed to get character from title %s: %w", title, error)
	}

	return char, nil
}

// GetCharacterRandom selects a random character
func (s *Service) GetCharacterRandom() (*Character, error) {

	characters, error := s.GetCharactersValid()
	if error != nil {
		return nil, fmt.Errorf("failed to get characters: %w", error)
	}

	if len(characters) == 0 {
		return nil, errors.New("no characters available")
	}

	randomIndex := random.NewRandom().Intn(len(characters))

	return &characters[randomIndex], nil
}

// /----- UTILITY FUNCTIONS -----/

// UpdateCharacterAsPlayed marks a character as played or updates his date if already played
func (s *Service) UpdateCharacterAsPlayed(id uint) error {

	if id <= 0 {
		return errors.New("invalid character ID")
	}

	char, error := s.repository.GetByID(id)
	if error != nil {
		return fmt.Errorf("character not found: %w", error)
	}

	char.UpdateAsPlayed()

	error = s.repository.Update(char)
	if error != nil {
		return fmt.Errorf("failed to update character: %w", error)
	}

	return nil
}

// UpdateCharacterAsUnplayed set a character as unplayed
func (s *Service) UpdateCharacterAsUnplayed(id uint) error {

	if id <= 0 {
		return errors.New("invalid character ID")
	}

	char, error := s.repository.GetByID(id)
	if error != nil {
		return fmt.Errorf("character not found: %w", error)
	}

	char.UpdateAsUnplayed()

	error = s.repository.Update(char)
	if error != nil {
		return fmt.Errorf("failed to update character: %w", error)
	}

	return nil
}

// isValidForGame checks if a character is valid for the game
func (s *Service) IsValidForGame(c *Character) bool {

	if c.Name == "" || c.Race == "" {
		return false
	}

	if len(c.Games) == 0 && c.MainGame == "" {
		return false
	}

	return !c.IsPlayed()
}
