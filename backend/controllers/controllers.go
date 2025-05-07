package controllers

import (
	"MatchManiaAPI/services"
)

type Controllers struct {
	AuthController            AuthController
	UserController            UserController
	PlayerController          PlayerController
	LeagueController          LeagueController
	TeamController            TeamController
	ResultController          ResultController
	TrackmaniaOAuthController TrackmaniaOAuthController
	QueueController           QueueController
	MatchController           MatchController
}

func NewControllers(
	services *services.Services,
) *Controllers {
	return &Controllers{
		AuthController:            NewAuthController(services.AuthService, services.UserService),
		UserController:            NewUserController(services.UserService),
		PlayerController:          NewPlayerController(services.PlayerService),
		LeagueController:          NewLeagueController(services.LeagueService),
		TeamController:            NewTeamController(services.LeagueService, services.TeamService),
		ResultController:          NewResultController(services.TeamService, services.ResultService),
		TrackmaniaOAuthController: NewTrackmaniaOAuthController(services.TrackmaniaOAuthService, services.UserService),
		QueueController:           NewQueueController(services.QueueService),
		MatchController:           NewMatchController(services.MatchService),
	}
}
