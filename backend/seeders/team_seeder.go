package seeders

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func SeedTeams(db *config.DB, env *config.Env) error {
	var count int64
	if err := db.Model(&models.Team{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var userId string
	if err := db.
		Model(&models.User{}).
		Order("username ASC").
		Select("id").
		Limit(1).
		Scan(&userId).
		Error; err != nil {
		return errors.New("no users found in the database")
	}
	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		return fmt.Errorf("failed to parse user ID: %w", err)
	}

	var leagues []models.League
	if err := db.
		Order("end_date DESC").
		First(&leagues).
		Error; err != nil {
		return errors.New("no leagues found in the database")
	}

	teams := []models.Team{
		{UserId: parsedUserId, Name: "BIG CLAN", Leagues: leagues, Elo: 999},
		{UserId: parsedUserId, Name: "Astralis", Leagues: leagues, Elo: 967},
		{UserId: parsedUserId, Name: "Natus Vincere", Leagues: leagues, Elo: 1245},
		{UserId: parsedUserId, Name: "G2 Esports", Leagues: leagues, Elo: 945},
		{UserId: parsedUserId, Name: "Team Liquid", Leagues: leagues, Elo: 885},
		{UserId: parsedUserId, Name: "FaZe Clan", Leagues: leagues, Elo: 812},
		{UserId: parsedUserId, Name: "Fnatic", Leagues: leagues, Elo: 789},
	}

	if err := db.Create(&teams).Error; err != nil {
		return err
	}

	return nil
}
