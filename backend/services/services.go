package services

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/repositories"
)

type Services struct {
	UserService            UserService
	AuthService            AuthService
	LeagueService          LeagueService
	TeamService            TeamService
	ResultService          ResultService
	TrackmaniaOAuthService TrackmaniaOAuthService
	MatchmakingService     MatchmakingService
}

func SetupServices(repos *repositories.Repositories, env *config.Env) *Services {
	return &Services{
		UserService:            NewUserService(repos.UserRepository, repos.TrackmaniaTrackRepository),
		AuthService:            NewAuthService(repos.SessionRepository, repos.UserRepository, env),
		LeagueService:          NewLeagueService(repos.LeagueRepository),
		TeamService:            NewTeamService(repos.TeamRepository),
		ResultService:          NewResultService(repos.ResultRepository),
		TrackmaniaOAuthService: NewTrackmaniaOAuthService(repos.TrackmaniaOAuthStateRepository, env),
		MatchmakingService:     NewMatchmakingService(),
	}
}
