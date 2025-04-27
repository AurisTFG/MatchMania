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

	var user models.User
	if err := db.First(&user).Error; err != nil {
		return errors.New("no users found in the database")
	}

	var season models.Season
	if err := db.First(&season).Error; err != nil {
		return errors.New("no seasons found in the database")
	}

	var teams []models.Team
	if err := db.Limit(2).Find(&teams).Error; err != nil {
		return errors.New("no teams found in the database")
	}

	results := []models.Result{
		{
			UserId:         user.Id,
			StartDate:      time.Now(),
			EndDate:        time.Now().Add(40 * time.Minute),
			Score:          "19",
			OpponentScore:  "9",
			TeamId:         teams[0].Id,
			OpponentTeamId: teams[1].Id,
			SeasonId:       season.Id,
		},
		{
			UserId:         user.Id,
			StartDate:      time.Now(),
			EndDate:        time.Now().Add(40 * time.Minute),
			Score:          "15",
			OpponentScore:  "5",
			TeamId:         teams[0].Id,
			OpponentTeamId: teams[1].Id,
			SeasonId:       season.Id,
		},
		{
			UserId:         user.Id,
			StartDate:      time.Now(),
			EndDate:        time.Now().Add(40 * time.Minute),
			Score:          "16",
			OpponentScore:  "13",
			TeamId:         teams[0].Id,
			OpponentTeamId: teams[1].Id,
			SeasonId:       season.Id,
		},
		{
			UserId:         user.Id,
			StartDate:      time.Now(),
			EndDate:        time.Now().Add(40 * time.Minute),
			Score:          "8",
			OpponentScore:  "6",
			TeamId:         teams[0].Id,
			OpponentTeamId: teams[1].Id,
			SeasonId:       season.Id,
		},
		{
			UserId:         user.Id,
			StartDate:      time.Now(),
			EndDate:        time.Now().Add(40 * time.Minute),
			Score:          "11",
			OpponentScore:  "2",
			TeamId:         teams[0].Id,
			OpponentTeamId: teams[1].Id,
			SeasonId:       season.Id,
		},
		{
			UserId:         user.Id,
			StartDate:      time.Now(),
			EndDate:        time.Now().Add(40 * time.Minute),
			Score:          "7",
			OpponentScore:  "8",
			TeamId:         teams[0].Id,
			OpponentTeamId: teams[1].Id,
			SeasonId:       season.Id,
		},
		{
			UserId:         user.Id,
			StartDate:      time.Now(),
			EndDate:        time.Now().Add(40 * time.Minute),
			Score:          "12",
			OpponentScore:  "15",
			TeamId:         teams[0].Id,
			OpponentTeamId: teams[1].Id,
			SeasonId:       season.Id,
		},
	}

	if err := db.Create(&results).Error; err != nil {
		return err
	}

	return nil
}
