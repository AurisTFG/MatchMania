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

	var users []models.User
	if err := db.Find(&users).Error; err != nil {
		return err
	}

	if len(users) == 0 {
		return errors.New("no users found in the database")
	}

	teams := []models.Team{
		{UserUUID: users[0].UUID, Name: "BIG CLAN", SeasonID: 1, Elo: 1123},
		{UserUUID: users[0].UUID, Name: "Astralis", SeasonID: 1, Elo: 1245},
		{UserUUID: users[0].UUID, Name: "Natus Vincere", SeasonID: 1, Elo: 182},
		{UserUUID: users[0].UUID, Name: "G2 Esports", SeasonID: 1, Elo: 945},
		{UserUUID: users[0].UUID, Name: "Team Liquid", SeasonID: 1, Elo: 885},
		{UserUUID: users[0].UUID, Name: "FaZe Clan", SeasonID: 1, Elo: 812},
		{UserUUID: users[0].UUID, Name: "Fnatic", SeasonID: 1, Elo: 789},
	}

	for _, team := range teams {
		if err := db.Create(&team).Error; err != nil {
			return err
		}
	}

	return nil
}
