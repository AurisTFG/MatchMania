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
var userService services.UserService

func SetupRoutes(
	server *gin.Engine,
	c *controllers.Controllers,
	s *services.Services,
) {
	authService = s.AuthService
	userService = s.UserService
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
			users.GET("", requirePerm(enums.ViewUserPermission), c.UserController.GetAllUsers)
			users.GET(":userId", requirePerm(enums.ManageUserPermission), c.UserController.GetUserById)
		}

		players := v1.Group("/players")
		{
			players.GET("", requirePerm(enums.ViewPlayerPermission), c.PlayerController.GetAllPlayers)
		}

		trackmania := v1.Group("/trackmania")
		{
			oauth := trackmania.Group("/oauth")
			{
				oauth.GET("/url", requirePerm(enums.LoggedInPermission), c.TrackmaniaOAuthController.GetAuthUrl)
				oauth.GET("/callback", c.TrackmaniaOAuthController.HandleCallback)
			}
		}

		teams := v1.Group("/teams")
		{
			teams.GET("", requirePerm(enums.ViewTeamPermission), c.TeamController.GetAllTeams)
			teams.GET(":teamId", requirePerm(enums.ViewTeamPermission), c.TeamController.GetTeam)
			teams.POST("", requirePerm(enums.ManageTeamPermission), c.TeamController.CreateTeam)
			teams.PATCH(":teamId", requirePerm(enums.ManageTeamPermission), c.TeamController.UpdateTeam)
			teams.DELETE(":teamId", requirePerm(enums.ManageTeamPermission), c.TeamController.DeleteTeam)
		}

		matchmaking := v1.Group("/matchmaking")
		{
			queues := matchmaking.Group("/queues")
			{
				queues.GET("", requirePerm(enums.ViewQueuePermission), c.QueueController.GetAllQueues)
				queues.POST("/join", requirePerm(enums.ManageQueuePermission), c.QueueController.JoinQueue)
				queues.POST("/leave", requirePerm(enums.ManageQueuePermission), c.QueueController.LeaveQueue)
			}

			matches := matchmaking.Group("/matches")
			{
				matches.GET("", requirePerm(enums.ViewMatchPermission), c.MatchController.GetAllMatches)
				matches.POST("/end", requirePerm(enums.ManageMatchPermission), c.MatchController.EndMatch)
			}
		}

		leagues := v1.Group("/leagues")
		{
			leagues.GET("", requirePerm(enums.ViewLeaguePermission), c.LeagueController.GetAllLeagues)
			leagues.GET(":leagueId", requirePerm(enums.ViewLeaguePermission), c.LeagueController.GetLeague)
			leagues.POST("", requirePerm(enums.ManageLeaguePermission), c.LeagueController.CreateLeague)
			leagues.PATCH(":leagueId", requirePerm(enums.ManageLeaguePermission), c.LeagueController.UpdateLeague)
			leagues.DELETE(":leagueId", requirePerm(enums.ManageLeaguePermission), c.LeagueController.DeleteLeague)
		}

		results := v1.Group("/results")
		{
			results.GET("", requirePerm(enums.ViewResultPermission), c.ResultController.GetAllResults)
			results.GET(":resultId", requirePerm(enums.ViewResultPermission), c.ResultController.GetResult)
			results.POST("", requirePerm(enums.ManageResultPermission), c.ResultController.CreateResult)
			results.PATCH(":resultId", requirePerm(enums.ManageResultPermission), c.ResultController.UpdateResult)
			results.DELETE(":resultId", requirePerm(enums.ManageResultPermission), c.ResultController.DeleteResult)
		}
	}
}

func requirePerm(permission enums.Permission) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := middlewares.AuthMiddleware(c, authService, userService)
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
