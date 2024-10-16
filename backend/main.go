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
}

func main() {
	server := gin.Default()

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
		log.Fatal("[Error] failed to start Gin server due to: " + err.Error())
	}
}
