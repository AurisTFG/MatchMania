package utils

import (
	"MatchManiaAPI/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetParamId(ctx *gin.Context, paramName string) (uuid.UUID, error) {
	idStr := ctx.Param(paramName)
	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
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
