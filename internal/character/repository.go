package character

import (
	"errors"

	"gorm.io/gorm"
)

type CharacterRepository struct {
	db *gorm.DB
}

func NewCharacterRepository(db *gorm.DB) *CharacterRepository {
	return &CharacterRepository{db: db}
}

// /----- CREATE -----/

// Add creates a new character record in the database
func (r *CharacterRepository) Add(character *Character) error {

	result := r.db.Create(character)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// /----- SELECT -----/

// GetAll retrieves all characters with optional pagination
func (r *CharacterRepository) GetAll(limit, offset int) ([]Character, error) {
	var characters []Character

	query := r.db
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	result := query.Find(&characters)
	if result.Error != nil {
		return nil, result.Error
	}

	return characters, nil
}

// GetByID retrieves a character by its ID
func (r *CharacterRepository) GetByID(id uint) (*Character, error) {

	var character Character
	result := r.db.First(&character, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("character not found")
		}
		return nil, result.Error
	}

	return &character, nil
}

// GetByWikiTitle retrieves a character by its WikiTitle
func (r *CharacterRepository) GetByWikiTitle(wikiTitle string) (*Character, error) {
	var character Character

	result := r.db.Where("wiki_title = ?", wikiTitle).First(&character)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("character not found")
		}
		return nil, result.Error
	}

	return &character, nil
}

// GetByName retrieves a character by its name
func (r *CharacterRepository) GetByName(name string) (*Character, error) {

	var character Character
	result := r.db.Where("name = ?", name).First(&character)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("character not found")
		}
		return nil, result.Error
	}

	return &character, nil
}

// /----- DELETE -----/

// Delete removes a character by ID
func (r *CharacterRepository) DeleteByID(id uint) error {
	if id == 0 {
		return errors.New("invalid character ID")
	}

	result := r.db.Delete(&Character{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("character not found")
	}

	return nil
}

// DeleteByWikiTitle removes a character by WikiTitle
func (r *CharacterRepository) DeleteByWikiTitle(wikiTitle string) error {
	if wikiTitle == "" {
		return errors.New("wiki title cannot be empty")
	}

	result := r.db.Where("wiki_title = ?", wikiTitle).Delete(&Character{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("character not found")
	}

	return nil
}
