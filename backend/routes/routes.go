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
			seasons.GET("/:id", controllers.GetSeason)
			seasons.PUT("/:id", controllers.UpdateSeason)
			seasons.DELETE("/:id", controllers.DeleteSeason)
		}

		teams := v1.Group("/teams")
		{
			teams.GET("/", controllers.GetAllTeams)
			teams.POST("/", controllers.CreateTeam)
			teams.GET("/:id", controllers.GetTeam)
			teams.PUT("/:id", controllers.UpdateTeam)
			teams.DELETE("/:id", controllers.DeleteTeam)
		}

		results := v1.Group("/results")
		{
			results.GET("/", controllers.GetAllResults)
			results.POST("/", controllers.CreateResult)
			results.GET("/:id", controllers.GetResult)
			results.PUT("/:id", controllers.UpdateResult)
			results.DELETE("/:id", controllers.DeleteResult)
		}
	}
}
