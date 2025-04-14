package utils

import (
	"MatchManiaAPI/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParamUint(ctx *gin.Context, paramName string) (uint, error) {
	id := ctx.Param(paramName)
	integer, err := strconv.Atoi(id)

	if err != nil {
		return 0, err
	}

	return uint(integer), nil
}

func GetParamString(ctx *gin.Context, paramName string) string {
	return ctx.Param(paramName)
}

func GetAuthUser(ctx *gin.Context) *models.User {
	user, ok := ctx.Get("user")
	if !ok {
		return nil
	}

	userObj, ok := user.(*models.User)
	if !ok {
		return nil
	}

	return userObj
}
