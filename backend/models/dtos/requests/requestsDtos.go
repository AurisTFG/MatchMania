package requests

import (
	"MatchManiaAPI/models/dtos/requests/auth"
	"MatchManiaAPI/models/dtos/requests/results"
	"MatchManiaAPI/models/dtos/requests/seasons"
	"MatchManiaAPI/models/dtos/requests/teams"
)

type LoginDto = auth.LoginDto
type SignupDto = auth.SignupDto
type CreateResultDto = results.CreateResultDto
type UpdateResultDto = results.UpdateResultDto
type CreateSeasonDto = seasons.CreateSeasonDto
type UpdateSeasonDto = seasons.UpdateSeasonDto
type CreateTeamDto = teams.CreateTeamDto
type UpdateTeamDto = teams.UpdateTeamDto
