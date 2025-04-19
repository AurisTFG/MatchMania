package requests

import (
	"time"

	"github.com/google/uuid"
)

type CreateResultDto struct {
	MatchStartDate time.Time `example:"2025-06-01T00:00:00Z" json:"matchStartDate" validate:"required,startDate"`
	MatchEndDate   time.Time `example:"2025-06-01T00:40:00Z" json:"matchEndDate"   validate:"required,endDate,dateDiff,gtfield=MatchStartDate"`
	Score          string    `example:"16"                   json:"score"          validate:"score"`
	OpponentScore  string    `example:"14"                   json:"opponentScore"  validate:"score"`
	OpponentTeamId uuid.UUID `example:"4"                    json:"opponentTeamId" validate:"required"`
}
