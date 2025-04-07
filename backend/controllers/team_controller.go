package controllers

import (
	"MatchManiaAPI/models"
	r "MatchManiaAPI/responses"
	"MatchManiaAPI/services"
	"MatchManiaAPI/utils"

	"github.com/gin-gonic/gin"
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
// @Param seasonId path string true "Season ID" default(2)
// @Success 200 {object} []models.TeamDto
// @Failure 400 {object} models.ErrorDto
// @Failure 422 {object} models.ErrorDto
// @Router /seasons/{seasonId}/teams [get]
func (c *TeamController) GetAllTeams(ctx *gin.Context) {
	seasonID, err := utils.GetParamUint(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	teams, err := c.teamService.GetAllTeams(seasonID)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.OK(ctx, models.ToTeamDtos(teams))
}

// @Summary Get a team
// @Description Get a team
// @Tags teams
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID" default(2)
// @Param teamId path string true "Team ID" default(2)
// @Success 200 {object} models.TeamDto
// @Failure 400 {object} models.ErrorDto
// @Failure 404 {object} models.ErrorDto
// @Router /seasons/{seasonId}/teams/{teamId} [get]
func (c *TeamController) GetTeam(ctx *gin.Context) {
	seasonID, err := utils.GetParamUint(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	teamID, err := utils.GetParamUint(ctx, "teamId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	team, err := c.teamService.GetTeamByID(seasonID, teamID)
	if err != nil {
		r.NotFound(ctx, "Team not found in season")
		return
	}

	r.OK(ctx, team.ToDto())
}

// @Summary Create a team
// @Description Create a team
// @Tags teams
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID" default(2)
// @Param team body models.CreateTeamDto true "Team object that needs to be created"
// @Success 201 {object} models.TeamDto
// @Failure 400 {object} models.ErrorDto
// @Failure 401 {object} models.ErrorDto
// @Failure 404 {object} models.ErrorDto
// @Failure 422 {object} models.ErrorDto
// @Router /seasons/{seasonId}/teams [post]
func (c *TeamController) CreateTeam(ctx *gin.Context) {
	seasonID, err := utils.GetParamUint(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	var bodyDto models.CreateTeamDto
	if err = ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err = bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	season, err := c.seasonService.GetSeasonByID(seasonID)
	if err != nil {
		r.NotFound(ctx, "Season not found")
		return
	}

	user := utils.GetAuthUser(ctx)
	if user == nil {
		r.Unauthorized(ctx, "User not found")
		return
	}

	newTeam, err := c.teamService.CreateTeam(&bodyDto, season.ID, user.UUID)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.Created(ctx, newTeam.ToDto())
}

// @Summary Update a team
// @Description Update a team
// @Tags teams
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID" default(2)
// @Param teamId path string true "Team ID" default(2)
// @Param team body models.UpdateTeamDto true "Team object that needs to be updated"
// @Success 200 {object} models.TeamDto
// @Failure 400 {object} models.ErrorDto
// @Failure 401 {object} models.ErrorDto
// @Failure 403 {object} models.ErrorDto
// @Failure 404 {object} models.ErrorDto
// @Failure 422 {object} models.ErrorDto
// @Router /seasons/{seasonId}/teams/{teamId} [patch]
func (c *TeamController) UpdateTeam(ctx *gin.Context) {
	seasonID, err := utils.GetParamUint(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	teamID, err := utils.GetParamUint(ctx, "teamId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	var bodyDto models.UpdateTeamDto

	if err = ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err = bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	user := utils.GetAuthUser(ctx)
	if user == nil {
		r.Unauthorized(ctx, "User not found")
		return
	}

	currentTeam, err := c.teamService.GetTeamByID(seasonID, teamID)
	if err != nil {
		r.NotFound(ctx, "Team not found in season")
		return
	}

	if user.Role != models.AdminRole && user.Role != models.ModeratorRole && user.UUID != currentTeam.UserUUID {
		r.Forbidden(ctx, "This action is forbidden")
		return
	}

	updatedTeam, err := c.teamService.UpdateTeam(currentTeam, &bodyDto)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.OK(ctx, updatedTeam.ToDto())
}

// @Summary Delete a team
// @Description Delete a team
// @Tags teams
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID" default(1)
// @Param teamId path string true "Team ID" default(1)
// @Success 204
// @Failure 400 {object} models.ErrorDto
// @Failure 401 {object} models.ErrorDto
// @Failure 403 {object} models.ErrorDto
// @Failure 404 {object} models.ErrorDto
// @Failure 422 {object} models.ErrorDto
// @Router /seasons/{seasonId}/teams/{teamId} [delete]
func (c *TeamController) DeleteTeam(ctx *gin.Context) {
	seasonID, err := utils.GetParamUint(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	teamID, err := utils.GetParamUint(ctx, "teamId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	user := utils.GetAuthUser(ctx)
	if user == nil {
		r.Unauthorized(ctx, "User not found")
		return
	}

	team, err := c.teamService.GetTeamByID(seasonID, teamID)
	if err != nil {
		r.NotFound(ctx, "Team not found in season")
		return
	}

	if user.Role != models.AdminRole && user.UUID != team.UserUUID {
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
