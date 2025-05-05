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

type TeamController struct {
	leagueService services.LeagueService
	teamService   services.TeamService
}

func NewTeamController(leagueService services.LeagueService, teamService services.TeamService) TeamController {
	return TeamController{leagueService: leagueService, teamService: teamService}
}

// @Summary Get all teams
// @Description Get all teams
// @Tags teams
// @Accept json
// @Produce json
// @Success 200 {object} []responses.TeamDto
// @Failure 400 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /teams [get]
func (c *TeamController) GetAllTeams(ctx *gin.Context) {
	teams, err := c.teamService.GetAllTeams()
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	dto := utils.MustCopy[[]responses.TeamDto](teams)

	r.OK(ctx, dto)
}

// @Summary Get a team
// @Description Get a team
// @Tags teams
// @Accept json
// @Produce json
// @Param teamId path string true "Team ID" default(0deecf6a-289b-49a0-8f1b-9bc4185f99df)
// @Success 200 {object} responses.TeamDto
// @Failure 400 {object} responses.ErrorDto
// @Failure 404 {object} responses.ErrorDto
// @Router /teams/{teamId} [get]
func (c *TeamController) GetTeam(ctx *gin.Context) {
	teamId, err := utils.GetParamId(ctx, "teamId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	team, err := c.teamService.GetTeamById(teamId)
	if err != nil {
		r.NotFound(ctx, "Team not found in league")
		return
	}

	dto := utils.MustCopy[responses.TeamDto](team)

	r.OK(ctx, dto)
}

// @Summary Create a team
// @Description Create a team
// @Tags teams
// @Accept json
// @Produce json
// @Param team body requests.CreateTeamDto true "Team object that needs to be created"
// @Success 201
// @Failure 400 {object} responses.ErrorDto
// @Failure 401 {object} responses.ErrorDto
// @Failure 404 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /teams [post]
func (c *TeamController) CreateTeam(ctx *gin.Context) {
	var bodyDto requests.CreateTeamDto
	if err := ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err := validators.Validate(&bodyDto); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	userId := utils.MustGetRequestingUserId(ctx)

	if err := c.teamService.CreateTeam(&bodyDto, userId); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.Created(ctx)
}

// @Summary Update a team
// @Description Update a team
// @Tags teams
// @Accept json
// @Produce json
// @Param teamId path string true "Team ID" default(0deecf6a-289b-49a0-8f1b-9bc4185f99df)
// @Param team body requests.UpdateTeamDto true "Team object that needs to be updated"
// @Success 204
// @Failure 400 {object} responses.ErrorDto
// @Failure 401 {object} responses.ErrorDto
// @Failure 403 {object} responses.ErrorDto
// @Failure 404 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /teams/{teamId} [patch]
func (c *TeamController) UpdateTeam(ctx *gin.Context) {
	teamId, err := utils.GetParamId(ctx, "teamId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	var bodyDto requests.UpdateTeamDto

	if err = ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err = validators.Validate(&bodyDto); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	currentTeam, err := c.teamService.GetTeamById(teamId)
	if err != nil {
		r.NotFound(ctx, "Team not found in league")
		return
	}

	if err = c.teamService.UpdateTeam(currentTeam, &bodyDto); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.NoContent(ctx)
}

// @Summary Delete a team
// @Description Delete a team
// @Tags teams
// @Accept json
// @Produce json
// @Param teamId path string true "Team ID" default(0deecf6a-289b-49a0-8f1b-9bc4185f99df)
// @Success 204
// @Failure 400 {object} responses.ErrorDto
// @Failure 401 {object} responses.ErrorDto
// @Failure 403 {object} responses.ErrorDto
// @Failure 404 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /teams/{teamId} [delete]
func (c *TeamController) DeleteTeam(ctx *gin.Context) {
	teamId, err := utils.GetParamId(ctx, "teamId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	team, err := c.teamService.GetTeamById(teamId)
	if err != nil {
		r.NotFound(ctx, "Team not found in league")
		return
	}

	err = c.teamService.DeleteTeam(team)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.NoContent(ctx)
}
