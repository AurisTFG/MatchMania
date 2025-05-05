package results

import (
	"time"

	"github.com/google/uuid"
)

type CreateResultDto struct {
	LeagueId       uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000" json:"leagueId"       validate:"required"`
	StartDate      time.Time `example:"2025-06-01T00:00:000Z"                json:"startDate"      validate:"required,minDate=-30,maxDate=3650"`
	EndDate        time.Time `example:"2025-08-31T00:00:000Z"                json:"endDate"        validate:"required,maxDate=3650,maxDateDiff=1"`
	TeamId         uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000" json:"teamId"         validate:"required"`
	OpponentTeamId uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000" json:"opponentTeamId" validate:"required"`
	Score          string    `example:"16"                                   json:"score"          validate:"score"`
	OpponentScore  string    `example:"14"                                   json:"opponentScore"  validate:"score,nefield=Score"`
}
