package routes

import (
	"MatchManiaAPI/controllers"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(server *gin.Engine) {
	v1 := server.Group("/api/v1")
	{
		v1.GET("/health-check", controllers.HealthCheck)

		seasons := v1.Group("/seasons")
		{
			seasons.GET("/", controllers.GetAllSeasons)
			seasons.POST("/", controllers.CreateSeason)
			seasons.GET("/:seasonId", controllers.GetSeason)
			seasons.PUT("/:seasonId", controllers.UpdateSeason)
			seasons.DELETE("/:seasonId", controllers.DeleteSeason)

			teams := seasons.Group("/:seasonId/teams")
			{
				teams.GET("/", controllers.GetAllTeams)
				teams.POST("/", controllers.CreateTeam)
				teams.GET("/:teamId", controllers.GetTeam)
				teams.PUT("/:teamId", controllers.UpdateTeam)
				teams.DELETE("/:teamId", controllers.DeleteTeam)

				results := teams.Group("/:teamId/results")
				{
					results.GET("/", controllers.GetAllResults)
					results.POST("/", controllers.CreateResult)
					results.GET("/:resultId", controllers.GetResult)
					results.PUT("/:resultId", controllers.UpdateResult)
					results.DELETE("/:resultId", controllers.DeleteResult)
				}
			}
		}
	}
}
