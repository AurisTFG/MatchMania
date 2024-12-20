package routes

import (
	"MatchManiaAPI/controllers"
	"MatchManiaAPI/middlewares"

	_ "MatchManiaAPI/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ApplyRoutes(
	server *gin.Engine,
	c *Controllers,
) {
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := server.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", c.AuthController.AuthSignUp)
			auth.POST("/login", c.AuthController.AuthLogIn)
			auth.POST("/logout", c.AuthController.AuthLogOut)
			auth.POST("/refresh-token", c.AuthController.AuthRefreshToken)
		}

		users := v1.Group("/users")
		{
			users.GET("", c.UserController.GetAllUsers)
			users.GET(":userId", c.UserController.GetUserByID)
		}

		seasons := v1.Group("/seasons")
		{
			seasons.GET("", c.SeasonController.GetAllSeasons)
			seasons.GET(":seasonId", c.SeasonController.GetSeason)
			seasons.POST("", c.AuthMiddleware.RequireAuth, c.SeasonController.CreateSeason)
			seasons.PATCH(":seasonId", c.AuthMiddleware.RequireAuth, c.SeasonController.UpdateSeason)
			seasons.DELETE(":seasonId", c.AuthMiddleware.RequireAuth, c.SeasonController.DeleteSeason)

			teams := seasons.Group("/:seasonId/teams")
			{
				teams.GET("", c.TeamController.GetAllTeams)
				teams.GET(":teamId", c.TeamController.GetTeam)
				teams.POST("", c.AuthMiddleware.RequireAuth, c.TeamController.CreateTeam)
				teams.PATCH(":teamId", c.AuthMiddleware.RequireAuth, c.TeamController.UpdateTeam)
				teams.DELETE(":teamId", c.AuthMiddleware.RequireAuth, c.TeamController.DeleteTeam)

				results := teams.Group("/:teamId/results")
				{
					results.GET("", c.ResultController.GetAllResults)
					results.GET(":resultId", c.ResultController.GetResult)
					results.POST("", c.AuthMiddleware.RequireAuth, c.ResultController.CreateResult)
					results.PATCH(":resultId", c.AuthMiddleware.RequireAuth, c.ResultController.UpdateResult)
					results.DELETE(":resultId", c.AuthMiddleware.RequireAuth, c.ResultController.DeleteResult)
				}
			}
		}
	}
}

type Controllers struct {
	AuthMiddleware   middlewares.AuthMiddleware
	AuthController   controllers.AuthController
	UserController   controllers.UserController
	SeasonController controllers.SeasonController
	TeamController   controllers.TeamController
	ResultController controllers.ResultController
}

func NewControllers(
	AuthMiddleware middlewares.AuthMiddleware,
	AuthController controllers.AuthController,
	userController controllers.UserController,
	SeasonController controllers.SeasonController,
	TeamController controllers.TeamController,
	ResultController controllers.ResultController,
) Controllers {
	return Controllers{
		AuthMiddleware:   AuthMiddleware,
		AuthController:   AuthController,
		UserController:   userController,
		SeasonController: SeasonController,
		TeamController:   TeamController,
		ResultController: ResultController,
	}
}
