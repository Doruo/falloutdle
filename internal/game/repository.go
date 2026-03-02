package game

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewGameRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// /----- CREATE -----/

// Add creates a new game record in the database.
func (r *Repository) Add(game *Game) error {

	result := r.db.Create(game)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// /----- READ -----/

// GetAll retrieves all games with optional pagination.
func (r *Repository) GetAll(limit, offset int) ([]Game, error) {

	var games []Game
	query := r.db
	if limit > 0 {
		query = query.Limit(limit)
	}
	if offset > 0 {
		query = query.Offset(offset)
	}

	result := query.Find(&games)
	if result.Error != nil {
		return nil, result.Error
	}

	return games, nil
}

// GetByID retrieves a game by its ID.
func (r *Repository) GetByID(id uint) (*Game, error) {

	var game Game
	result := r.db.First(&game, id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("game not found")
		}
		return nil, result.Error
	}

	return &game, nil
}

// GetByDate retrieves a game by its date.
func (r *Repository) GetByDate(date time.Time) (*Game, error) {

	var game Game
	result := r.db.First(&game, date)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("game not found")
		}
		return nil, result.Error
	}

	return &game, nil
}

// /----- UPDATE -----/

// Update modifies an existing game.
func (r *Repository) Update(game *Game) error {

	if game == nil {
		return errors.New("game cannot be nil")
	}

	result := r.db.Save(game)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("game not found")
	}

	return nil
}

// /----- DELETE -----/

// Delete removes a game by ID.
func (r *Repository) DeleteByID(id uint) error {
	if id == 0 {
		return errors.New("invalid game ID")
	}

	result := r.db.Delete(&Game{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("game not found")
	}

	return nil
}

func (r *Repository) DeleteByDate(date time.Time) error {

	result := r.db.Delete(&Game{}, date)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("game not found")
	}

	return nil
}
