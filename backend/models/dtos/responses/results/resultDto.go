package responses

import (
	"time"

	"github.com/google/uuid"
)

type ResultDto struct {
	Id             uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000" json:"id"`
	MatchStartDate time.Time `example:"2025-06-01T00:00:00Z" json:"matchStartDate"`
	MatchEndDate   time.Time `example:"2025-06-01T00:40:00Z" json:"matchEndDate"`
	Score          string    `example:"16"                   json:"score"`
	OpponentScore  string    `example:"14"                   json:"opponentScore"`
	SeasonId       uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000" json:"seasonId"`
	TeamId         uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000" json:"teamId"`
	OpponentTeamId uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000" json:"opponentTeamId"`
	UserId         uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000" json:"userId"`
}
