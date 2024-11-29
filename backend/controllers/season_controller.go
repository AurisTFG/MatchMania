package controllers

import (
	"MatchManiaAPI/models"
	"MatchManiaAPI/responses"
	r "MatchManiaAPI/responses"
	"MatchManiaAPI/services"
	"MatchManiaAPI/utils"

	"github.com/gin-gonic/gin"
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
// @Success 200 {object} models.SeasonsResponse
// @Failure 502 {object} models.BadGatewayResponse
// @Router /seasons [get]
func (c *SeasonController) GetAllSeasons(ctx *gin.Context) {
	seasons, err := c.seasonService.GetAllSeasons()
	if err != nil {
		r.BadGateway(ctx, err.Error())
		return
	}

	r.OK(ctx, responses.SeasonsResponse{Seasons: models.ToSeasonDtos(seasons)})
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
func (c *SeasonController) GetSeason(ctx *gin.Context) {
	seasonID, err := utils.GetParamUint(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	season, err := c.seasonService.GetSeasonByID(seasonID)
	if err != nil {
		r.NotFound(ctx, "Season not found")
		return
	}

	r.OK(ctx, responses.SeasonResponse{Season: season.ToDto()})
}

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
func (c *SeasonController) CreateSeason(ctx *gin.Context) {
	var bodyDto models.CreateSeasonDto
	if err := ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	user := utils.GetAuthUser(ctx)
	if user == nil {
		r.Unauthorized(ctx, "User not found")
		return
	}

	newSeason, err := c.seasonService.CreateSeason(&bodyDto, user.UUID)
	if err != nil {
		r.BadGateway(ctx, err.Error())
		return
	}

	r.Created(ctx, responses.SeasonResponse{Season: newSeason.ToDto()})
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
// @Router /seasons/{seasonId} [patch]
func (c *SeasonController) UpdateSeason(ctx *gin.Context) {
	seasonID, err := utils.GetParamUint(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	var bodyDto models.UpdateSeasonDto
	if err := ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	user := utils.GetAuthUser(ctx)
	if user == nil {
		r.Unauthorized(ctx, "User not found")
		return
	}

	season, err := c.seasonService.GetSeasonByID(seasonID)
	if err != nil {
		r.NotFound(ctx, "Season not found")
		return
	}

	if user.Role != models.AdminRole && user.UUID != season.UserUUID {
		r.Forbidden(ctx, "This action is forbidden")
		return
	}

	updatedSeason, err := c.seasonService.UpdateSeason(seasonID, &bodyDto)
	if err != nil {
		r.BadGateway(ctx, err.Error())
		return
	}

	r.OK(ctx, responses.SeasonResponse{Season: updatedSeason.ToDto()})
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
func (c *SeasonController) DeleteSeason(ctx *gin.Context) {
	seasonID, err := utils.GetParamUint(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	user := utils.GetAuthUser(ctx)
	if user == nil {
		r.Unauthorized(ctx, "User not found")
		return
	}

	season, err := c.seasonService.GetSeasonByID(seasonID)
	if err != nil {
		r.NotFound(ctx, "Season not found")
		return
	}

	if user.Role != models.AdminRole && user.UUID != season.UserUUID {
		r.Forbidden(ctx, "This action is forbidden")
		return
	}

	err = c.seasonService.DeleteSeason(season)
	if err != nil {
		r.BadGateway(ctx, err.Error())
		return
	}

	r.Deleted(ctx)
}
