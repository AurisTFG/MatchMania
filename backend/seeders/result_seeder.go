package seeders

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func SeedResults(db *config.DB, env *config.Env) error {
	var count int64
	if err := db.Model(&models.Result{}).Count(&count).Error; err != nil {
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

	var league models.League
	if err := db.
		Order("end_date DESC").
		Find(&league).
		Error; err != nil {
		return errors.New("no leagues found in the database")
	}

	var teams []models.Team
	if err := db.
		Order("elo DESC").
		Limit(2).
		Find(&teams).
		Error; err != nil {
		return errors.New("no teams found in the database")
	}

	results := []models.Result{
		{
			UserId:          &parsedUserId,
			StartDate:       time.Now(),
			EndDate:         time.Now().Add(40 * time.Minute),
			Score:           4,
			OpponentScore:   1,
			EloDiff:         15,
			OpponentEloDiff: -8,
			NewElo:          1015,
			NewOpponentElo:  992,
			TeamId:          teams[0].Id,
			OpponentTeamId:  teams[1].Id,
			LeagueId:        league.Id,
		},
		{
			UserId:          &parsedUserId,
			StartDate:       time.Now(),
			EndDate:         time.Now().Add(40 * time.Minute),
			Score:           1,
			OpponentScore:   3,
			EloDiff:         -11,
			OpponentEloDiff: 15,
			NewElo:          989,
			NewOpponentElo:  1015,
			TeamId:          teams[0].Id,
			OpponentTeamId:  teams[1].Id,
			LeagueId:        league.Id,
		},
		{
			UserId:          &parsedUserId,
			StartDate:       time.Now(),
			EndDate:         time.Now().Add(40 * time.Minute),
			Score:           3,
			OpponentScore:   2,
			EloDiff:         12,
			OpponentEloDiff: -3,
			NewElo:          1012,
			NewOpponentElo:  997,
			TeamId:          teams[0].Id,
			OpponentTeamId:  teams[1].Id,
			LeagueId:        league.Id,
		},
		{
			UserId:          &parsedUserId,
			StartDate:       time.Now(),
			EndDate:         time.Now().Add(40 * time.Minute),
			Score:           0,
			OpponentScore:   4,
			EloDiff:         -20,
			OpponentEloDiff: 10,
			NewElo:          980,
			NewOpponentElo:  1010,
			TeamId:          teams[0].Id,
			OpponentTeamId:  teams[1].Id,
			LeagueId:        league.Id,
		},
		{
			UserId:          &parsedUserId,
			StartDate:       time.Now(),
			EndDate:         time.Now().Add(40 * time.Minute),
			Score:           4,
			OpponentScore:   3,
			EloDiff:         10,
			OpponentEloDiff: -10,
			NewElo:          1010,
			NewOpponentElo:  990,
			TeamId:          teams[0].Id,
			OpponentTeamId:  teams[1].Id,
			LeagueId:        league.Id,
		},
		{
			UserId:          &parsedUserId,
			StartDate:       time.Now(),
			EndDate:         time.Now().Add(40 * time.Minute),
			Score:           0,
			OpponentScore:   1,
			EloDiff:         -5,
			OpponentEloDiff: 5,
			NewElo:          995,
			NewOpponentElo:  1005,
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
			NewElo:          1009,
			NewOpponentElo:  996,
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
