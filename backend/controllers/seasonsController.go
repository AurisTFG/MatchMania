package controllers

import (
	"MatchManiaAPI/models"
	r "MatchManiaAPI/responses"
	"MatchManiaAPI/services"

	"github.com/gin-gonic/gin"
)

// @Summary Create a season
// @Description Create a season
// @Tags seasons
// @Accept json
// @Produce json
// @Param season body models.CreateSeasonDto true "Season object that needs to be created"
// @Success 201 {object} models.SeasonResponse
// @Failure 400 {object} models.BadRequestResponse
// @Failure 422 {object} models.UnprocessableEntityResponse
// @Failure 502 {object} models.BadGatewayResponse
// @Router /seasons [post]
func CreateSeason(c *gin.Context) {
	var bodyDto models.CreateSeasonDto

	if err := c.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(c, err.Error())
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(c, err.Error())
		return
	}

	newSeason, err := services.CreateSeason(&bodyDto)
	if err != nil {
		r.BadGateway(c, err.Error())
		return
	}

	r.Created(c, models.SeasonResponse{Season: newSeason.ToDto()})
}

// @Summary Update a season
// @Description Update a season
// @Tags seasons
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID"
// @Param season body models.UpdateSeasonDto true "Season object that needs to be updated"
// @Success 200 {object} models.SeasonResponse
// @Failure 400 {object} models.BadRequestResponse
// @Failure 404 {object} models.NotFoundResponse
// @Failure 502 {object} models.BadGatewayResponse
// @Router /seasons/{seasonId} [put]
func UpdateSeason(c *gin.Context) {
	id := c.Param("seasonId")
	var bodyDto models.UpdateSeasonDto

	if err := c.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(c, err.Error())
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(c, err.Error())
		return
	}

	season, err := services.GetSeasonByID(id)
	if err != nil {
		r.NotFound(c, "Season with id "+id+" not found")
		return
	}

	updatedSeason, err := services.UpdateSeason(season, &bodyDto)
	if err != nil {
		r.BadGateway(c, err.Error())
		return
	}

	r.OK(c, models.SeasonResponse{Season: updatedSeason.ToDto()})
}

// @Summary Get a season
// @Description Get a season
// @Tags seasons
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID"
// @Success 200 {object} models.SeasonResponse
// @Failure 404 {object} models.NotFoundResponse
// @Router /seasons/{seasonId} [get]
func GetSeason(c *gin.Context) {
	id := c.Param("seasonId")

	season, err := services.GetSeasonByID(id)
	if err != nil {
		r.NotFound(c, "Season with id "+id+" not found")
		return
	}

	r.OK(c, models.SeasonResponse{Season: season.ToDto()})
}

// @Summary Get all seasons
// @Description Get all seasons
// @Tags seasons
// @Accept json
// @Produce json
// @Success 200 {object} models.SeasonsResponse
// @Failure 502 {object} models.BadGatewayResponse
// @Router /seasons [get]
func GetAllSeasons(c *gin.Context) {
	seasons, err := services.GetAllSeasons()
	if err != nil {
		r.BadGateway(c, err.Error())
		return
	}

	r.OK(c, models.SeasonsResponse{Seasons: models.ToSeasonDtos(seasons)})
}

// @Summary Delete a season
// @Description Delete a season
// @Tags seasons
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID"
// @Success 204
// @Failure 404 {object} models.NotFoundResponse
// @Failure 502 {object} models.BadGatewayResponse
// @Router /seasons/{seasonId} [delete]
func DeleteSeason(c *gin.Context) {
	id := c.Param("seasonId")

	season, err := services.GetSeasonByID(id)
	if err != nil {
		r.NotFound(c, "Season with id "+id+" not found")
		return
	}

	err = services.DeleteSeason(season)
	if err != nil {
		r.BadGateway(c, err.Error())
		return
	}

	r.Deleted(c)
}
