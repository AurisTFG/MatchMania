package initializers

import (
	"MatchManiaAPI/models"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/logger"
)

type CustomTeamCreationType struct {
	dto      models.CreateTeamDto
	seasonID uint
}

type CustomResultCreationType struct {
	dto      models.CreateResultDto
	seasonID uint
	teamID   uint
}

func SeedDatabase() error {
	originalLogger := DB.Logger
	DB.Logger = originalLogger.LogMode(logger.Silent)

	tables := []string{"results", "teams", "seasons", "users"}

	// Delete all rows from all tables
	for _, table := range tables {
		if err := DB.Exec("DELETE FROM " + table).Error; err != nil {
			return fmt.Errorf("failed to delete from table %s: %w", table, err)
		}
	}

	// Reset ID counters for all tables
	for _, table := range tables {
		seqName := fmt.Sprintf("%s_id_seq", table)
		if err := DB.Exec("ALTER SEQUENCE " + seqName + " RESTART WITH 1").Error; err != nil {
			return fmt.Errorf("failed to reset sequence for table %s: %w", table, err)
		}
	}

	DB.Exec(`DROP TYPE IF EXISTS role;`)
	DB.Exec(`CREATE TYPE role AS ENUM ('admin', 'moderator', 'user');`)

	seasons := []models.CreateSeasonDto{
		{Name: "TO BE DELETED", StartDate: time.Now(), EndDate: time.Now().AddDate(0, 0, 30)},
		{Name: "Fall 2024", StartDate: time.Date(2024, 9, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)},
		{Name: "Winter 2025", StartDate: time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2025, 3, 31, 0, 0, 0, 0, time.UTC)},
		{Name: "Spring 2025", StartDate: time.Date(2025, 4, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2025, 6, 30, 0, 0, 0, 0, time.UTC)},
		{Name: "Summer 2025", StartDate: time.Date(2025, 7, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2025, 8, 31, 0, 0, 0, 0, time.UTC)},
		{Name: "Fall 2025", StartDate: time.Date(2025, 9, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2025, 12, 31, 0, 0, 0, 0, time.UTC)},
		{Name: "Winter 2026", StartDate: time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2026, 3, 31, 0, 0, 0, 0, time.UTC)},
		{Name: "Spring 2026", StartDate: time.Date(2026, 4, 1, 0, 0, 0, 0, time.UTC), EndDate: time.Date(2026, 6, 30, 0, 0, 0, 0, time.UTC)},
	}

	teams := []CustomTeamCreationType{
		{models.CreateTeamDto{Name: "TO BE DELETED"}, 1},
		{models.CreateTeamDto{Name: "BIG CLAN"}, 2},
		{models.CreateTeamDto{Name: "Astralis"}, 2},
		{models.CreateTeamDto{Name: "Natus Vincere"}, 2},
		{models.CreateTeamDto{Name: "G2 Esports"}, 2},
		{models.CreateTeamDto{Name: "Team Liquid"}, 2},
		{models.CreateTeamDto{Name: "FaZe Clan"}, 2},
		{models.CreateTeamDto{Name: "Fnatic"}, 2},
	}

	results := []CustomResultCreationType{
		{models.CreateResultDto{MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: 19, OpponentScore: 9, OpponentTeamID: 2}, 2, 3},
		{models.CreateResultDto{MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: 15, OpponentScore: 5, OpponentTeamID: 3}, 2, 4},
		{models.CreateResultDto{MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: 16, OpponentScore: 13, OpponentTeamID: 4}, 2, 5},
		{models.CreateResultDto{MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: 8, OpponentScore: 6, OpponentTeamID: 5}, 2, 6},
		{models.CreateResultDto{MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: 11, OpponentScore: 2, OpponentTeamID: 6}, 2, 7},
		{models.CreateResultDto{MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: 7, OpponentScore: 8, OpponentTeamID: 7}, 2, 8},
		{models.CreateResultDto{MatchStartDate: time.Now(), MatchEndDate: time.Now().Add(40 * time.Minute), Score: 12, OpponentScore: 15, OpponentTeamID: 8}, 2, 2},
	}

	users := []models.User{
		{Username: "AdminXD", Email: "adminemail@gmail.com", Password: "admin", Role: models.AdminRole},
		{Username: "ModeratorXDD", Email: "moderatoremail@gmail.com", Password: "mod", Role: models.ModeratorRole},
		{Username: "UserXDDD", Email: "userremail@gmail.com", Password: "user", Role: models.UserRole},
	}

	for _, season := range seasons {
		newSeason := season.ToSeason()

		result := DB.Create(&newSeason)
		if result.Error != nil {
			return result.Error
		}
	}

	for _, team := range teams {
		newTeam := team.dto.ToTeam()
		newTeam.SeasonID = team.seasonID
		newTeam.Elo = 1000

		result := DB.Create(&newTeam)
		if result.Error != nil {
			return result.Error
		}
	}

	for _, result := range results {
		newResult := result.dto.ToResult()
		newResult.SeasonID = result.seasonID
		newResult.TeamID = result.teamID

		result := DB.Create(&newResult)
		if result.Error != nil {
			return result.Error
		}
	}

	for _, user := range users {
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user.Password = string(hash)

		result := DB.Create(&user)
		if result.Error != nil {
			return result.Error
		}
	}

	DB.Logger = originalLogger

	return nil
}
