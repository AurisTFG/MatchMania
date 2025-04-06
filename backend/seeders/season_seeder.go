package seeders

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
	"errors"
	"time"
)

func SeedSeasons(db *config.DB, env *config.Env) error {
	var count int64
	if err := db.Model(&models.Season{}).Count(&count).Error; err != nil {
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

	seasons := []models.Season{
		{
			UserUUID:  users[0].UUID,
			Name:      "Fall 2024",
			StartDate: time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			UserUUID:  users[0].UUID,
			Name:      "Winter 2025",
			StartDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2025, 3, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			UserUUID:  users[0].UUID,
			Name:      "Spring 2025",
			StartDate: time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2025, 6, 30, 0, 0, 0, 0, time.UTC),
		},
		{
			UserUUID:  users[0].UUID,
			Name:      "Summer 2025",
			StartDate: time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2025, 8, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			UserUUID:  users[0].UUID,
			Name:      "Fall 2025",
			StartDate: time.Date(2025, 9, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			UserUUID:  users[0].UUID,
			Name:      "Winter 2026",
			StartDate: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2026, 3, 31, 0, 0, 0, 0, time.UTC),
		},
		{
			UserUUID:  users[0].UUID,
			Name:      "Spring 2026",
			StartDate: time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC),
			EndDate:   time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC),
		},
	}

	for _, season := range seasons {
		if err := db.Create(&season).Error; err != nil {
			return err
		}
	}

	return nil
}
