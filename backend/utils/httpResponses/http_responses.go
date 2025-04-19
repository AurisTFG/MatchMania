package httpResponses

import (
	dtos "MatchManiaAPI/models/dtos/responses/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, obj any) { // 200
	c.JSON(http.StatusOK, obj)
}

func Created(c *gin.Context, obj any) { // 201
	c.JSON(http.StatusCreated, obj)
}

func NoContent(c *gin.Context) { // 204
	c.JSON(http.StatusNoContent, nil)
}

func BadRequest(c *gin.Context, errorMessage string) { // 400
	c.JSON(http.StatusBadRequest, dtos.ErrorDto{Message: errorMessage})
}

func Unauthorized(c *gin.Context) { // 401
	c.JSON(http.StatusUnauthorized, dtos.ErrorDto{Message: "Unauthorized"})
	c.Abort()
}

func Forbidden(c *gin.Context, errorMessage string) { // 403
	c.JSON(http.StatusForbidden, dtos.ErrorDto{Message: errorMessage})
}

func NotFound(c *gin.Context, errorMessage string) { // 404
	c.JSON(http.StatusNotFound, dtos.ErrorDto{Message: errorMessage})
}

func UnprocessableEntity(c *gin.Context, errorMessage string) { // 422
	c.JSON(http.StatusUnprocessableEntity, dtos.ErrorDto{Message: errorMessage})
}

func InternalServerError(c *gin.Context, errorMessage string) { // 500
	c.JSON(http.StatusInternalServerError, dtos.ErrorDto{Message: errorMessage})
}
