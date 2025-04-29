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

	var allRoles []models.Role
	if err := db.Find(&allRoles).Error; err != nil {
		return err
	}

	users := []models.User{
		{
			Username: "BigBoss",
			Email:    env.AdminEmail,
			Password: env.AdminPassword,
			Roles:    getRolesByRole(allRoles, enums.AdminRole),
		},
		{
			Username: "SmallBoss",
			Email:    env.ModeratorEmail,
			Password: env.ModeratorPassword,
			Roles:    getRolesByRole(allRoles, enums.ModeratorRole),
		},
		{
			Username: "TrackmaniaPro",
			Email:    "TrackmaniaUser@example.com",
			Password: env.UserPassword,
			Roles:    getRolesByRole(allRoles, enums.TrackmaniaUserRole),
		},
		{
			Username: "CasualGamer",
			Email:    env.UserEmail,
			Password: env.UserPassword,
			Roles:    getRolesByRole(allRoles, enums.UserRole),
		},
		{
			Username: "RandomGuest",
			Email:    "Guest@example.com",
			Password: env.UserPassword,
			Roles:    getRolesByRole(allRoles, enums.GuestRole),
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

func getRolesByRole(allRoles []models.Role, role enums.Role) []models.Role {
	var matchedRoles []models.Role

	for _, r := range allRoles {
		if r.Name == string(role) {
			matchedRoles = append(matchedRoles, r)
		}
	}

	return matchedRoles
}
