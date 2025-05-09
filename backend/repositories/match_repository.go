package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"

	"github.com/google/uuid"
)

type MatchRepository interface {
	GetAll() ([]*models.Match, error)
	GetByID(matchId uuid.UUID) (*models.Match, error)
	Create(match *models.Match) error
	Save(match *models.Match) error
	Delete(id uuid.UUID) error
}

type matchRepository struct {
	db *config.DB
}

func NewMatchRepository(db *config.DB) MatchRepository {
	return &matchRepository{db: db}
}

func (r *matchRepository) GetAll() ([]*models.Match, error) {
	var matches []*models.Match

	result := r.db.
		Joins("League").
		Preload("Teams.Players").
		Order("created_at DESC").
		Find(&matches)

	return matches, result.Error
}

func (r *matchRepository) GetByID(id uuid.UUID) (*models.Match, error) {
	var match models.Match

	result := r.db.
		Joins("League").
		Preload("Teams.Players").
		First(&match, "id = ?", id)

	return &match, result.Error
}

func (r *matchRepository) Create(match *models.Match) error {
	result := r.db.Create(match)

	return result.Error
}

func (r *matchRepository) Save(match *models.Match) error {
	result := r.db.Save(match)

	return result.Error
}

func (r *matchRepository) Delete(id uuid.UUID) error {
	result := r.db.Delete(&models.Match{}, id)

	return result.Error
}
