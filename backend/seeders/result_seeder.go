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

	var league models.League
	if err := db.First(&league).Error; err != nil {
		return errors.New("no leagues found in the database")
	}

	var teams []models.Team
	if err := db.Limit(2).Find(&teams).Error; err != nil {
		return errors.New("no teams found in the database")
	}

	results := []models.Result{
		{
			UserId:          &user.Id,
			StartDate:       time.Now(),
			EndDate:         time.Now().Add(40 * time.Minute),
			Score:           4,
			OpponentScore:   1,
			EloDiff:         15,
			OpponentEloDiff: -8,
			TeamId:          teams[0].Id,
			OpponentTeamId:  teams[1].Id,
			LeagueId:        league.Id,
		},
		{
			UserId:          &user.Id,
			StartDate:       time.Now(),
			EndDate:         time.Now().Add(40 * time.Minute),
			Score:           1,
			OpponentScore:   3,
			EloDiff:         -11,
			OpponentEloDiff: 15,
			TeamId:          teams[0].Id,
			OpponentTeamId:  teams[1].Id,
			LeagueId:        league.Id,
		},
		{
			UserId:          &user.Id,
			StartDate:       time.Now(),
			EndDate:         time.Now().Add(40 * time.Minute),
			Score:           3,
			OpponentScore:   2,
			EloDiff:         12,
			OpponentEloDiff: -3,
			TeamId:          teams[0].Id,
			OpponentTeamId:  teams[1].Id,
			LeagueId:        league.Id,
		},
		{
			UserId:          &user.Id,
			StartDate:       time.Now(),
			EndDate:         time.Now().Add(40 * time.Minute),
			Score:           0,
			OpponentScore:   4,
			EloDiff:         -20,
			OpponentEloDiff: 10,
			TeamId:          teams[0].Id,
			OpponentTeamId:  teams[1].Id,
			LeagueId:        league.Id,
		},
		{
			UserId:          &user.Id,
			StartDate:       time.Now(),
			EndDate:         time.Now().Add(40 * time.Minute),
			Score:           4,
			OpponentScore:   3,
			EloDiff:         10,
			OpponentEloDiff: -10,
			TeamId:          teams[0].Id,
			OpponentTeamId:  teams[1].Id,
			LeagueId:        league.Id,
		},
		{
			UserId:          &user.Id,
			StartDate:       time.Now(),
			EndDate:         time.Now().Add(40 * time.Minute),
			Score:           0,
			OpponentScore:   1,
			EloDiff:         -5,
			OpponentEloDiff: 5,
			TeamId:          teams[0].Id,
			OpponentTeamId:  teams[1].Id,
			LeagueId:        league.Id,
		},
		{
			UserId:          nil,
			StartDate:       time.Now(),
			EndDate:         time.Now().Add(40 * time.Minute),
			Score:           1,
			OpponentScore:   0,
			EloDiff:         9,
			OpponentEloDiff: -4,
			TeamId:          teams[0].Id,
			OpponentTeamId:  teams[1].Id,
			LeagueId:        league.Id,
		},
	}

	if err := db.Create(&results).Error; err != nil {
		return err
	}

	return nil
}
