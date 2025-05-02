package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"

	"github.com/google/uuid"
)

type LeagueRepository interface {
	FindAll() ([]models.League, error)
	FindById(uuid.UUID) (*models.League, error)
	Create(*models.League) error
	Update(*models.League, *models.League) error
	Delete(*models.League) error
}

type leagueRepository struct {
	db *config.DB
}

func NewLeagueRepository(db *config.DB) LeagueRepository {
	return &leagueRepository{db: db}
}

func (r *leagueRepository) FindAll() ([]models.League, error) {
	var leagues []models.League

	result := r.db.Joins("User").Find(&leagues)

	return leagues, result.Error
}

func (r *leagueRepository) FindById(leagueId uuid.UUID) (*models.League, error) {
	var league models.League

	result := r.db.Joins("User").First(&league, leagueId)

	return &league, result.Error
}

func (r *leagueRepository) Create(league *models.League) error {
	result := r.db.Create(league)

	return result.Error
}

func (r *leagueRepository) Update(currentLeague *models.League, updatedLeague *models.League) error {
	result := r.db.Model(currentLeague).Updates(updatedLeague)

	return result.Error
}

func (r *leagueRepository) Delete(league *models.League) error {
	result := r.db.Delete(league)

	return result.Error
}
