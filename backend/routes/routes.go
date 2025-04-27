package routes

import (
	"MatchManiaAPI/controllers"
	"MatchManiaAPI/middlewares"
	"MatchManiaAPI/models/enums"
	"MatchManiaAPI/services"
	r "MatchManiaAPI/utils/httpresponses"

	"github.com/gin-gonic/gin"
)

var authService services.AuthService

func SetupRoutes(
	server *gin.Engine,
	c *controllers.Controllers,
	s *services.Services,
) {
	authService = s.AuthService
	server.Use(middlewares.ErrorMiddleware())

	v1 := server.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.GET("/me", requirePerm(enums.LoggedInPermission), c.AuthController.GetMe)
			auth.POST("/signup", c.AuthController.SignUp)
			auth.POST("/login", c.AuthController.LogIn)
			auth.POST("/logout", c.AuthController.LogOut)
			auth.POST("/refresh", c.AuthController.RefreshToken)
		}

		users := v1.Group("/users")
		{
			users.GET("", requirePerm(enums.ManageUserPermission), c.UserController.GetAllUsers)
			users.GET(":userId", requirePerm(enums.ManageUserPermission), c.UserController.GetUserById)
		}

		trackmania := v1.Group("/trackmania")
		{
			oauth := trackmania.Group("/oauth")
			{
				oauth.GET("/url", requirePerm(enums.LoggedInPermission), c.TrackmaniaOAuthController.GetAuthUrl)
				oauth.GET("/callback", c.TrackmaniaOAuthController.HandleCallback)
			}
		}

		matchmaking := v1.Group("/matchmaking")
		{
			queue := matchmaking.Group("/queue")
			{
				queue.POST("/join", requirePerm(enums.ManageQueuePermission), c.MatchmakingController.JoinQueue)
				queue.POST("/leave", requirePerm(enums.ManageQueuePermission), c.MatchmakingController.LeaveQueue)
				queue.GET(
					"/teams-count/:seasonId",
					requirePerm(enums.LoggedInPermission),
					c.MatchmakingController.GetQueueTeamsCount,
				)
				queue.GET(
					"/status/:teamId",
					requirePerm(enums.ManageQueuePermission),
					c.MatchmakingController.CheckMatchStatus,
				)
			}
		}

		seasons := v1.Group("/seasons")
		{
			seasons.GET("", requirePerm(enums.LoggedInPermission), c.SeasonController.GetAllSeasons)
			seasons.GET(":seasonId", requirePerm(enums.LoggedInPermission), c.SeasonController.GetSeason)
			seasons.POST("", requirePerm(enums.ManageSeasonPermission), c.SeasonController.CreateSeason)
			seasons.PATCH(":seasonId", requirePerm(enums.ManageSeasonPermission), c.SeasonController.UpdateSeason)
			seasons.DELETE(":seasonId", requirePerm(enums.ManageSeasonPermission), c.SeasonController.DeleteSeason)

			teams := seasons.Group("/:seasonId/teams")
			{
				teams.GET("", requirePerm(enums.LoggedInPermission), c.TeamController.GetAllTeams)
				teams.GET(":teamId", requirePerm(enums.LoggedInPermission), c.TeamController.GetTeam)
				teams.POST("", requirePerm(enums.ManageTeamPermission), c.TeamController.CreateTeam)
				teams.PATCH(":teamId", requirePerm(enums.ManageTeamPermission), c.TeamController.UpdateTeam)
				teams.DELETE(":teamId", requirePerm(enums.ManageTeamPermission), c.TeamController.DeleteTeam)

				results := teams.Group("/:teamId/results")
				{
					results.GET("", requirePerm(enums.LoggedInPermission), c.ResultController.GetAllResults)
					results.GET(":resultId", requirePerm(enums.LoggedInPermission), c.ResultController.GetResult)
					results.POST("", requirePerm(enums.ManageResultPermission), c.ResultController.CreateResult)
					results.PATCH(
						":resultId",
						requirePerm(enums.ManageResultPermission),
						c.ResultController.UpdateResult,
					)
					results.DELETE(
						":resultId",
						requirePerm(enums.ManageResultPermission),
						c.ResultController.DeleteResult,
					)
				}
			}
		}
	}
}

func requirePerm(permission enums.Permission) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := middlewares.AuthMiddleware(c, authService)
		if err != nil {
			r.Unauthorized(c)
			return
		}

		if permission != enums.LoggedInPermission {
			err = middlewares.RequirePermission(c, string(permission))
			if err != nil {
				r.Forbidden(c)
				return
			}
		}

		c.Next()
	}
}
