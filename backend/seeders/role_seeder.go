package seeders

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/models"
	"MatchManiaAPI/models/enums"
	"fmt"
)

func SeedRoles(db *config.DB, env *config.Env) error {
	var count int64
	if err := db.Model(&models.Role{}).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	var allPermissions []models.Permission
	if err := db.Find(&allPermissions).Error; err != nil {
		return err
	}

	roles := make([]models.Role, 0, len(enums.AllRoles()))
	for roleName, roleDescription := range enums.AllRoles() {
		role := models.Role{
			Name:        roleName,
			Description: roleDescription,
		}

		permissionNames := enums.GetPermissionsForRole(enums.Role(roleName))
		permissions := getPermissionsByPermissionNames(allPermissions, permissionNames)

		fmt.Printf("Role: %s, Permissions: %v\n", role.Name, permissions)
		role.Permissions = permissions

		roles = append(roles, role)
	}

	if err := db.Create(&roles).Error; err != nil {
		return err
	}

	return nil
}

func getPermissionsByPermissionNames(allPermissions []models.Permission, permissionNames []string) []models.Permission {
	permissions := []models.Permission{}

	for _, permissionName := range permissionNames {
		for _, permission := range allPermissions {
			if permission.Name == permissionName {
				permissions = append(permissions, permission)
				break
			}
		}
	}

	return permissions
}
