package initializers

import (
	"MatchManiaAPI/models"
	"log"
	"time"
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

func SeedDatabase() {
	DB.Exec("DELETE FROM results;")
	DB.Exec("DELETE FROM teams;")
	DB.Exec("DELETE FROM seasons;")

	DB.Exec("ALTER SEQUENCE seasons_id_seq RESTART WITH 1;")
	DB.Exec("ALTER SEQUENCE teams_id_seq RESTART WITH 1;")
	DB.Exec("ALTER SEQUENCE results_id_seq RESTART WITH 1;")

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

	for _, season := range seasons {
		DB.Create(&season)
	}

	for _, team := range teams {
		newTeam := team.dto.ToTeam()
		newTeam.SeasonID = team.seasonID
		newTeam.Elo = 1000

		DB.Create(&newTeam)
	}

	for _, result := range results {
		newResult := result.dto.ToResult()
		newResult.SeasonID = result.seasonID
		newResult.TeamID = result.teamID

		DB.Create(&newResult)
	}

	log.Println("Database seeded successfully!")
}
