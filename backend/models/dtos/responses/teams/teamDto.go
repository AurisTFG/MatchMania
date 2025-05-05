package teams

import (
	"MatchManiaAPI/models/dtos/responses/leagues"
	"MatchManiaAPI/models/dtos/responses/players"
	"MatchManiaAPI/models/dtos/responses/users"

	"github.com/google/uuid"
)

type TeamDto struct {
	Id      uuid.UUID `example:"550e8400-e29b-41d4-a716-446655440000" json:"id"`
	Name    string    `example:"BIG Clan"                             json:"name"`
	LogoUrl string    `example:"https://example.com/logo.png"         json:"logoUrl"`
	Elo     uint      `example:"1200"                                 json:"elo"`

	User    users.UserMinimalDto       `json:"user"`
	Players []players.PlayerMinimalDto `json:"players"`
	Leagues []leagues.LeagueMinimalDto `json:"leagues"`
}
