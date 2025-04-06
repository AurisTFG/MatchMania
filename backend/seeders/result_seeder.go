package seeders

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
	"errors"
	"time"
)

func SeedResults(db *config.DB, env *config.Env) error {
	var count int64
	if err := db.Model(&models.Result{}).Count(&count).Error; err != nil {
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

	results := []models.Result{
		{
			UserUUID:       users[0].UUID,
			MatchStartDate: time.Now(),
			MatchEndDate:   time.Now().Add(40 * time.Minute),
			Score:          "19",
			OpponentScore:  "9",
			TeamID:         2,
			OpponentTeamID: 1,
			SeasonID:       1,
		},
		{
			UserUUID:       users[0].UUID,
			MatchStartDate: time.Now(),
			MatchEndDate:   time.Now().Add(40 * time.Minute),
			Score:          "15",
			OpponentScore:  "5",
			TeamID:         3,
			OpponentTeamID: 2,
			SeasonID:       1,
		},
		{
			UserUUID:       users[0].UUID,
			MatchStartDate: time.Now(),
			MatchEndDate:   time.Now().Add(40 * time.Minute),
			Score:          "16",
			OpponentScore:  "13",
			TeamID:         4,
			OpponentTeamID: 3,
			SeasonID:       1,
		},
		{
			UserUUID:       users[0].UUID,
			MatchStartDate: time.Now(),
			MatchEndDate:   time.Now().Add(40 * time.Minute),
			Score:          "8",
			OpponentScore:  "6",
			TeamID:         5,
			OpponentTeamID: 4,
			SeasonID:       1,
		},
		{
			UserUUID:       users[0].UUID,
			MatchStartDate: time.Now(),
			MatchEndDate:   time.Now().Add(40 * time.Minute),
			Score:          "11",
			OpponentScore:  "2",
			TeamID:         6,
			OpponentTeamID: 5,
			SeasonID:       1,
		},
		{
			UserUUID:       users[0].UUID,
			MatchStartDate: time.Now(),
			MatchEndDate:   time.Now().Add(40 * time.Minute),
			Score:          "7",
			OpponentScore:  "8",
			TeamID:         7,
			OpponentTeamID: 6,
			SeasonID:       1,
		},
		{
			UserUUID:       users[0].UUID,
			MatchStartDate: time.Now(),
			MatchEndDate:   time.Now().Add(40 * time.Minute),
			Score:          "12",
			OpponentScore:  "15",
			TeamID:         1,
			OpponentTeamID: 7,
			SeasonID:       1,
		},
	}

	for _, result := range results {
		if err := db.Create(&result).Error; err != nil {
			return err
		}
	}

	return nil
}
