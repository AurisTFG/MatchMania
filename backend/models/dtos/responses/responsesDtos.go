package responses

import (
	"MatchManiaAPI/models/dtos/responses/errors"
	"MatchManiaAPI/models/dtos/responses/leagues"
	"MatchManiaAPI/models/dtos/responses/matches"
	"MatchManiaAPI/models/dtos/responses/players"
	"MatchManiaAPI/models/dtos/responses/queues"
	"MatchManiaAPI/models/dtos/responses/results"
	"MatchManiaAPI/models/dtos/responses/teams"
	"MatchManiaAPI/models/dtos/responses/trackmanioauth"
	"MatchManiaAPI/models/dtos/responses/users"
)

type ErrorDto = errors.ErrorDto
type ResultDto = results.ResultDto
type LeagueDto = leagues.LeagueDto
type LeagueMinimalDto = leagues.LeagueMinimalDto
type TeamDto = teams.TeamDto
type UserDto = users.UserDto
type UserMinimalDto = users.UserMinimalDto
type PlayerMinimalDto = players.PlayerMinimalDto
type TrackmaniaOAuthAccessTokenDto = trackmanioauth.TrackmaniaOAuthAccessTokenDto
type TrackmaniaOAuthUrlDto = trackmanioauth.TrackmaniaOAuthUrlDto
type TrackmaniaOAuthUserDto = trackmanioauth.TrackmaniaOAuthUserDto
type TrackmaniaOAuthFavoritesDto = trackmanioauth.TrackmaniaOAuthFavoritesDto
type QueueDto = queues.QueueDto
type MatchDto = matches.MatchDto
