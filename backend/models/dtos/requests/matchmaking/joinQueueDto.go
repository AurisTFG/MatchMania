package matchmaking

import "github.com/google/uuid"

type JoinQueueDto struct {
	LeagueId uuid.UUID `json:"leagueId" validate:"required"`
	TeamId   uuid.UUID `json:"teamId"   validate:"required"`
}
