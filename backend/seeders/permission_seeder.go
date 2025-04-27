package seeders

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
	"MatchManiaAPI/models/enums"
)

func SeedPermissions(db *config.DB, env *config.Env) error {
	var count int64
	if err := db.Model(&models.Permission{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	permissions := make([]models.Permission, 0, len(enums.AllPermissionsWithDesc()))
	for permissionName, permissionDescription := range enums.AllPermissionsWithDesc() {
		permission := models.Permission{
			Name:        permissionName,
			Description: permissionDescription,
		}

		permissions = append(permissions, permission)
	}

	if err := db.Create(&permissions).Error; err != nil {
		return err
	}

	return nil
}
