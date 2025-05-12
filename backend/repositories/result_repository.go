package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"

	"github.com/google/uuid"
)

type ResultRepository interface {
	GetAll() ([]models.Result, error)
	GetById(uuid.UUID) (*models.Result, error)
	Create(*models.Result) error
	Update(*models.Result, *models.Result) error
	Delete(*models.Result) error
}

type resultRepository struct {
	db *config.DB
}

func NewResultRepository(db *config.DB) ResultRepository {
	return &resultRepository{db: db}
}

func (r *resultRepository) GetAll() ([]models.Result, error) {
	var results []models.Result

	result := r.db.
		Joins("User").
		Joins("Team").
		Joins("OpponentTeam").
		Joins("League").
		Order("start_date DESC").
		Find(&results)

	return results, result.Error
}

func (r *resultRepository) GetById(resultId uuid.UUID) (*models.Result, error) {
	var resultModel models.Result

	result := r.db.First(&resultModel, resultId)

	return &resultModel, result.Error
}

func (r *resultRepository) Create(newResult *models.Result) error {
	result := r.db.Create(newResult)

	return result.Error
}

func (r *resultRepository) Update(currentResult *models.Result, updatedResult *models.Result) error {
	result := r.db.
		Model(currentResult).
		Updates(updatedResult)

	return result.Error
}

func (r *resultRepository) Delete(resultModel *models.Result) error {
	result := r.db.Delete(resultModel)

	return result.Error
}
