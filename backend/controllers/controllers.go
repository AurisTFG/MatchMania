package controllers

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/services"
)

type Controllers struct {
	AuthController   AuthController
	UserController   UserController
	SeasonController SeasonController
	TeamController   TeamController
	ResultController ResultController
}

func NewControllers(
	AuthController AuthController,
	userController UserController,
	SeasonController SeasonController,
	TeamController TeamController,
	ResultController ResultController,
) Controllers {
	return Controllers{
		AuthController:   AuthController,
		UserController:   userController,
		SeasonController: SeasonController,
		TeamController:   TeamController,
		ResultController: ResultController,
	}
}

func SetupControllers(
	db *config.DB,
	env *config.Env,
) *Controllers {
	sessionRepository := repositories.NewSessionRepository(db)
	userRepository := repositories.NewUserRepository(db)
	seasonRepository := repositories.NewSeasonRepository(db)
	teamRepository := repositories.NewTeamRepository(db)
	resultRepository := repositories.NewResultRepository(db)

	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService(sessionRepository, userRepository, env)
	seasonService := services.NewSeasonService(seasonRepository)
	teamService := services.NewTeamService(teamRepository)
	resultService := services.NewResultService(resultRepository)

	authController := NewAuthController(authService, userService, env)
	userController := NewUserController(userService)
	seasonController := NewSeasonController(seasonService)
	teamController := NewTeamController(seasonService, teamService)
	resultController := NewResultController(teamService, resultService)

	return &Controllers{
		AuthController:   authController,
		UserController:   userController,
		SeasonController: seasonController,
		TeamController:   teamController,
		ResultController: resultController,
	}
}
