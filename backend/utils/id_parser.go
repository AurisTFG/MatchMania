package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseID(ctx *gin.Context, paramName string) (uint, error) {
	id := ctx.Param(paramName)
	integer, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return uint(integer), nil
}
