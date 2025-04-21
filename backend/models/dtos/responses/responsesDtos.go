package responses

import (
	"MatchManiaAPI/models/dtos/responses/errors"
	"MatchManiaAPI/models/dtos/responses/results"
	"MatchManiaAPI/models/dtos/responses/seasons"
	"MatchManiaAPI/models/dtos/responses/teams"
	"MatchManiaAPI/models/dtos/responses/users"
)

type ErrorDto = errors.ErrorDto
type ResultDto = results.ResultDto
type SeasonDto = seasons.SeasonDto
type TeamDto = teams.TeamDto
type UserDto = users.UserDto
type UserMinimalDto = users.UserMinimalDto
