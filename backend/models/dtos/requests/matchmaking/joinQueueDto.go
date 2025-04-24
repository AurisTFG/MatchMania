package matchmaking

import "github.com/google/uuid"

type JoinQueueDto struct {
	SeasonId uuid.UUID `json:"seasonId" validate:"required"`
	TeamId   uuid.UUID `json:"teamId"   validate:"required"`
}
