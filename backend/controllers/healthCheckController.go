package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	response := gin.H{
		"status":  "success",
		"message": "API is running",
	}

	c.JSON(http.StatusOK, response)
}
