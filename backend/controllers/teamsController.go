package controllers

import (
	"MatchManiaAPI/models"
	r "MatchManiaAPI/responses"
	"MatchManiaAPI/services"

	"github.com/gin-gonic/gin"
)

func CreateTeam(c *gin.Context) {
	seasonID := c.Param("seasonId")
	var bodyDto models.CreateTeamDto

	if err := c.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(c, err)
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(c, err)
		return
	}

	season, err := services.GetSeasonByID(seasonID)
	if err != nil {
		r.NotFound(c, "Season with id "+seasonID+" not found")
		return
	}

	newTeam, err := services.CreateTeam(&bodyDto, season.ID)
	if err != nil {
		r.BadGateway(c, err)
		return
	}

	r.Created(c, "team", newTeam.ToDto())
}

func UpdateTeam(c *gin.Context) {
	seasonID := c.Param("seasonId")
	teamID := c.Param("teamId")
	var bodyDto models.UpdateTeamDto

	if err := c.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(c, err)
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(c, err)
		return
	}

	team, err := services.GetTeamByID(seasonID, teamID)
	if err != nil {
		r.NotFound(c, "Team with id "+teamID+" not found in season with id "+seasonID)
		return
	}

	updatedTeam, err := services.UpdateTeam(team, &bodyDto)
	if err != nil {
		r.BadGateway(c, err)
		return
	}

	r.OK(c, "team", updatedTeam.ToDto())
}

func GetTeam(c *gin.Context) {
	seasonID := c.Param("seasonId")
	teamID := c.Param("teamId")

	team, err := services.GetTeamByID(seasonID, teamID)
	if err != nil {
		r.NotFound(c, "Team with id "+teamID+" not found in season with id "+seasonID)
		return
	}

	r.OK(c, "team", team.ToDto())
}

func GetAllTeams(c *gin.Context) {
	var seasonId = c.Param("seasonId")

	teams, err := services.GetAllTeams(seasonId)
	if err != nil {
		r.BadGateway(c, err)
		return
	}

	r.OK(c, "teams", models.ToTeamDtos(teams))
}

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
		r.BadGateway(c, err)
		return
	}

	r.Deleted(c)
}
