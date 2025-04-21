package controllers

import (
	"MatchManiaAPI/models/dtos/requests"
	"MatchManiaAPI/models/dtos/responses"
	"MatchManiaAPI/services"
	"MatchManiaAPI/utils"
	r "MatchManiaAPI/utils/httpresponses"
	"MatchManiaAPI/validators"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SeasonController struct {
	seasonService services.SeasonService
}

func NewSeasonController(seasonService services.SeasonService) SeasonController {
	return SeasonController{seasonService: seasonService}
}

// @Summary Get all seasons
// @Description Get all seasons
// @Tags seasons
// @Accept json
// @Produce json
// @Success 200 {object} []responses.SeasonDto
// @Router /seasons [get]
func (c *SeasonController) GetAllSeasons(ctx *gin.Context) {
	seasons, err := c.seasonService.GetAllSeasons()
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	dto := utils.CopyOrPanic[[]responses.SeasonDto](seasons)

	r.OK(ctx, dto)
}

// @Summary Get a season
// @Description Get a season
// @Tags seasons
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID" default(0deecf6a-289b-49a0-8f1b-9bc4185f99df)
// @Success 200 {object} responses.SeasonDto
// @Failure 400 {object} responses.ErrorDto
// @Failure 404 {object} responses.ErrorDto
// @Router /seasons/{seasonId} [get]
func (c *SeasonController) GetSeason(ctx *gin.Context) {
	seasonId, err := utils.GetParamId(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	season, err := c.seasonService.GetSeasonById(seasonId)
	if err != nil {
		r.NotFound(ctx, "Season not found")
		return
	}

	dto := utils.CopyOrPanic[responses.SeasonDto](season)

	r.OK(ctx, dto)
}

// @Summary Create a season
// @Description Create a season
// @Tags seasons
// @Accept json
// @Produce json
// @Param season body requests.CreateSeasonDto true "Season object that needs to be created"
// @Success 204
// @Failure 400 {object} responses.ErrorDto
// @Failure 401 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /seasons [post]
func (c *SeasonController) CreateSeason(ctx *gin.Context) {
	var bodyDto requests.CreateSeasonDto
	if err := ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err := validators.Validate(&bodyDto); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	userId := utils.GetAuthUserId(ctx)
	if userId == uuid.Nil {
		r.Unauthorized(ctx)
		return
	}

	if err := c.seasonService.CreateSeason(&bodyDto, userId); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.NoContent(ctx)
}

// @Summary Update a season
// @Description Update a season
// @Tags seasons
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID" default(0deecf6a-289b-49a0-8f1b-9bc4185f99df)
// @Param season body requests.UpdateSeasonDto true "Season object that needs to be updated"
// @Success 204
// @Failure 400 {object} responses.ErrorDto
// @Failure 401 {object} responses.ErrorDto
// @Failure 404 {object} responses.ErrorDto
// @Failure 403 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /seasons/{seasonId} [patch]
func (c *SeasonController) UpdateSeason(ctx *gin.Context) {
	seasonId, err := utils.GetParamId(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	var bodyDto requests.UpdateSeasonDto
	if err = ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err = validators.Validate(&bodyDto); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	currentSeason, err := c.seasonService.GetSeasonById(seasonId)
	if err != nil {
		r.NotFound(ctx, "Season not found")
		return
	}

	if err = c.seasonService.UpdateSeason(currentSeason, &bodyDto); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.NoContent(ctx)
}

// @Summary Delete a season
// @Description Delete a season
// @Tags seasons
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID" default(0deecf6a-289b-49a0-8f1b-9bc4185f99df)
// @Success 204
// @Failure 400 {object} responses.ErrorDto
// @Failure 401 {object} responses.ErrorDto
// @Failure 403 {object} responses.ErrorDto
// @Failure 404 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /seasons/{seasonId} [delete]
func (c *SeasonController) DeleteSeason(ctx *gin.Context) {
	seasonId, err := utils.GetParamId(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	season, err := c.seasonService.GetSeasonById(seasonId)
	if err != nil {
		r.NotFound(ctx, "Season not found")
		return
	}

	err = c.seasonService.DeleteSeason(season)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.NoContent(ctx)
}
