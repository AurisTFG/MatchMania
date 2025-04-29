package middlewares

import (
	"MatchManiaAPI/constants"
	"MatchManiaAPI/services"
	"MatchManiaAPI/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context, authService services.AuthService, userService services.UserService) error {
	accessToken, err := c.Cookie(constants.AccessTokenName)
	if err != nil || accessToken == "" {
		return err
	}

	user, err := authService.VerifyAccessToken(accessToken)
	if err != nil {
		return err
	}

	utils.SetRequestingUserId(c, user.Id)
	permissions, err := userService.GetDistinctPermissionsByUserId(user.Id)
	if err != nil {
		return err
	}

	utils.SetRequestingUserPermissions(c, permissions)

	return nil
}
