package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OK(c *gin.Context, dataName string, data interface{}) {
	c.JSON(
		http.StatusOK,
		gin.H{
			dataName: data,
			"status": "success",
		},
	)
}

func Created(c *gin.Context, dataName string, data interface{}) {
	c.JSON(
		http.StatusCreated,
		gin.H{
			dataName: data,
			"status": "success",
		},
	)
}

func Deleted(c *gin.Context) {
	c.JSON(
		http.StatusNoContent,
		nil,
	)
}

func BadRequest(c *gin.Context, err error) {
	c.JSON(
		http.StatusBadRequest,
		gin.H{
			"status": "error",
			"error":  err.Error(),
		},
	)
}

func NotFound(c *gin.Context, message string) {
	c.JSON(
		http.StatusNotFound,
		gin.H{
			"status": "error",
			"error":  message,
		},
	)
}

func UnprocessableEntity(c *gin.Context, err error) {
	c.JSON(
		http.StatusUnprocessableEntity,
		gin.H{
			"status": "error",
			"error":  err.Error(),
		},
	)
}

func BadGateway(c *gin.Context, err error) {
	c.JSON(
		http.StatusBadGateway,
		gin.H{
			"status": "error",
			"error":  err.Error(),
		},
	)
}
