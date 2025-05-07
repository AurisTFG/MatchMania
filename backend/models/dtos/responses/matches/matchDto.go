package matches

import (
	"MatchManiaAPI/models/dtos/responses/leagues"
	"MatchManiaAPI/models/dtos/responses/teams"

	"github.com/google/uuid"
)

type MatchDto struct {
	Id       uuid.UUID                `example:"b2c4f3e0-5d8a-4c1b-9f3e-7a1d2f3e4b5c" json:"id"`
	GameMode string                   `example:"2v2"                                  json:"gameMode"`
	League   leagues.LeagueMinimalDto `                                               json:"league"`
	Teams    []teams.TeamMinimalDto   `                                               json:"teams"`
}
