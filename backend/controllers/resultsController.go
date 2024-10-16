package controllers

import (
	"MatchManiaAPI/initializers"
	"MatchManiaAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateResult(c *gin.Context) {
	var result models.Result
	if err := c.ShouldBindJSON(&result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Create(&result)

	c.JSON(http.StatusCreated, gin.H{"data": result})
}

func GetResult(c *gin.Context) {
	var result models.Result
	if err := initializers.DB.First(&result, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetAllResults(c *gin.Context) {
	var results []models.Result
	initializers.DB.Find(&results)

	c.JSON(http.StatusOK, gin.H{"data": results})
}

func UpdateResult(c *gin.Context) {
	var result models.Result
	if err := initializers.DB.First(&result, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	if err := c.ShouldBindJSON(&result); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Save(&result)

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func DeleteResult(c *gin.Context) {
	var result models.Result
	if err := initializers.DB.First(&result, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
		return
	}

	initializers.DB.Delete(&result)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
