package services

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/repositories"
)

type Services struct {
	UserService            UserService
	PlayerService          PlayerService
	AuthService            AuthService
	LeagueService          LeagueService
	TeamService            TeamService
	ResultService          ResultService
	TrackmaniaOAuthService TrackmaniaOAuthService
	MatchmakingService     MatchmakingService
}

func SetupServices(
	env *config.Env,
	repos *repositories.Repositories,
) *Services {
	return &Services{
		UserService:            NewUserService(repos.UserRepository, repos.TrackmaniaTrackRepository),
		PlayerService:          NewPlayerService(repos.PlayerRepository),
		AuthService:            NewAuthService(repos.SessionRepository, repos.UserRepository, env),
		LeagueService:          NewLeagueService(repos.LeagueRepository),
		TeamService:            NewTeamService(repos.TeamRepository),
		ResultService:          NewResultService(repos.ResultRepository),
		TrackmaniaOAuthService: NewTrackmaniaOAuthService(repos.TrackmaniaOAuthStateRepository, env),
		MatchmakingService:     NewMatchmakingService(),
	}
}
