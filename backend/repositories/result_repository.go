package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
)

type ResultRepository interface {
	FindAll() ([]models.Result, error)
	FindAllBySeasonID(uint) ([]models.Result, error)
	FindAllByTeamID(uint) ([]models.Result, error)
	FindAllBySeasonIDAndTeamID(uint, uint) ([]models.Result, error)
	FindByID(uint) (*models.Result, error)
	FindByIDAndSeasonIDAndTeamID(uint, uint, uint) (*models.Result, error)
	Create(*models.Result) (*models.Result, error)
	Update(*models.Result) (*models.Result, error)
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

func (r *resultRepository) FindAllBySeasonID(seasonID uint) ([]models.Result, error) {
	var results []models.Result

	result := r.db.Where("season_id = ?", seasonID).Find(&results)

	return results, result.Error
}

func (r *resultRepository) FindAllByTeamID(teamID uint) ([]models.Result, error) {
	var results []models.Result

	result := r.db.Where("team_id = ? OR opponent_team_id = ?", teamID, teamID).Find(&results)

	return results, result.Error
}

func (r *resultRepository) FindAllBySeasonIDAndTeamID(seasonID uint, teamID uint) ([]models.Result, error) {

	var results []models.Result

	result := r.db.Where(
		"season_id = ? AND (team_id = ? OR opponent_team_id = ?)",
		seasonID, teamID, teamID,
	).Find(&results)

	return results, result.Error
}

func (r *resultRepository) FindByID(resultID uint) (*models.Result, error) {
	var resultModel models.Result

	result := r.db.First(&resultModel, resultID)

	return &resultModel, result.Error
}

func (r *resultRepository) FindByIDAndSeasonIDAndTeamID(seasonID uint, teamID uint, resultID uint) (*models.Result, error) {
	var resultModel models.Result

	result := r.db.Where(
		"season_id = ? AND (team_id = ? OR opponent_team_id = ?) AND id = ?",
		seasonID, teamID, teamID, resultID,
	).First(&resultModel)

	return &resultModel, result.Error
}

func (r *resultRepository) Create(resultModel *models.Result) (*models.Result, error) {
	result := r.db.Create(resultModel)

	return resultModel, result.Error
}

func (r *resultRepository) Update(resultModel *models.Result) (*models.Result, error) {
	result := r.db.Save(resultModel)

	return resultModel, result.Error
}

func (r *resultRepository) Delete(resultModel *models.Result) error {
	result := r.db.Delete(resultModel)

	return result.Error
}
