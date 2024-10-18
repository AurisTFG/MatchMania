package controllers

import (
	"MatchManiaAPI/models"
	"fmt"

	r "MatchManiaAPI/responses"
	"MatchManiaAPI/services"

	"github.com/gin-gonic/gin"
)

func CreateResult(c *gin.Context) {
	seasonID := c.Param("seasonId")
	teamID := c.Param("teamId")
	var bodyDto models.CreateResultDto

	if err := c.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(c, err)
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(c, err)
		return
	}

	if fmt.Sprint(bodyDto.OpponentTeamID) == teamID {
		r.UnprocessableEntity(c, fmt.Errorf("opponent Team ID cannot be the same as Team ID"))
		return
	}

	team, err := services.GetTeamByID(seasonID, teamID)
	if err != nil {
		r.NotFound(c, "Team with id "+teamID+" not found in season with id "+seasonID)
		return
	}

	_, err = services.GetTeamByID(seasonID, fmt.Sprint(bodyDto.OpponentTeamID))
	if err != nil {
		r.NotFound(c, "Opponent Team with id "+fmt.Sprint(bodyDto.OpponentTeamID)+" not found in season with id "+seasonID)
		return
	}

	newResult, err := services.CreateResult(&bodyDto, team.SeasonID, team.ID)
	if err != nil {
		r.BadGateway(c, err)
		return
	}

	r.Created(c, "result", newResult.ToDto())
}

func UpdateResult(c *gin.Context) {
	seasonID := c.Param("seasonId")
	teamID := c.Param("teamId")
	resultID := c.Param("resultId")
	var bodyDto models.UpdateResultDto

	if err := c.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(c, err)
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(c, err)
		return
	}

	resultModel, err := services.GetResultByID(seasonID, teamID, resultID)
	if err != nil {
		r.NotFound(c, "Result with id "+resultID+" not found in team with id "+teamID+" in season with id "+seasonID)
		return
	}

	updatedResult, err := services.UpdateResult(resultModel, &bodyDto)
	if err != nil {
		r.BadGateway(c, err)
		return
	}

	r.OK(c, "result", updatedResult.ToDto())
}

func GetResult(c *gin.Context) {
	seasonID := c.Param("seasonId")
	teamID := c.Param("teamId")
	resultID := c.Param("resultId")

	resultModel, err := services.GetResultByID(seasonID, teamID, resultID)
	if err != nil {
		r.NotFound(c, "Result with id "+resultID+" not found in team with id "+teamID+" in season with id "+seasonID)
		return
	}

	r.OK(c, "result", resultModel.ToDto())
}

func GetAllResults(c *gin.Context) {
	seasonID := c.Param("seasonId")
	teamID := c.Param("teamId")

	resultModels, err := services.GetAllResults(seasonID, teamID)
	if err != nil {
		r.BadGateway(c, err)
		return
	}

	r.OK(c, "results", models.ToResultDtos(resultModels))
}

func DeleteResult(c *gin.Context) {
	seasonID := c.Param("seasonId")
	teamID := c.Param("teamId")
	resultID := c.Param("resultId")

	resultModel, err := services.GetResultByID(seasonID, teamID, resultID)
	if err != nil {
		r.NotFound(c, "Result with id "+resultID+" not found in team with id "+teamID+" in season with id "+seasonID)
		return
	}

	err = services.DeleteResult(resultModel)
	if err != nil {
		r.BadGateway(c, err)
		return
	}

	r.Deleted(c)
}
