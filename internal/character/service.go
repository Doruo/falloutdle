package character

import (
	"errors"
	"fmt"

	"github.com/doruo/falloutdle/pkg/random"
)

// characterService implements Repository using CharacterRepository
type Service struct {
	repo *Repository
}

// NewCharacterService creates a new character service
func NewCharacterService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// /----- GET FUNCTIONS -----/

// GetByID retrieves a character by ID
func (s *Service) GetByID(id int) (*Character, error) {

	if id <= 0 {
		return nil, errors.New("invalid ID")
	}

	char, err := s.repo.GetByID(uint(id))
	if err != nil {
		return nil, fmt.Errorf("failed to get character ID %d: %w", id, err)
	}

	return char, nil
}

func (s *Service) GetByWikiTitle(title string) (*Character, error) {

	if title == "" {
		return nil, errors.New("invalid title")
	}

	char, err := s.repo.GetByWikiTitle(title)
	if err != nil {
		return nil, fmt.Errorf("failed to get character from tite %s: %w", title, err)
	}

	return char, nil
}

// GetAllValidCharacters retrieves all valid characters for the game
func (s *Service) GetAllValidCharacters() ([]Character, error) {

	characters, err := s.repo.GetAll(0, 0)

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

// GetRandomCharacter selects a random character
func (s *Service) GetRandomCharacter() (*Character, error) {

	characters, err := s.GetAllValidCharacters()
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

// UpdateAsPlayed marks a character as played or updates his date if already played
func (s *Service) UpdateAsPlayed(characterID uint) error {

	if characterID <= 0 {
		return errors.New("invalid character ID")
	}

	char, err := s.repo.GetByID(characterID)
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

// UpdateAsUnplayed set a character as unplayed
func (s *Service) UpdateAsUnplayed(characterID uint) error {

	if characterID <= 0 {
		return errors.New("invalid character ID")
	}

	char, err := s.repo.GetByID(characterID)
	if err != nil {
		return fmt.Errorf("character not found: %w", err)
	}

	char.UpdateAsUnplayed()

	err = s.repo.Update(char)
	if err != nil {
		return fmt.Errorf("failed to update character: %w", err)
	}

	return nil
}

// isValidForGame checks if a character is valid for the game
func (s *Service) IsValidForGame(char *Character) bool {

	if char.Name == "" || char.Race == "" {
		return false
	}

	if len(char.Games) == 0 && char.MainGame == "" {
		return false
	}

	return !char.IsPlayed()
}
