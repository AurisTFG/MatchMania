package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
	"fmt"

	"github.com/google/uuid"
)

type SeasonRepository interface {
	FindAll() ([]models.Season, error)
	FindById(uuid.UUID) (*models.Season, error)
	Create(*models.Season) (*models.Season, error)
	Update(*models.Season, *models.Season) (*models.Season, error)
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

	result := r.db.Preload("User").Find(&seasons)

	fmt.Println("Seasons found:", seasons)

	return seasons, result.Error
}

func (r *seasonRepository) FindById(seasonId uuid.UUID) (*models.Season, error) {
	var season models.Season

	result := r.db.First(&season, seasonId)

	return &season, result.Error
}

func (r *seasonRepository) Create(season *models.Season) (*models.Season, error) {
	result := r.db.Create(season)

	return season, result.Error
}

func (r *seasonRepository) Update(currentSeason *models.Season, updatedSeason *models.Season) (*models.Season, error) {
	result := r.db.Model(currentSeason).Updates(updatedSeason)

	return currentSeason, result.Error
}

func (r *seasonRepository) Delete(season *models.Season) error {
	result := r.db.Delete(season)

	return result.Error
}
