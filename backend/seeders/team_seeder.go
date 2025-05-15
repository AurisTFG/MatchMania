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
	if err = db.
		Order("end_date DESC").
		First(&leagues).
		Error; err != nil {
		return errors.New("no leagues found in the database")
	}

	teams := []models.Team{
		createTeam(parsedUserId, leagues, "BIG CLAN", 999),
		createTeam(parsedUserId, leagues, "Astralis", 967),
		createTeam(parsedUserId, leagues, "Natus Vincere", 1245),
		createTeam(parsedUserId, leagues, "G2 Esports", 945),
		createTeam(parsedUserId, leagues, "Team Liquid", 885),
		createTeam(parsedUserId, leagues, "FaZe Clan", 812),
		createTeam(parsedUserId, leagues, "Fnatic", 789),
	}

	if err = db.Create(&teams).Error; err != nil {
		return err
	}

	return nil
}

func createTeam(userId uuid.UUID, leagues []models.League, name string, elo uint) models.Team {
	return models.Team{
		Name:    name,
		LogoUrl: nil,
		Elo:     elo,
		UserId:  userId,
		Leagues: leagues,
	}
}
