package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
)

type SeasonRepository interface {
	FindAll() ([]models.Season, error)
	FindByID(uint) (*models.Season, error)
	Create(*models.Season) (*models.Season, error)
	Update(uint, *models.Season) (*models.Season, error)
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

	result := r.db.Find(&seasons)

	return seasons, result.Error
}

func (r *seasonRepository) FindByID(seasonID uint) (*models.Season, error) {
	var season models.Season

	result := r.db.First(&season, seasonID)

	return &season, result.Error
}

func (r *seasonRepository) Create(season *models.Season) (*models.Season, error) {
	result := r.db.Create(season)

	return season, result.Error
}

func (r *seasonRepository) Update(seasonID uint, season *models.Season) (*models.Season, error) {
	result := r.db.Model(&models.Season{}).Where("id = ?", seasonID).Updates(season)

	return season, result.Error
}

func (r *seasonRepository) Delete(season *models.Season) error {
	result := r.db.Delete(season)

	return result.Error
}
