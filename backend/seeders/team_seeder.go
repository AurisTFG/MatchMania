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

	var leagues []models.League
	if err := db.First(&leagues).Error; err != nil {
		return errors.New("no leagues found in the database")
	}

	teams := []models.Team{
		{UserId: user.Id, Name: "BIG CLAN", Leagues: leagues, Elo: 1123},
		{UserId: user.Id, Name: "Astralis", Leagues: leagues, Elo: 1245},
		{UserId: user.Id, Name: "Natus Vincere", Leagues: leagues, Elo: 182},
		{UserId: user.Id, Name: "G2 Esports", Leagues: leagues, Elo: 945},
		{UserId: user.Id, Name: "Team Liquid", Leagues: leagues, Elo: 885},
		{UserId: user.Id, Name: "FaZe Clan", Leagues: leagues, Elo: 812},
		{UserId: user.Id, Name: "Fnatic", Leagues: leagues, Elo: 789},
	}

	if err := db.Create(&teams).Error; err != nil {
		return err
	}

	return nil
}
