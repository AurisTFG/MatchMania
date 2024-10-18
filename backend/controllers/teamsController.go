package controllers

import (
	"MatchManiaAPI/models"
	r "MatchManiaAPI/responses"
	"MatchManiaAPI/services"

	"github.com/gin-gonic/gin"
)

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
func CreateTeam(c *gin.Context) {
	seasonID := c.Param("seasonId")
	var bodyDto models.CreateTeamDto

	if err := c.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(c, err.Error())
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(c, err.Error())
		return
	}

	season, err := services.GetSeasonByID(seasonID)
	if err != nil {
		r.NotFound(c, "Season with id "+seasonID+" not found")
		return
	}

	newTeam, err := services.CreateTeam(&bodyDto, season.ID)
	if err != nil {
		r.BadGateway(c, err.Error())
		return
	}

	r.Created(c, models.TeamResponse{Team: newTeam.ToDto()})
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
func UpdateTeam(c *gin.Context) {
	seasonID := c.Param("seasonId")
	teamID := c.Param("teamId")
	var bodyDto models.UpdateTeamDto

	if err := c.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(c, err.Error())
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(c, err.Error())
		return
	}

	team, err := services.GetTeamByID(seasonID, teamID)
	if err != nil {
		r.NotFound(c, "Team with id "+teamID+" not found in season with id "+seasonID)
		return
	}

	updatedTeam, err := services.UpdateTeam(team, &bodyDto)
	if err != nil {
		r.BadGateway(c, err.Error())
		return
	}

	r.OK(c, models.TeamResponse{Team: updatedTeam.ToDto()})
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
func GetTeam(c *gin.Context) {
	seasonID := c.Param("seasonId")
	teamID := c.Param("teamId")

	team, err := services.GetTeamByID(seasonID, teamID)
	if err != nil {
		r.NotFound(c, "Team with id "+teamID+" not found in season with id "+seasonID)
		return
	}

	r.OK(c, models.TeamResponse{Team: team.ToDto()})
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
func GetAllTeams(c *gin.Context) {
	var seasonId = c.Param("seasonId")

	teams, err := services.GetAllTeams(seasonId)
	if err != nil {
		r.BadGateway(c, err.Error())
		return
	}

	r.OK(c, models.TeamsResponse{Teams: models.ToTeamDtos(teams)})
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
func DeleteTeam(c *gin.Context) {
	seasonID := c.Param("seasonId")
	teamID := c.Param("teamId")

	team, err := services.GetTeamByID(seasonID, teamID)
	if err != nil {
		r.NotFound(c, "Team with id "+teamID+" not found in season with id "+seasonID)
		return
	}

	err = services.DeleteTeam(team)
	if err != nil {
		r.BadGateway(c, err.Error())
		return
	}

	r.Deleted(c)
}
