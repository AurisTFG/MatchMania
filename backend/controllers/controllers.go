package controllers

import (
	"MatchManiaAPI/services"
)

type Controllers struct {
	AuthController   AuthController
	UserController   UserController
	SeasonController SeasonController
	TeamController   TeamController
	ResultController ResultController
}

func SetupControllers(
	services *services.Services,
) *Controllers {
	authController := NewAuthController(services.AuthService, services.UserService)
	userController := NewUserController(services.UserService)
	seasonController := NewSeasonController(services.SeasonService)
	teamController := NewTeamController(services.SeasonService, services.TeamService)
	resultController := NewResultController(services.TeamService, services.ResultService)

	return &Controllers{
		AuthController:   authController,
		UserController:   userController,
		SeasonController: seasonController,
		TeamController:   teamController,
		ResultController: resultController,
	}
}
