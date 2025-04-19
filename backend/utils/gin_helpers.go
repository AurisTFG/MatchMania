package utils

import (
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

func GetAuthUserId(ctx *gin.Context) uuid.UUID {
	userId, ok := ctx.Get("userId")
	if !ok {
		return uuid.Nil
	}

	userIdObj, ok := userId.(uuid.UUID)
	if !ok {
		return uuid.Nil
	}

	return userIdObj
}
