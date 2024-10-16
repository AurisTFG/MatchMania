package main

import (
	"MatchManiaAPI/initializers"
	"MatchManiaAPI/models"
	"log"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
}

func main() {
	err := initializers.DB.AutoMigrate(
		&models.Season{},
		&models.Team{},
		&models.Result{},
	)
	if err != nil {
		log.Fatal("failed to migrate database", err)
	}

	log.Println("Migration complete")
}
