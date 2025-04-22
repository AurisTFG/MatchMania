package services

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/repositories"
)

type Services struct {
	UserService        UserService
	AuthService        AuthService
	SeasonService      SeasonService
	TeamService        TeamService
	ResultService      ResultService
	MatchmakingService MatchmakingService
}

func SetupServices(repos *repositories.Repositories, env *config.Env) *Services {
	return &Services{
		UserService:        NewUserService(repos.UserRepository),
		AuthService:        NewAuthService(repos.SessionRepository, repos.UserRepository, env),
		SeasonService:      NewSeasonService(repos.SeasonRepository),
		TeamService:        NewTeamService(repos.TeamRepository),
		ResultService:      NewResultService(repos.ResultRepository),
		MatchmakingService: NewMatchmakingService(),
	}
}
