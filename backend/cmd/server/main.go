package main

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/controllers"
	"MatchManiaAPI/middlewares"
	"MatchManiaAPI/repositories"
	"MatchManiaAPI/routes"
	"MatchManiaAPI/services"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	env *config.Env
	db  *config.DB
	err error
)

func init() {
	fmt.Print("(0/4) ")
	envName := os.Getenv("ENV")
	if envName == "" {
		envName = "dev"
	}
	fmt.Println("Environment:", envName)

	fmt.Print("(1/4) ")
	env, err = config.LoadEnv(envName)
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}
	fmt.Println("Environment variables successfully loaded")

	fmt.Print("(2/4) ")
	db, err = config.ConnectDatabase(env)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Successfully connected to database")

	fmt.Print("(3/4) ")
	err = config.MigrateDatabase(db)
	if err != nil {
		log.Fatalf("Failed to sync database: %v", err)
	}
	fmt.Println("Successfully synced database")

	fmt.Print("(4/4) ")
	err = config.SeedDatabase(db)
	if err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}
	fmt.Println("Successfully seeded database")
}

// @title MatchMania API
// @version 1.0
// @description This is the API server for the MatchMania application.
// @host localhost:8080
// @BasePath /api/v1
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	gin.SetMode(gin.DebugMode)

	server := gin.Default()

	if err := server.SetTrustedProxies(nil); err != nil {
		log.Fatal("Failed to set trusted proxies.")
	}

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{env.ClientURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	userRepository := repositories.NewUserRepository(db)
	seasonRepository := repositories.NewSeasonRepository(db)
	teamRepository := repositories.NewTeamRepository(db)
	resultRepository := repositories.NewResultRepository(db)

	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService(userService, env)
	seasonService := services.NewSeasonService(seasonRepository)
	teamService := services.NewTeamService(teamRepository)
	resultService := services.NewResultService(resultRepository)

	authMiddleware := middlewares.NewAuthMiddleware(authService)

	authController := controllers.NewAuthController(authService, env)
	seasonController := controllers.NewSeasonController(seasonService)
	teamController := controllers.NewTeamController(seasonService, teamService)
	resultController := controllers.NewResultController(teamService, resultService)

	routes.ApplyRoutes(server, authMiddleware, authController, seasonController, teamController, resultController)

	err := server.Run(":" + env.ServerPort)
	if err != nil {
		log.Fatal("Failed to start Gin server.")
	}
}
