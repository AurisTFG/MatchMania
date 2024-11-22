package responses

import (
	"MatchManiaAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, obj any) {
	c.JSON(http.StatusOK, obj)
}

func Created(c *gin.Context, obj any) {
	c.JSON(http.StatusCreated, obj)
}

func Deleted(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}

func BadRequest(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusBadRequest, models.BadRequestResponse{Error: errorMessage})
}

func Unauthorized(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusUnauthorized, models.UnauthorizedResponse{Error: errorMessage})
	c.Abort()
}

func NotFound(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusNotFound, models.NotFoundResponse{Error: errorMessage})
}

func UnprocessableEntity(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusUnprocessableEntity, models.UnprocessableEntityResponse{Error: errorMessage})
}

func BadGateway(c *gin.Context, errorMessage string) {
	c.JSON(http.StatusBadGateway, models.BadGatewayResponse{Error: errorMessage})
}
