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
			Name:      "Fall 2024",
			StartDate: time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			UserId:    parsedUserId,
			Name:      "Winter 2025",
			StartDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2025, 3, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			UserId:    parsedUserId,
			Name:      "Spring 2025",
			StartDate: time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2025, 6, 30, 0, 0, 0, 0, time.UTC),
		},
		{
			UserId:    parsedUserId,
			Name:      "Summer 2025",
			StartDate: time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2025, 8, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			UserId:    parsedUserId,
			Name:      "Fall 2025",
			StartDate: time.Date(2025, 9, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			UserId:    parsedUserId,
			Name:      "Winter 2026",
			StartDate: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2026, 3, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			UserId:    parsedUserId,
			Name:      "Spring 2026",
			StartDate: time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC),
		},
	}

	if err = db.Create(&leagues).Error; err != nil {
		return err
	}

	return nil
}
