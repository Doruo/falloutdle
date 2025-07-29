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

	characters, err := s.repository.GetAll(0, 0)

	if err != nil {
		return nil, fmt.Errorf("failed to get characters: %w", err)
	}

	return characters, nil
}

// GetCharactersValid retrieves all valid characters for the game
func (s *Service) GetCharactersValid() ([]Character, error) {

	characters, err := s.GetCharacters()

	if err != nil {
		return nil, fmt.Errorf("failed to get characters: %w", err)
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

	char, err := s.repository.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get character ID %d: %w", id, err)
	}

	return char, nil
}

func (s *Service) GetByWikiTitle(title string) (*Character, error) {

	if title == "" {
		return nil, errors.New("invalid title")
	}

	char, err := s.repository.GetByWikiTitle(title)
	if err != nil {
		return nil, fmt.Errorf("failed to get character from tite %s: %w", title, err)
	}

	return char, nil
}

// GetCharacterRandom selects a random character
func (s *Service) GetCharacterRandom() (*Character, error) {

	characters, err := s.GetCharactersValid()
	if err != nil {
		return nil, fmt.Errorf("failed to get characters: %w", err)
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

	char, err := s.repository.GetByID(id)
	if err != nil {
		return fmt.Errorf("character not found: %w", err)
	}

	char.UpdateAsPlayed()

	err = s.repository.Update(char)
	if err != nil {
		return fmt.Errorf("failed to update character: %w", err)
	}

	return nil
}

// UpdateCharacterAsUnplayed set a character as unplayed
func (s *Service) UpdateCharacterAsUnplayed(id uint) error {

	if id <= 0 {
		return errors.New("invalid character ID")
	}

	char, err := s.repository.GetByID(id)
	if err != nil {
		return fmt.Errorf("character not found: %w", err)
	}

	char.UpdateAsUnplayed()

	err = s.repository.Update(char)
	if err != nil {
		return fmt.Errorf("failed to update character: %w", err)
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
