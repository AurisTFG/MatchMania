package services

import (
	"MatchManiaAPI/initializers"
	"MatchManiaAPI/models"
)

func GetTeamByID(seasonID string, teamID string) (*models.Team, error) {
	var team models.Team

	result := initializers.DB.Where("season_id = ? AND id = ?", seasonID, teamID).First(&team)

	return &team, result.Error
}

func GetAllTeams(seasonID string) ([]models.Team, error) {
	var teams []models.Team

	result := initializers.DB.Where("season_id = ?", seasonID).Find(&teams)

	return teams, result.Error
}

func CreateTeam(teamDto *models.CreateTeamDto, seasonID uint) (*models.Team, error) {
	newTeam := teamDto.ToTeam()
	newTeam.SeasonID = seasonID
	newTeam.Elo = 1000

	result := initializers.DB.Create(&newTeam)

	return &newTeam, result.Error
}

func UpdateTeam(team *models.Team, teamDto *models.UpdateTeamDto) (*models.Team, error) {
	updatedTeam := teamDto.ToTeam()

	result := initializers.DB.Model(team).Updates(updatedTeam)

	return team, result.Error
}

func DeleteTeam(team *models.Team) error {
	result := initializers.DB.Delete(&team)

	return result.Error
}
