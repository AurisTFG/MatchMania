package initializers

import (
	"MatchManiaAPI/models"
	"fmt"
)

func SyncDatabase() error {
	err := DB.AutoMigrate(
		&models.User{},
		&models.Season{},
		&models.Team{},
		&models.Result{},
	)
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}
