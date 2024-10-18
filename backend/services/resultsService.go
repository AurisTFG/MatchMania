package services

import (
	"MatchManiaAPI/initializers"
	"MatchManiaAPI/models"
)

func GetResultByID(seasonID string, teamID string, resultID string) (*models.Result, error) {
	var resultModel models.Result

	result := initializers.DB.Where(
		"season_id = ? AND (team_id = ? OR opponent_team_id = ?) AND id = ?",
		seasonID, teamID, teamID, resultID,
	).First(&resultModel)

	return &resultModel, result.Error
}

func GetAllResults(seasonID string, teamID string) ([]models.Result, error) {
	var resultModels []models.Result

	result := initializers.DB.Where(
		"season_id = ? AND (team_id = ? OR opponent_team_id = ?)",
		seasonID, teamID, teamID,
	).Find(&resultModels)

	return resultModels, result.Error
}

func CreateResult(resultDto *models.CreateResultDto, seasonID uint, teamID uint) (*models.Result, error) {
	newResult := resultDto.ToResult()
	newResult.SeasonID = seasonID
	newResult.TeamID = teamID

	result := initializers.DB.Create(&newResult)

	return &newResult, result.Error
}

func UpdateResult(resultModel *models.Result, resultDto *models.UpdateResultDto) (*models.Result, error) {
	// updatedResult := resultDto.ToResult()

	updatedResult := map[string]interface{}{ // "hack" to make zero values go through :DDDDD
		"match_start_date": resultDto.MatchStartDate,
		"match_end_date":   resultDto.MatchEndDate,
		"score":            resultDto.Score,
		"opponent_score":   resultDto.OpponentScore,
	}

	result := initializers.DB.Model(resultModel).Updates(updatedResult)

	return resultModel, result.Error
}

func DeleteResult(resultModel *models.Result) error {
	result := initializers.DB.Delete(&resultModel)

	return result.Error
}
