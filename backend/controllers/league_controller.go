package controllers

import (
	"MatchManiaAPI/models/dtos/requests"
	"MatchManiaAPI/models/dtos/responses"
	"MatchManiaAPI/services"
	"MatchManiaAPI/utils"
	r "MatchManiaAPI/utils/httpresponses"
	"MatchManiaAPI/validators"

	"github.com/gin-gonic/gin"
)

type LeagueController struct {
	leagueService services.LeagueService
}

func NewLeagueController(leagueService services.LeagueService) LeagueController {
	return LeagueController{leagueService: leagueService}
}

// @Summary Get all leagues
// @Description Get all leagues
// @Tags leagues
// @Accept json
// @Produce json
// @Success 200 {object} []responses.LeagueDto
// @Router /leagues [get]
func (c *LeagueController) GetAllLeagues(ctx *gin.Context) {
	leagues, err := c.leagueService.GetAllLeagues()
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	dto := utils.MustCopy[[]responses.LeagueDto](leagues)

	r.OK(ctx, dto)
}

// @Summary Get a league
// @Description Get a league
// @Tags leagues
// @Accept json
// @Produce json
// @Param leagueId path string true "League ID" default(0deecf6a-289b-49a0-8f1b-9bc4185f99df)
// @Success 200 {object} responses.LeagueDto
// @Failure 400 {object} responses.ErrorDto
// @Failure 404 {object} responses.ErrorDto
// @Router /leagues/{leagueId} [get]
func (c *LeagueController) GetLeague(ctx *gin.Context) {
	leagueId, err := utils.GetParamId(ctx, "leagueId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	league, err := c.leagueService.GetLeagueById(leagueId)
	if err != nil {
		r.NotFound(ctx, "League not found")
		return
	}

	dto := utils.MustCopy[responses.LeagueDto](league)

	r.OK(ctx, dto)
}

// @Summary Create a league
// @Description Create a league
// @Tags leagues
// @Accept json
// @Produce json
// @Param league body requests.CreateLeagueDto true "League object that needs to be created"
// @Success 201
// @Failure 400 {object} responses.ErrorDto
// @Failure 401 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /leagues [post]
func (c *LeagueController) CreateLeague(ctx *gin.Context) {
	var bodyDto requests.CreateLeagueDto
	if err := ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err := validators.Validate(&bodyDto); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	userId := utils.MustGetRequestingUserId(ctx)

	if err := c.leagueService.CreateLeague(&bodyDto, userId); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.Created(ctx)
}

// @Summary Update a league
// @Description Update a league
// @Tags leagues
// @Accept json
// @Produce json
// @Param leagueId path string true "League ID" default(0deecf6a-289b-49a0-8f1b-9bc4185f99df)
// @Param league body requests.UpdateLeagueDto true "League object that needs to be updated"
// @Success 204
// @Failure 400 {object} responses.ErrorDto
// @Failure 401 {object} responses.ErrorDto
// @Failure 404 {object} responses.ErrorDto
// @Failure 403 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /leagues/{leagueId} [patch]
func (c *LeagueController) UpdateLeague(ctx *gin.Context) {
	leagueId, err := utils.GetParamId(ctx, "leagueId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	var bodyDto requests.UpdateLeagueDto
	if err = ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err = validators.Validate(&bodyDto); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	currentLeague, err := c.leagueService.GetLeagueById(leagueId)
	if err != nil {
		r.NotFound(ctx, "League not found")
		return
	}

	if err = c.leagueService.UpdateLeague(currentLeague, &bodyDto); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.NoContent(ctx)
}

// @Summary Delete a league
// @Description Delete a league
// @Tags leagues
// @Accept json
// @Produce json
// @Param leagueId path string true "League ID" default(0deecf6a-289b-49a0-8f1b-9bc4185f99df)
// @Success 204
// @Failure 400 {object} responses.ErrorDto
// @Failure 401 {object} responses.ErrorDto
// @Failure 403 {object} responses.ErrorDto
// @Failure 404 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /leagues/{leagueId} [delete]
func (c *LeagueController) DeleteLeague(ctx *gin.Context) {
	leagueId, err := utils.GetParamId(ctx, "leagueId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	league, err := c.leagueService.GetLeagueById(leagueId)
	if err != nil {
		r.NotFound(ctx, "League not found")
		return
	}

	err = c.leagueService.DeleteLeague(league)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.NoContent(ctx)
}
