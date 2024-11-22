package main

import (
	"MatchManiaAPI/initializers"
	"MatchManiaAPI/routes"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	if err := initializers.LoadEnvVars(); err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}
	fmt.Println("(1/4) Environment variables successfully loaded")

	if err := initializers.ConnectToDatabase(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("(2/4) Successfully connected to database")

	if err := initializers.SyncDatabase(); err != nil {
		log.Fatalf("Failed to sync database: %v", err)
	}
	fmt.Println("(3/4) Successfully synced database")

	if err := initializers.SeedDatabase(); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}
	fmt.Println("(4/4) Successfully seeded database")
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
