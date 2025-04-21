package utils

import (
	"MatchManiaAPI/constants"

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

func MustGetRequestingUserId(ctx *gin.Context) uuid.UUID {
	userIdObj, exists := ctx.Get(constants.RequestingUserIdKey)
	if !exists {
		panic("Requesting user ID not found in context")
	}

	userId, ok := userIdObj.(uuid.UUID)
	if !ok {
		panic("Requesting user ID is not of type uuid.UUID")
	}

	return userId
}

func SetRequestingUserId(ctx *gin.Context, userId uuid.UUID) {
	ctx.Set(constants.RequestingUserIdKey, userId)
}
