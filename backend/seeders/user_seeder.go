package seeders

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
	"MatchManiaAPI/models/enums"

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

	var allRoles []models.Role
	if err := db.Find(&allRoles).Error; err != nil {
		return err
	}

	users := []models.User{
		{
			Username: "BigBoss",
			Email:    env.AdminEmail,
			Password: env.AdminPassword,
			Country:  "lt",
			Roles:    getRolesByRole(allRoles, enums.AdminRole),
		},
		{
			Username: "SmallBoss",
			Email:    env.ModeratorEmail,
			Password: env.ModeratorPassword,
			Roles:    getRolesByRole(allRoles, enums.ModeratorRole),
		},
		{
			Username:       "Ziren",
			Email:          "TrackmaniaPlayer1@example.com",
			Password:       env.UserPassword,
			Country:        "lt",
			TrackmaniaId:   uuid.MustParse("99a9530f-63a3-4912-99d9-a770feedb989"),
			TrackmaniaName: "Ziren-",
			Roles:          getRolesByRole(allRoles, enums.TrackmaniaPlayerRole),
		},
		{
			Username:       "Bits",
			Email:          "TrackmaniaPlayer2@example.com",
			Password:       env.UserPassword,
			Country:        "lt",
			TrackmaniaId:   uuid.MustParse("23ae2d9d-ae19-4060-9901-19e19a8f6690"),
			TrackmaniaName: "B1tsy",
			Roles:          getRolesByRole(allRoles, enums.TrackmaniaPlayerRole),
		},
		{
			Username:       "Fain",
			Email:          "TrackmaniaPlayer3@example.com",
			Password:       env.UserPassword,
			Country:        "lt",
			TrackmaniaId:   uuid.MustParse("95fb5a24-532a-42cb-b6fb-0a60a857ff62"),
			TrackmaniaName: "faiN91",
			Roles:          getRolesByRole(allRoles, enums.TrackmaniaPlayerRole),
		},
		{
			Username: "CasualGamer",
			Email:    env.UserEmail,
			Password: env.UserPassword,
			Roles:    getRolesByRole(allRoles, enums.UserRole),
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
