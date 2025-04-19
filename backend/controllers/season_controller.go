package controllers

import (
	requests "MatchManiaAPI/models/dtos/requests/seasons"
	responses "MatchManiaAPI/models/dtos/responses/seasons"
	"MatchManiaAPI/models/enums"
	"MatchManiaAPI/services"
	"MatchManiaAPI/utils"
	r "MatchManiaAPI/utils/httpResponses"
	"MatchManiaAPI/validators"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
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

	var seasonsDto []responses.SeasonDto
	copier.Copy(&seasonsDto, seasons)

	r.OK(ctx, seasonsDto)
}

// @Summary Get a season
// @Description Get a season
// @Tags seasons
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID" default(2)
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

	var seasonDto responses.SeasonDto
	copier.Copy(&seasonDto, season)

	r.OK(ctx, seasonDto)
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

	user := utils.GetAuthUser(ctx)
	if user == nil {
		r.Unauthorized(ctx, "User not found")
		return
	}

	_, err := c.seasonService.CreateSeason(&bodyDto, user.Id)
	if err != nil {
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
// @Param seasonId path string true "Season ID" default(1)
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

	user := utils.GetAuthUser(ctx)
	if user == nil {
		r.Unauthorized(ctx, "User not found")
		return
	}

	currentSeason, err := c.seasonService.GetSeasonById(seasonId)
	if err != nil {
		r.NotFound(ctx, "Season not found")
		return
	}

	if user.Role != enums.AdminRole && user.Role != enums.ModeratorRole && user.Id != currentSeason.UserId {
		r.Forbidden(ctx, "This action is forbidden")
		return
	}

	_, err = c.seasonService.UpdateSeason(currentSeason, &bodyDto)
	if err != nil {
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
// @Param seasonId path string true "Season ID" default(1)
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

	user := utils.GetAuthUser(ctx)
	if user == nil {
		r.Unauthorized(ctx, "User not found")
		return
	}

	season, err := c.seasonService.GetSeasonById(seasonId)
	if err != nil {
		r.NotFound(ctx, "Season not found")
		return
	}

	if user.Role != enums.AdminRole && user.Id != season.UserId {
		r.Forbidden(ctx, "This action is forbidden")
		return
	}

	err = c.seasonService.DeleteSeason(season)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.NoContent(ctx)
}
