package routes

import (
	"MatchManiaAPI/controllers"
	"MatchManiaAPI/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	server *gin.Engine,
	c *controllers.Controllers,
	m *middlewares.Middlewares,
) {
	v1 := server.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.GET("/me", m.AuthMiddleware.RequireAuth, c.AuthController.GetMe)
			auth.POST("/signup", c.AuthController.SignUp)
			auth.POST("/login", c.AuthController.LogIn)
			auth.POST("/logout", c.AuthController.LogOut)
			auth.POST("/refresh", c.AuthController.RefreshToken)
		}

		users := v1.Group("/users")
		{
			users.GET("", m.AuthMiddleware.RequireAuth, c.UserController.GetAllUsers)
			users.GET(":userId", m.AuthMiddleware.RequireAuth, c.UserController.GetUserByID)
		}

		seasons := v1.Group("/seasons")
		{
			seasons.GET("", c.SeasonController.GetAllSeasons)
			seasons.GET(":seasonId", c.SeasonController.GetSeason)
			seasons.POST("", m.AuthMiddleware.RequireAuth, c.SeasonController.CreateSeason)
			seasons.PATCH(":seasonId", m.AuthMiddleware.RequireAuth, c.SeasonController.UpdateSeason)
			seasons.DELETE(":seasonId", m.AuthMiddleware.RequireAuth, c.SeasonController.DeleteSeason)

			teams := seasons.Group("/:seasonId/teams")
			{
				teams.GET("", c.TeamController.GetAllTeams)
				teams.GET(":teamId", c.TeamController.GetTeam)
				teams.POST("", m.AuthMiddleware.RequireAuth, c.TeamController.CreateTeam)
				teams.PATCH(":teamId", m.AuthMiddleware.RequireAuth, c.TeamController.UpdateTeam)
				teams.DELETE(":teamId", m.AuthMiddleware.RequireAuth, c.TeamController.DeleteTeam)

				results := teams.Group("/:teamId/results")
				{
					results.GET("", c.ResultController.GetAllResults)
					results.GET(":resultId", c.ResultController.GetResult)
					results.POST("", m.AuthMiddleware.RequireAuth, c.ResultController.CreateResult)
					results.PATCH(":resultId", m.AuthMiddleware.RequireAuth, c.ResultController.UpdateResult)
					results.DELETE(":resultId", m.AuthMiddleware.RequireAuth, c.ResultController.DeleteResult)
				}
			}
		}
	}
}
