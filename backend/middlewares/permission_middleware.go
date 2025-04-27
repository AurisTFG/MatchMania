package middlewares

import (
	"MatchManiaAPI/utils"

	"errors"
	"slices"

	"github.com/gin-gonic/gin"
)

func RequirePermission(c *gin.Context, requiredPerm string) error {
	userPermissions := utils.GetRequestingUserPermissions(c)
	if userPermissions == nil {
		return errors.New("user permissions not found")
	}

	if !slices.Contains(userPermissions, requiredPerm) {
		return errors.New("user does not have the required permission")
	}

	return nil
}
