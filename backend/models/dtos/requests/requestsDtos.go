package requests

import (
	"MatchManiaAPI/models/dtos/requests/auth"
	"MatchManiaAPI/models/dtos/requests/leagues"
	"MatchManiaAPI/models/dtos/requests/queues"
	"MatchManiaAPI/models/dtos/requests/results"
	"MatchManiaAPI/models/dtos/requests/teams"
	"MatchManiaAPI/models/dtos/requests/trackmaniaapi"
)

type LoginDto = auth.LoginDto
type SignupDto = auth.SignupDto
type CreateResultDto = results.CreateResultDto
type UpdateResultDto = results.UpdateResultDto
type CreateLeagueDto = leagues.CreateLeagueDto
type UpdateLeagueDto = leagues.UpdateLeagueDto
type CreateTeamDto = teams.CreateTeamDto
type UpdateTeamDto = teams.UpdateTeamDto
type JoinQueueDto = queues.JoinQueueDto
type LeaveQueueDto = queues.LeaveQueueDto

type CreateCompetitionDto = trackmaniaapi.CreateCompetitionDto
type CreateCompetitionTeamDto = trackmaniaapi.CreateCompetitionTeamDto
