package results

import (
	"MatchManiaAPI/models/dtos/responses/leagues"
	"MatchManiaAPI/models/dtos/responses/teams"
	"MatchManiaAPI/models/dtos/responses/users"
	"time"

	"github.com/google/uuid"
)

type ResultDto struct {
	Id              uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000" json:"id"`
	StartDate       time.Time `example:"2025-06-01T00:00:00Z"                 json:"startDate"`
	EndDate         time.Time `example:"2025-06-01T00:40:00Z"                 json:"endDate"`
	Score           uint      `example:"4"                                    json:"score"`
	OpponentScore   uint      `example:"0"                                    json:"opponentScore"`
	EloDiff         int       `example:"15"                                   json:"eloDiff"`
	OpponentEloDiff int       `example:"-8"                                   json:"opponentEloDiff"`

	League       leagues.LeagueMinimalDto `json:"league"`
	Team         teams.TeamMinimalDto     `json:"team"`
	OpponentTeam teams.TeamMinimalDto     `json:"opponentTeam"`
	User         users.UserMinimalDto     `json:"user"`
}
