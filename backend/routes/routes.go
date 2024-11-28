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
	authMiddleware middlewares.AuthMiddleware,
	authController controllers.AuthController,
	seasonController controllers.SeasonController,
	teamController controllers.TeamController,
	resultController controllers.ResultController,
) {
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := server.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", authController.AuthSignUp)
			auth.POST("/login", authController.AuthLogIn)
			auth.POST("/logout", authController.AuthLogOut)
			auth.POST("/refresh-token", authController.AuthRefreshToken)
		}

		seasons := v1.Group("/seasons")
		{
			seasons.GET("/", seasonController.GetAllSeasons)
			seasons.POST("/", seasonController.CreateSeason)
			seasons.GET("/:seasonId", seasonController.GetSeason)
			seasons.PUT("/:seasonId", seasonController.UpdateSeason)
			seasons.DELETE("/:seasonId", seasonController.DeleteSeason)

			teams := seasons.Group("/:seasonId/teams")
			{
				teams.GET("/", teamController.GetAllTeams)
				teams.POST("/", teamController.CreateTeam)
				teams.GET("/:teamId", teamController.GetTeam)
				teams.PUT("/:teamId", teamController.UpdateTeam)
				teams.DELETE("/:teamId", teamController.DeleteTeam)

				results := teams.Group("/:teamId/results")
				{
					results.GET("/", resultController.GetAllResults)
					results.POST("/", authMiddleware.RequireAuth, resultController.CreateResult)
					results.GET("/:resultId", resultController.GetResult)
					results.PUT("/:resultId", resultController.UpdateResult)
					results.DELETE("/:resultId", resultController.DeleteResult)
				}
			}
		}
	}
}
