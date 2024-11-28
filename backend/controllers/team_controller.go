package controllers

import (
	"MatchManiaAPI/models"
	"MatchManiaAPI/responses"
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

// @Summary Create a team
// @Description Create a team
// @Tags teams
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID"
// @Param team body models.CreateTeamDto true "Team object that needs to be created"
// @Success 201 {object} models.TeamResponse
// @Failure 400 {object} models.BadRequestResponse
// @Failure 422 {object} models.UnprocessableEntityResponse
// @Failure 502 {object} models.BadGatewayResponse
// @Router /seasons/{seasonId}/teams [post]
func (c *TeamController) CreateTeam(ctx *gin.Context) {
	seasonID, err := utils.ParseID(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	var bodyDto models.CreateTeamDto
	if err := ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	season, err := c.seasonService.GetSeasonByID(seasonID)
	if err != nil {
		r.NotFound(ctx, "Season not found")
		return
	}

	newTeam, err := c.teamService.CreateTeam(&bodyDto, season.ID)
	if err != nil {
		r.BadGateway(ctx, err.Error())
		return
	}

	r.Created(ctx, responses.TeamResponse{Team: newTeam.ToDto()})
}

// @Summary Update a team
// @Description Update a team
// @Tags teams
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID"
// @Param teamId path string true "Team ID"
// @Param team body models.UpdateTeamDto true "Team object that needs to be updated"
// @Success 200 {object} models.TeamResponse
// @Failure 400 {object} models.BadRequestResponse
// @Failure 404 {object} models.NotFoundResponse
// @Failure 502 {object} models.BadGatewayResponse
// @Router /seasons/{seasonId}/teams/{teamId} [put]
func (c *TeamController) UpdateTeam(ctx *gin.Context) {
	seasonID, err := utils.ParseID(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	teamID, err := utils.ParseID(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	var bodyDto models.UpdateTeamDto

	if err := ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	_, err = c.teamService.GetTeamByID(seasonID, teamID)
	if err != nil {
		r.NotFound(ctx, "Team not found in season")
		return
	}

	updatedTeam, err := c.teamService.UpdateTeam(teamID, &bodyDto)
	if err != nil {
		r.BadGateway(ctx, err.Error())
		return
	}

	r.OK(ctx, responses.TeamResponse{Team: updatedTeam.ToDto()})
}

// @Summary Get a team
// @Description Get a team
// @Tags teams
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID"
// @Param teamId path string true "Team ID"
// @Success 200 {object} models.TeamResponse
// @Failure 404 {object} models.NotFoundResponse
// @Router /seasons/{seasonId}/teams/{teamId} [get]
func (c *TeamController) GetTeam(ctx *gin.Context) {
	seasonID, err := utils.ParseID(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	teamID, err := utils.ParseID(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	team, err := c.teamService.GetTeamByID(seasonID, teamID)
	if err != nil {
		r.NotFound(ctx, "Team not found in season")
		return
	}

	r.OK(ctx, responses.TeamResponse{Team: team.ToDto()})
}

// @Summary Get all teams
// @Description Get all teams
// @Tags teams
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID"
// @Success 200 {object} models.TeamsResponse
// @Failure 502 {object} models.BadGatewayResponse
// @Router /seasons/{seasonId}/teams [get]
func (c *TeamController) GetAllTeams(ctx *gin.Context) {
	seasonID, err := utils.ParseID(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	teams, err := c.teamService.GetAllTeams(seasonID)
	if err != nil {
		r.BadGateway(ctx, err.Error())
		return
	}

	r.OK(ctx, responses.TeamsResponse{Teams: models.ToTeamDtos(teams)})
}

// @Summary Delete a team
// @Description Delete a team
// @Tags teams
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID"
// @Param teamId path string true "Team ID"
// @Success 204
// @Failure 404 {object} models.NotFoundResponse
// @Failure 502 {object} models.BadGatewayResponse
// @Router /seasons/{seasonId}/teams/{teamId} [delete]
func (c *TeamController) DeleteTeam(ctx *gin.Context) {
	seasonID, err := utils.ParseID(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	teamID, err := utils.ParseID(ctx, "seasonId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	team, err := c.teamService.GetTeamByID(seasonID, teamID)
	if err != nil {
		r.NotFound(ctx, "Team not found in season")
		return
	}

	err = c.teamService.DeleteTeam(team)
	if err != nil {
		r.BadGateway(ctx, err.Error())
		return
	}

	r.Deleted(ctx)
}
