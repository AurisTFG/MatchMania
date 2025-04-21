package results

import (
	"time"

	"github.com/google/uuid"
)

type CreateResultDto struct {
	StartDate      time.Time `example:"2025-06-01" json:"startDate"      validate:"required,minDate=-30,maxDate=3650"`
	EndDate        time.Time `example:"2025-06-01" json:"endDate"        validate:"required,maxDate=3650,dateDiff=7"`
	Score          string    `example:"16"         json:"score"          validate:"score"`
	OpponentScore  string    `example:"14"         json:"opponentScore"  validate:"score"`
	OpponentTeamId uuid.UUID `example:"4"          json:"opponentTeamId" validate:"required"`
}
