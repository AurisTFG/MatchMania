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
		{
			UserId:    parsedUserId,
			Name:      "Red Bull Faster",
			StartDate: time.Date(2025, 1, 24, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC),
		},

		{
			UserId:    parsedUserId,
			Name:      "ComicCon Baltics 2024",
			StartDate: time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 5, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			UserId:    parsedUserId,
			Name:      "Trackmania World Cup 2024",
			StartDate: time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			UserId:    parsedUserId,
			Name:      "Winter 2024",
			StartDate: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			UserId:    parsedUserId,
			Name:      "Spring 2024",
			StartDate: time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 6, 30, 0, 0, 0, 0, time.UTC),
		},
		{
			UserId:    parsedUserId,
			Name:      "Summer 2024",
			StartDate: time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 8, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			UserId:    parsedUserId,
			Name:      "Fall 2024",
			StartDate: time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		},
	}

	if err = db.Create(&leagues).Error; err != nil {
		return err
	}

	return nil
}
