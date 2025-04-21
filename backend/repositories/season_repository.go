package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"

	"github.com/google/uuid"
)

type SeasonRepository interface {
	FindAll() ([]models.Season, error)
	FindById(uuid.UUID) (*models.Season, error)
	Create(*models.Season) error
	Update(*models.Season, *models.Season) error
	Delete(*models.Season) error
}

type seasonRepository struct {
	db *config.DB
}

func NewSeasonRepository(db *config.DB) SeasonRepository {
	return &seasonRepository{db: db}
}

func (r *seasonRepository) FindAll() ([]models.Season, error) {
	var seasons []models.Season

	result := r.db.Joins("User").Find(&seasons)

	return seasons, result.Error
}

func (r *seasonRepository) FindById(seasonId uuid.UUID) (*models.Season, error) {
	var season models.Season

	result := r.db.Joins("User").First(&season, seasonId)

	return &season, result.Error
}

func (r *seasonRepository) Create(season *models.Season) error {
	result := r.db.Create(season)

	return result.Error
}

func (r *seasonRepository) Update(currentSeason *models.Season, updatedSeason *models.Season) error {
	result := r.db.Model(currentSeason).Updates(updatedSeason)

	return result.Error
}

func (r *seasonRepository) Delete(season *models.Season) error {
	result := r.db.Delete(season)

	return result.Error
}
