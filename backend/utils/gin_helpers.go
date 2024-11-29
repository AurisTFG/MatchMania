package utils

import (
	"MatchManiaAPI/models"
	"log"
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

func GetAuthUser(ctx *gin.Context) *models.User {
	user, ok := ctx.Get("user")
	if !ok {
		return nil
	}

	userObj, ok := user.(*models.User)
	if !ok {
		log.Printf("Expected user object, but got: %v", user)
		return nil
	}

	return userObj
}
