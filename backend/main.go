package main

import (
	"MatchManiaAPI/initializers"
	"MatchManiaAPI/routes"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.SyncDatabase()
	initializers.SeedDatabase()
}

// @title MatchMania API
// @version 1.0
// @description This is an API for managing matchmaking seasons, teams, and results
// @host localhost:8000
// @BasePath /api/v1
func main() {
	gin.SetMode(gin.DebugMode)

	server := gin.Default()

	server.SetTrustedProxies(nil)
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{initializers.Cfg.ClientURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.ApplyRoutes(server)

	err := server.Run(":" + initializers.Cfg.ServerPort)
	if err != nil {
		log.Fatal("Failed to start Gin server.")
	}
}
