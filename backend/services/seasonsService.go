package services

import (
	"MatchManiaAPI/initializers"
	"MatchManiaAPI/models"
)

func GetSeasonByID(seasonID string) (*models.Season, error) {
	var season models.Season

	result := initializers.DB.First(&season, seasonID)

	return &season, result.Error
}

func GetAllSeasons() ([]models.Season, error) {
	var seasons []models.Season

	result := initializers.DB.Find(&seasons)

	return seasons, result.Error
}

func CreateSeason(seasonDto *models.CreateSeasonDto) (*models.Season, error) {
	newSeason := seasonDto.ToSeason()

	result := initializers.DB.Create(&newSeason)

	return &newSeason, result.Error
}

func UpdateSeason(season *models.Season, seasonDto *models.UpdateSeasonDto) (*models.Season, error) {
	updatedSeason := seasonDto.ToSeason()

	result := initializers.DB.Model(season).Updates(updatedSeason)

	return season, result.Error
}

func DeleteSeason(season *models.Season) error {
	result := initializers.DB.Delete(&season)

	return result.Error
}
