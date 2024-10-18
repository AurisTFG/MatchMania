package controllers

import (
	"MatchManiaAPI/models"
	r "MatchManiaAPI/responses"
	"MatchManiaAPI/services"

	"github.com/gin-gonic/gin"
)

func CreateSeason(c *gin.Context) {
	var bodyDto models.CreateSeasonDto

	if err := c.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(c, err)
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(c, err)
		return
	}

	newSeason, err := services.CreateSeason(&bodyDto)
	if err != nil {
		r.BadGateway(c, err)
		return
	}

	r.Created(c, "season", newSeason.ToDto())
}

func UpdateSeason(c *gin.Context) {
	id := c.Param("seasonId")
	var bodyDto models.UpdateSeasonDto

	if err := c.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(c, err)
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(c, err)
		return
	}

	season, err := services.GetSeasonByID(id)
	if err != nil {
		r.NotFound(c, "Season with id "+id+" not found")
		return
	}

	updatedSeason, err := services.UpdateSeason(season, &bodyDto)
	if err != nil {
		r.BadGateway(c, err)
		return
	}

	r.OK(c, "season", updatedSeason.ToDto())
}

func GetSeason(c *gin.Context) {
	id := c.Param("seasonId")

	season, err := services.GetSeasonByID(id)
	if err != nil {
		r.NotFound(c, "Season with id "+id+" not found")
		return
	}

	r.OK(c, "season", season.ToDto())
}

func GetAllSeasons(c *gin.Context) {
	seasons, err := services.GetAllSeasons()
	if err != nil {
		r.BadGateway(c, err)
		return
	}

	r.OK(c, "seasons", models.ToSeasonDtos(seasons))
}

func DeleteSeason(c *gin.Context) {
	id := c.Param("seasonId")

	season, err := services.GetSeasonByID(id)
	if err != nil {
		r.NotFound(c, "Season with id "+id+" not found")
		return
	}

	err = services.DeleteSeason(season)
	if err != nil {
		r.BadGateway(c, err)
		return
	}

	r.Deleted(c)
}
