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
	fmt.Print("(1/6) ")
	envName := os.Getenv("MATCHMANIA_ENV")
	if envName == "" {
		log.Fatal("Failed to load environment variables: MATCHMANIA_ENV not set.")
	}
	fmt.Println("Environment:", envName)

	fmt.Print("(2/6) ")
	env, err = config.LoadEnv(envName)
	if err != nil {
		log.Fatalf("Failed to load environment variables: %v", err)
	}
	fmt.Println("Environment variables successfully loaded")

	fmt.Print("(3/6) ")
	db, err = config.ConnectDatabase(env)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Successfully connected to database")

	fmt.Print("(4/6) ")
	err = config.MigrateDatabase(db)
	if err != nil {
		log.Fatalf("Failed to sync database: %v", err)
	}
	fmt.Println("Successfully synced database")

	fmt.Print("(5/6) ")
	err = config.SeedDatabase(db, env)
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
	if env.IsDev {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.Default()

	if err := server.SetTrustedProxies(nil); err != nil {
		log.Fatal("Failed to set trusted proxies.")
	}

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{env.ClientURL},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	sessionRepository := repositories.NewSessionRepository(db)
	userRepository := repositories.NewUserRepository(db)
	seasonRepository := repositories.NewSeasonRepository(db)
	teamRepository := repositories.NewTeamRepository(db)
	resultRepository := repositories.NewResultRepository(db)

	sessionService := services.NewSessionService(sessionRepository, env)
	userService := services.NewUserService(userRepository)
	authService := services.NewAuthService(userService, env)
	seasonService := services.NewSeasonService(seasonRepository)
	teamService := services.NewTeamService(teamRepository)
	resultService := services.NewResultService(resultRepository)

	authMiddleware := middlewares.NewAuthMiddleware(authService)

	authController := controllers.NewAuthController(authService, sessionService, env)
	userController := controllers.NewUserController(userService)
	seasonController := controllers.NewSeasonController(seasonService)
	teamController := controllers.NewTeamController(seasonService, teamService)
	resultController := controllers.NewResultController(teamService, resultService)

	allControllers := routes.NewControllers(authMiddleware, authController, userController, seasonController, teamController, resultController)

	routes.ApplyRoutes(server, &allControllers)

	fmt.Println("(6/6) Starting server on " + env.ServerHost + ":" + env.ServerPort + " . . . ")

	err := server.Run(env.ServerHost + ":" + env.ServerPort)
	if err != nil {
		log.Fatal("Failed to start Gin server.")
	}
}
