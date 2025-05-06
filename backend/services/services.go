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
	EloService             EloService
	MatchmakingService     MatchmakingService
}

func SetupServices(
	env *config.Env,
	repos *repositories.Repositories,
) *Services {
	eloService := NewEloService()

	return &Services{
		UserService:            NewUserService(repos.UserRepository, repos.TrackmaniaTrackRepository),
		PlayerService:          NewPlayerService(repos.PlayerRepository),
		AuthService:            NewAuthService(repos.SessionRepository, repos.UserRepository, env),
		LeagueService:          NewLeagueService(repos.LeagueRepository),
		TeamService:            NewTeamService(repos.TeamRepository),
		ResultService:          NewResultService(repos.ResultRepository, repos.TeamRepository, eloService),
		TrackmaniaOAuthService: NewTrackmaniaOAuthService(repos.TrackmaniaOAuthStateRepository, env),
		EloService:             eloService,
		MatchmakingService:     NewMatchmakingService(),
	}
}
