package controllers

import (
	"MatchManiaAPI/initializers"
	"MatchManiaAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTeam(c *gin.Context) {
	var team models.Team
	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Create(&team)

	c.JSON(http.StatusCreated, gin.H{"data": team})
}

func GetTeam(c *gin.Context) {
	var team models.Team
	if err := initializers.DB.First(&team, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": team})
}

func GetAllTeams(c *gin.Context) {
	var teams []models.Team
	initializers.DB.Find(&teams)

	c.JSON(http.StatusOK, gin.H{"data": teams})
}

func UpdateTeam(c *gin.Context) {
	var team models.Team
	if err := initializers.DB.First(&team, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&team); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&team)

	c.JSON(http.StatusOK, gin.H{"data": team})
}

func DeleteTeam(c *gin.Context) {
	var team models.Team
	if err := initializers.DB.First(&team, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	initializers.DB.Delete(&team)

	c.JSON(http.StatusOK, gin.H{"data": team})
}
