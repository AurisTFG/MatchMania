package seeders

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func SeedLeagues(db *config.DB, env *config.Env) error {
	var count int64
	if err := db.Model(&models.League{}).Count(&count).Error; err != nil {
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
		Scan(&userId).Error; err != nil {
		return errors.New("no users found in the database")
	}

	parsedUserId, err := uuid.Parse(userId)
	if err != nil {
		return fmt.Errorf("failed to parse user ID: %w", err)
	}

	leagues := []models.League{
		createLeague(parsedUserId, "Winter 2024", 2024, 1),
		createLeague(parsedUserId, "Spring 2024", 2024, 4),
		createLeague(parsedUserId, "Summer 2024", 2024, 7),
		createLeague(parsedUserId, "Fall 2024", 2024, 10),
		createLeague(parsedUserId, "Red Bull Faster", 2025, 1),
		createLeague(parsedUserId, "ComicCon Baltics 2024", 2024, 2),
		createLeague(parsedUserId, "Trackmania World Cup 2024", 2024, 9),
	}

	if err = db.Create(&leagues).Error; err != nil {
		return err
	}

	return nil
}

func createLeague(userId uuid.UUID, name string, year int, month int) models.League {
	return models.League{
		UserId:    userId,
		Name:      name,
		StartDate: time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC),
		EndDate:   time.Date(year, time.Month(month+3), 30, 0, 0, 0, 0, time.UTC),
	}
}
