package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"

	"github.com/google/uuid"
)

type ResultRepository interface {
	FindAll() ([]models.Result, error)
	FindAllBySeasonID(uuid.UUID) ([]models.Result, error)
	FindAllByTeamID(uuid.UUID) ([]models.Result, error)
	FindAllBySeasonIDAndTeamID(uuid.UUID, uuid.UUID) ([]models.Result, error)
	FindById(uuid.UUID) (*models.Result, error)
	FindByIdAndSeasonIDAndTeamID(seasonId uuid.UUID, teamId uuid.UUID, resultId uuid.UUID) (*models.Result, error)
	Create(*models.Result) (*models.Result, error)
	Update(*models.Result, *models.Result) (*models.Result, error)
	Delete(*models.Result) error
}

type resultRepository struct {
	db *config.DB
}

func NewResultRepository(db *config.DB) ResultRepository {
	return &resultRepository{db: db}
}

func (r *resultRepository) FindAll() ([]models.Result, error) {
	var results []models.Result

	result := r.db.Find(&results)

	return results, result.Error
}

func (r *resultRepository) FindAllBySeasonID(seasonId uuid.UUID) ([]models.Result, error) {
	var results []models.Result

	result := r.db.Where("season_id = ?", seasonId).Find(&results)

	return results, result.Error
}

func (r *resultRepository) FindAllByTeamID(teamId uuid.UUID) ([]models.Result, error) {
	var results []models.Result

	result := r.db.Where("team_id = ? OR opponent_team_id = ?", teamId, teamId).Find(&results)

	return results, result.Error
}

func (r *resultRepository) FindAllBySeasonIDAndTeamID(seasonId uuid.UUID, teamId uuid.UUID) ([]models.Result, error) {
	var results []models.Result

	result := r.db.Where(
		"season_id = ? AND (team_id = ? OR opponent_team_id = ?)",
		seasonId, teamId, teamId,
	).Find(&results)

	return results, result.Error
}

func (r *resultRepository) FindById(resultId uuid.UUID) (*models.Result, error) {
	var resultModel models.Result

	result := r.db.First(&resultModel, resultId)

	return &resultModel, result.Error
}

func (r *resultRepository) FindByIdAndSeasonIDAndTeamID(
	seasonId uuid.UUID,
	teamId uuid.UUID,
	resultId uuid.UUID,
) (*models.Result, error) {
	var resultModel models.Result

	result := r.db.Where(
		"season_id = ? AND (team_id = ? OR opponent_team_id = ?) AND id = ?",
		seasonId, teamId, teamId, resultId,
	).First(&resultModel)

	return &resultModel, result.Error
}

func (r *resultRepository) Create(newResult *models.Result) (*models.Result, error) {
	result := r.db.Create(newResult)

	return newResult, result.Error
}

func (r *resultRepository) Update(currentResult *models.Result, updatedResult *models.Result) (*models.Result, error) {
	result := r.db.Model(currentResult).Updates(updatedResult)

	return currentResult, result.Error
}

func (r *resultRepository) Delete(resultModel *models.Result) error {
	result := r.db.Delete(resultModel)

	return result.Error
}
