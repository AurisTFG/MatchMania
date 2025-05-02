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

	var league models.League
	if err := db.First(&league).Error; err != nil {
		return errors.New("no leagues found in the database")
	}

	teams := []models.Team{
		{UserId: user.Id, Name: "BIG CLAN", LeagueId: league.Id, Elo: 1123},
		{UserId: user.Id, Name: "Astralis", LeagueId: league.Id, Elo: 1245},
		{UserId: user.Id, Name: "Natus Vincere", LeagueId: league.Id, Elo: 182},
		{UserId: user.Id, Name: "G2 Esports", LeagueId: league.Id, Elo: 945},
		{UserId: user.Id, Name: "Team Liquid", LeagueId: league.Id, Elo: 885},
		{UserId: user.Id, Name: "FaZe Clan", LeagueId: league.Id, Elo: 812},
		{UserId: user.Id, Name: "Fnatic", LeagueId: league.Id, Elo: 789},
	}

	if err := db.Create(&teams).Error; err != nil {
		return err
	}

	return nil
}
