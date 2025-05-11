package teams

import (
	"MatchManiaAPI/models/dtos/responses/players"

	"github.com/google/uuid"
)

type TeamMinimalDto struct {
	Id      uuid.UUID                  `example:"550e8400-e29b-41d4-a716-446655440000" json:"id"`
	Name    string                     `example:"BIG Clan"                             json:"name"`
	LogoUrl string                     `example:"https://example.com/logo.png"         json:"logoUrl"`
	Players []players.PlayerMinimalDto `                                               json:"players"`
}
