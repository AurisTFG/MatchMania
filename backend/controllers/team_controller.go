package controllers

import (
	requests "MatchManiaAPI/models/dtos/requests/teams"
	responses "MatchManiaAPI/models/dtos/responses/teams"
	"MatchManiaAPI/models/enums"
	"MatchManiaAPI/services"
	"MatchManiaAPI/utils"
	r "MatchManiaAPI/utils/httpResponses"
	"MatchManiaAPI/validators"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type TeamController struct {
	seasonService services.SeasonService
	teamService   services.TeamService
}

func NewTeamController(seasonService services.SeasonService, teamService services.TeamService) TeamController {
	return TeamController{seasonService: seasonService, teamService: teamService}
}

// @Summary Get all teams
// @Description Get all teams
// @Tags teams
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID" default(1)
// @Success 200 {object} []responses.TeamDto
// @Failure 400 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /seasons/{seasonId}/teams [get]
func (c *TeamController) GetAllTeams(ctx *gin.Context) {
	seasonId, err := utils.GetParamId(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	teams, err := c.teamService.GetAllTeams(seasonId)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	var teamsDto []responses.TeamDto
	copier.Copy(&teamsDto, teams)

	r.OK(ctx, teamsDto)
}

// @Summary Get a team
// @Description Get a team
// @Tags teams
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID" default(1)
// @Param teamId path string true "Team ID" default(1)
// @Success 200 {object} responses.TeamDto
// @Failure 400 {object} responses.ErrorDto
// @Failure 404 {object} responses.ErrorDto
// @Router /seasons/{seasonId}/teams/{teamId} [get]
func (c *TeamController) GetTeam(ctx *gin.Context) {
	seasonId, err := utils.GetParamId(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	teamId, err := utils.GetParamId(ctx, "teamId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	team, err := c.teamService.GetTeamById(seasonId, teamId)
	if err != nil {
		r.NotFound(ctx, "Team not found in season")
		return
	}

	var teamDto responses.TeamDto
	copier.Copy(&teamDto, team)

	r.OK(ctx, teamDto)
}

// @Summary Create a team
// @Description Create a team
// @Tags teams
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID" default(1)
// @Param team body requests.CreateTeamDto true "Team object that needs to be created"
// @Success 204
// @Failure 400 {object} responses.ErrorDto
// @Failure 401 {object} responses.ErrorDto
// @Failure 404 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /seasons/{seasonId}/teams [post]
func (c *TeamController) CreateTeam(ctx *gin.Context) {
	seasonId, err := utils.GetParamId(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	var bodyDto requests.CreateTeamDto
	if err = ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err = validators.Validate(&bodyDto); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	season, err := c.seasonService.GetSeasonById(seasonId)
	if err != nil {
		r.NotFound(ctx, "Season not found")
		return
	}

	user := utils.GetAuthUser(ctx)
	if user == nil {
		r.Unauthorized(ctx, "User not found")
		return
	}

	_, err = c.teamService.CreateTeam(&bodyDto, season.Id, user.Id)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.NoContent(ctx)
}

// @Summary Update a team
// @Description Update a team
// @Tags teams
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID" default(1)
// @Param teamId path string true "Team ID" default(1)
// @Param team body requests.UpdateTeamDto true "Team object that needs to be updated"
// @Success 204
// @Failure 400 {object} responses.ErrorDto
// @Failure 401 {object} responses.ErrorDto
// @Failure 403 {object} responses.ErrorDto
// @Failure 404 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /seasons/{seasonId}/teams/{teamId} [patch]
func (c *TeamController) UpdateTeam(ctx *gin.Context) {
	seasonId, err := utils.GetParamId(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

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

	user := utils.GetAuthUser(ctx)
	if user == nil {
		r.Unauthorized(ctx, "User not found")
		return
	}

	currentTeam, err := c.teamService.GetTeamById(seasonId, teamId)
	if err != nil {
		r.NotFound(ctx, "Team not found in season")
		return
	}

	if user.Role != enums.AdminRole && user.Role != enums.ModeratorRole && user.Id != currentTeam.UserId {
		r.Forbidden(ctx, "This action is forbidden")
		return
	}

	_, err = c.teamService.UpdateTeam(currentTeam, &bodyDto)
	if err != nil {
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
// @Param seasonId path string true "Season ID" default(1)
// @Param teamId path string true "Team ID" default(1)
// @Success 204
// @Failure 400 {object} responses.ErrorDto
// @Failure 401 {object} responses.ErrorDto
// @Failure 403 {object} responses.ErrorDto
// @Failure 404 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /seasons/{seasonId}/teams/{teamId} [delete]
func (c *TeamController) DeleteTeam(ctx *gin.Context) {
	seasonId, err := utils.GetParamId(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	teamId, err := utils.GetParamId(ctx, "teamId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	user := utils.GetAuthUser(ctx)
	if user == nil {
		r.Unauthorized(ctx, "User not found")
		return
	}

	team, err := c.teamService.GetTeamById(seasonId, teamId)
	if err != nil {
		r.NotFound(ctx, "Team not found in season")
		return
	}

	if user.Role != enums.AdminRole && user.Id != team.UserId {
		r.Forbidden(ctx, "This action is forbidden")
		return
	}

	err = c.teamService.DeleteTeam(team)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.NoContent(ctx)
}
