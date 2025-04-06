package seeders

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"

	"github.com/google/uuid"
)

func SeedUsers(db *config.DB, env *config.Env) error {
	var count int64
	if err := db.Model(&models.User{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	users := []models.User{
		{
			UUID:     uuid.New(),
			Username: "Admin",
			Email:    env.AdminEmail,
			Password: env.AdminPassword,
			Role:     models.AdminRole,
		},
		{
			UUID:     uuid.New(),
			Username: "Moderator",
			Email:    env.ModeratorEmail,
			Password: env.ModeratorPassword,
			Role:     models.ModeratorRole,
		},
		{
			UUID:     uuid.New(),
			Username: "User",
			Email:    env.UserEmail,
			Password: env.UserPassword,
			Role:     models.UserRole,
		},
	}

	for _, user := range users {
		if err := user.HashPassword(); err != nil {
			return err
		}

		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	return nil
}
