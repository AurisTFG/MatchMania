package services

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/repositories"
)

type Services struct {
	AppSettingService      AppSettingService
	UserService            UserService
	PlayerService          PlayerService
	AuthService            AuthService
	LeagueService          LeagueService
	TeamService            TeamService
	ResultService          ResultService
	EloService             EloService
	QueueService           QueueService
	MatchService           MatchService
	UbisoftApiService      UbisoftApiService
	NadeoApiService        NadeoApiService
	TrackmaniaApiService   TrackmaniaApiService
	TrackmaniaOAuthService TrackmaniaOAuthService
}

func NewServices(
	env *config.Env,
	repos *repositories.Repositories,
) *Services {
	appSettingService := NewAppSettingService(repos.AppSettingRepository)
	ubisoftApiService := NewUbisoftApiService(env, appSettingService)
	nadeoApiService := NewNadeoApiService(appSettingService)
	trackmaniaApiService := NewTrackmaniaApiService(env, ubisoftApiService, nadeoApiService, appSettingService)
	eloService := NewEloService()
	resultService := NewResultService(
		repos.ResultRepository,
		repos.TeamRepository,
		eloService,
	)

	return &Services{
		AppSettingService: appSettingService,
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
		EloService:             eloService,
		QueueService:           NewQueueService(repos.QueueRepository, repos.TeamRepository),
		MatchService:           NewMatchService(repos.MatchRepository, resultService, trackmaniaApiService),
		UbisoftApiService:      ubisoftApiService,
		NadeoApiService:        nadeoApiService,
		TrackmaniaApiService:   trackmaniaApiService,
		TrackmaniaOAuthService: NewTrackmaniaOAuthService(repos.TrackmaniaOAuthStateRepository, env),
	}
}
