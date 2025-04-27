package middlewares

import (
	"MatchManiaAPI/constants"
	"MatchManiaAPI/services"
	"MatchManiaAPI/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context, authService services.AuthService) error {
	accessToken, err := c.Cookie(constants.AccessTokenName)
	if err != nil || accessToken == "" {
		return err
	}

	user, err := authService.VerifyAccessToken(accessToken)
	if err != nil {
		return err
	}

	utils.SetRequestingUserId(c, user.Id)

	var userPermissions []string
	for _, role := range user.Roles {
		for _, p := range role.Permissions {
			userPermissions = append(userPermissions, p.Name)
		}
	}

	utils.SetRequestingUserPermissions(c, userPermissions)

	return nil
}
