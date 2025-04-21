package results

import (
	"MatchManiaAPI/models/dtos/responses/teams"
	"MatchManiaAPI/models/dtos/responses/users"
	"time"

	"github.com/google/uuid"
)

type ResultDto struct {
	Id            uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000" json:"id"`
	StartDate     time.Time `example:"2025-06-01T00:00:00Z"                 json:"startDate"`
	EndDate       time.Time `example:"2025-06-01T00:40:00Z"                 json:"endDate"`
	Score         string    `example:"16"                                   json:"score"`
	OpponentScore string    `example:"14"                                   json:"opponentScore"`

	Team         teams.TeamMinimalDto `json:"team"`
	OpponentTeam teams.TeamMinimalDto `json:"opponentTeam"`
	User         users.UserMinimalDto `json:"user"`
}
