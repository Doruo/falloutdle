package character

import (
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewCharacterRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// /----- CREATE -----/

// Add creates a new character record in the database
func (r *Repository) Add(character *Character) error {

	result := r.db.Create(character)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// WARNING: VERY EXPENSIVE FUNCTION FOR WIKI API AND DATABASE,
// DO NOT USE IT WITHOUT CAUTION !
// AddAllCharactersFromWiki creates all new characters record in the database from Wiki
/*
func AddAllCharactersFromWiki() {

	db := database.NewDatabaseConnection()
	repo := NewCharacterRepository(db)
	client := wiki.NewWikiClient()

	chars, err := client.FetchAllCharacters()

	if err != nil {
		fmt.Print("Error during fetch: %t", err)
	}

	for _, character := range chars {
		repo.Add(character)
	}
}*/

// /----- SELECT -----/

// GetAll retrieves all characters with optional pagination
func (r *Repository) GetAll(limit, offset int) ([]Character, error) {
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
func (r *Repository) GetByID(id uint) (*Character, error) {

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
func (r *Repository) GetByWikiTitle(wikiTitle string) (*Character, error) {
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
func (r *Repository) GetByName(name string) (*Character, error) {

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

// /----- UPDATE -----/

// Update modifies an existing character
func (r *Repository) Update(character *Character) error {

	if character == nil {
		return errors.New("character cannot be nil")
	}

	if character.ID == 0 {
		return errors.New("character ID is required for update")
	}

	result := r.db.Save(character)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("character not found")
	}

	return nil
}

// /----- DELETE -----/

// Delete removes a character by ID
func (r *Repository) DeleteByID(id uint) error {
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
func (r *Repository) DeleteByWikiTitle(wikiTitle string) error {
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
