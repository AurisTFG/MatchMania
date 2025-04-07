package responses

import (
	"MatchManiaAPI/models"
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
	c.JSON(http.StatusBadRequest, models.ErrorDto{Error: errorMessage})
}

func Unauthorized(c *gin.Context, errorMessage string) { // 401
	c.JSON(http.StatusUnauthorized, models.ErrorDto{Error: errorMessage})
	c.Abort()
}

func Forbidden(c *gin.Context, errorMessage string) { // 403
	c.JSON(http.StatusForbidden, models.ErrorDto{Error: errorMessage})
}

func NotFound(c *gin.Context, errorMessage string) { // 404
	c.JSON(http.StatusNotFound, models.ErrorDto{Error: errorMessage})
}

func UnprocessableEntity(c *gin.Context, errorMessage string) { // 422
	c.JSON(http.StatusUnprocessableEntity, models.ErrorDto{Error: errorMessage})
}

func InternalServerError(c *gin.Context, errorMessage string) { // 500
	c.JSON(http.StatusInternalServerError, models.ErrorDto{Error: errorMessage})
}
