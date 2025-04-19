package seeders

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
	"errors"
)

func SeedTeams(db *config.DB, env *config.Env) error {
	var count int64
	if err := db.Model(&models.Team{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var user models.User
	if err := db.First(&user).Error; err != nil {
		return errors.New("no users found in the database")
	}

	var season models.Season
	if err := db.First(&season).Error; err != nil {
		return errors.New("no seasons found in the database")
	}

	teams := []models.Team{
		{UserId: user.Id, Name: "BIG CLAN", SeasonId: season.Id, Elo: 1123},
		{UserId: user.Id, Name: "Astralis", SeasonId: season.Id, Elo: 1245},
		{UserId: user.Id, Name: "Natus Vincere", SeasonId: season.Id, Elo: 182},
		{UserId: user.Id, Name: "G2 Esports", SeasonId: season.Id, Elo: 945},
		{UserId: user.Id, Name: "Team Liquid", SeasonId: season.Id, Elo: 885},
		{UserId: user.Id, Name: "FaZe Clan", SeasonId: season.Id, Elo: 812},
		{UserId: user.Id, Name: "Fnatic", SeasonId: season.Id, Elo: 789},
	}

	for _, team := range teams {
		if err := db.Create(&team).Error; err != nil {
			return err
		}
	}

	return nil
}
