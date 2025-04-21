package teams

import (
	"github.com/google/uuid"
)

type TeamMinimalDto struct {
	Id   uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000" json:"id"`
	Name string    `example:"BIG Clan"                             json:"name"`
}
