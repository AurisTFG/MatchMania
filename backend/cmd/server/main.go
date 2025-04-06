package main

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/controllers"
	"MatchManiaAPI/docs"
	"MatchManiaAPI/middlewares"
	"MatchManiaAPI/routes"
	"fmt"
	"log"
	"os"
)

var (
	env *config.Env
	db  *config.DB
)

func init() {
	var err error

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
// @version 0.1.0
// @description Documentation for MatchMania API
// @host localhost:8080
// @BasePath /api/v1
// @schemes http
// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
// @contact.name AurisTFG
// @contact.url https://github.com/AurisTFG
func main() {
	server, err := config.SetupServer(env)
	if err != nil {
		log.Fatalf("Failed to setup server: %v", err)
	}

	docs.SetupSwagger(server, env)

	controllers := controllers.SetupControllers(db, env)
	middlewares := middlewares.SetupMiddlewares(db, env)
	routes.SetupRoutes(server, controllers, middlewares)

	fmt.Println("(6/6) Starting server on " + env.ServerURL + " . . . ")

	err = server.Run(env.ServerURL)
	if err != nil {
		log.Fatal("Failed to start Gin server.")
	}
}
