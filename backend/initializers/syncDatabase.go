package initializers

import (
	"MatchManiaAPI/models"
	"log"
)

func SyncDatabase() {
	err := DB.AutoMigrate(
		&models.Season{},
		&models.Team{},
		&models.Result{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}

	log.Println("Migration complete")
}
