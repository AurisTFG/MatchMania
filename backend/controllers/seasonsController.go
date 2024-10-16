package controllers

import (
	"MatchManiaAPI/initializers"
	"MatchManiaAPI/models"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateSeason(c *gin.Context) {
	season := models.Season{Name: "2021-2022", StartDate: time.Now(), EndDate: time.Now().AddDate(1, 0, 0)}

	result := initializers.DB.Create(&season)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"season": season,
	})
}

func GetSeason(c *gin.Context) {
	season := models.Season{}
	initializers.DB.First(&season, c.Param("id"))

	c.JSON(200, gin.H{
		"season": season,
	})
}

func GetAllSeasons(c *gin.Context) {
	seasons := []models.Season{}
	initializers.DB.Find(&seasons)

	c.JSON(200, gin.H{
		"seasons": seasons,
	})
}

func UpdateSeason(c *gin.Context) {
	season := models.Season{}
	initializers.DB.First(&season, c.Param("id"))

	season.Name = "2022-2023"
	initializers.DB.Save(&season)

	c.JSON(200, gin.H{
		"season": season,
	})
}

func DeleteSeason(c *gin.Context) {
	season := models.Season{}
	initializers.DB.First(&season, c.Param("id"))

	initializers.DB.Delete(&season)

	c.JSON(200, gin.H{
		"season": season,
	})
}
