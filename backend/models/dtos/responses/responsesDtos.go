package responses

import (
	"MatchManiaAPI/models/dtos/responses/errors"
	"MatchManiaAPI/models/dtos/responses/leagues"
	"MatchManiaAPI/models/dtos/responses/matchmaking"
	"MatchManiaAPI/models/dtos/responses/results"
	"MatchManiaAPI/models/dtos/responses/teams"
	"MatchManiaAPI/models/dtos/responses/trackmanioauth"
	"MatchManiaAPI/models/dtos/responses/users"
)

type ErrorDto = errors.ErrorDto
type ResultDto = results.ResultDto
type LeagueDto = leagues.LeagueDto
type TeamDto = teams.TeamDto
type UserDto = users.UserDto
type UserMinimalDto = users.UserMinimalDto
type QueueTeamsCountDto = matchmaking.QueueTeamsCountDto
type MatchStatusDto = matchmaking.MatchStatusDto
type TrackmaniaOAuthAccessTokenDto = trackmanioauth.TrackmaniaOAuthAccessTokenDto
type TrackmaniaOAuthUrlDto = trackmanioauth.TrackmaniaOAuthUrlDto
type TrackmaniaOAuthUserDto = trackmanioauth.TrackmaniaOAuthUserDto
type TrackmaniaTracksDto = trackmanioauth.TrackmaniaTracksDto
