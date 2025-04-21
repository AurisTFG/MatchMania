package seeders

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
	"MatchManiaAPI/models/enums"
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
		{Username: "Admin", Email: env.AdminEmail, Password: env.AdminPassword, Role: enums.AdminRole},
		{Username: "Moderator", Email: env.ModeratorEmail, Password: env.ModeratorPassword, Role: enums.ModeratorRole},
		{Username: "User", Email: env.UserEmail, Password: env.UserPassword, Role: enums.UserRole},
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
