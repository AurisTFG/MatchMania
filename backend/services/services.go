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
	QueueService           QueueService
	MatchService           MatchService
}

func NewServices(
	env *config.Env,
	repos *repositories.Repositories,
) *Services {
	eloService := NewEloService()
	resultService := NewResultService(
		repos.ResultRepository,
		repos.TeamRepository,
		eloService,
	)

	return &Services{
		UserService: NewUserService(
			repos.UserRepository,
			repos.RoleRepository,
			repos.TrackmaniaTrackRepository,
		),
		PlayerService:          NewPlayerService(repos.PlayerRepository),
		AuthService:            NewAuthService(repos.SessionRepository, repos.UserRepository, env),
		LeagueService:          NewLeagueService(repos.LeagueRepository),
		TeamService:            NewTeamService(repos.TeamRepository),
		ResultService:          resultService,
		TrackmaniaOAuthService: NewTrackmaniaOAuthService(repos.TrackmaniaOAuthStateRepository, env),
		EloService:             eloService,
		QueueService:           NewQueueService(repos.QueueRepository, repos.TeamRepository),
		MatchService:           NewMatchService(repos.MatchRepository, resultService),
	}
}
