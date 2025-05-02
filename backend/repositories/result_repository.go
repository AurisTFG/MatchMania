package repositories

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"

	"github.com/google/uuid"
)

type ResultRepository interface {
	FindAll() ([]models.Result, error)
	FindAllByLeagueID(uuid.UUID) ([]models.Result, error)
	FindAllByTeamID(uuid.UUID) ([]models.Result, error)
	FindAllByLeagueIDAndTeamID(uuid.UUID, uuid.UUID) ([]models.Result, error)
	FindById(uuid.UUID) (*models.Result, error)
	FindByIdAndLeagueIDAndTeamID(leagueId uuid.UUID, teamId uuid.UUID, resultId uuid.UUID) (*models.Result, error)
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

func (r *resultRepository) FindAll() ([]models.Result, error) {
	var results []models.Result

	result := r.db.Joins("User").Joins("Team").Joins("OpponentTeam").Find(&results)

	return results, result.Error
}

func (r *resultRepository) FindAllByLeagueID(leagueId uuid.UUID) ([]models.Result, error) {
	var results []models.Result

	result := r.db.Joins("User").Joins("Team").Joins("OponnentTeam").Where("league_id = ?", leagueId).Find(&results)

	return results, result.Error
}

func (r *resultRepository) FindAllByTeamID(teamId uuid.UUID) ([]models.Result, error) {
	var results []models.Result

	result := r.db.Joins("User").
		Joins("Team").
		Joins("OponnentTeam").
		Where("team_id = ? OR opponent_team_id = ?", teamId, teamId).
		Find(&results)

	return results, result.Error
}

func (r *resultRepository) FindAllByLeagueIDAndTeamID(leagueId uuid.UUID, teamId uuid.UUID) ([]models.Result, error) {
	var results []models.Result

	result := r.db.Joins("User").Joins("Team").Joins("OpponentTeam").Where(
		"\"results\".\"league_id\" = ? AND (team_id = ? OR opponent_team_id = ?)",
		leagueId, teamId, teamId,
	).Find(&results)

	return results, result.Error
}

func (r *resultRepository) FindById(resultId uuid.UUID) (*models.Result, error) {
	var resultModel models.Result

	result := r.db.Joins("User").Joins("Team").Joins("OponnentTeam").First(&resultModel, resultId)

	return &resultModel, result.Error
}

func (r *resultRepository) FindByIdAndLeagueIDAndTeamID(
	leagueId uuid.UUID,
	teamId uuid.UUID,
	resultId uuid.UUID,
) (*models.Result, error) {
	var resultModel models.Result

	result := r.db.Joins("User").Joins("Team").Joins("OponnentTeam").Where(
		"league_id = ? AND (team_id = ? OR opponent_team_id = ?) AND id = ?",
		leagueId, teamId, teamId, resultId,
	).First(&resultModel)

	return &resultModel, result.Error
}

func (r *resultRepository) Create(newResult *models.Result) error {
	result := r.db.Create(newResult)

	return result.Error
}

func (r *resultRepository) Update(currentResult *models.Result, updatedResult *models.Result) error {
	result := r.db.Model(currentResult).Updates(updatedResult)

	return result.Error
}

func (r *resultRepository) Delete(resultModel *models.Result) error {
	result := r.db.Delete(resultModel)

	return result.Error
}
