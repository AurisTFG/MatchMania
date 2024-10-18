package controllers

import (
	"MatchManiaAPI/models"
	"fmt"

	r "MatchManiaAPI/responses"
	"MatchManiaAPI/services"

	"github.com/gin-gonic/gin"
)

// @Summary Create a result
// @Description Create a result
// @Tags results
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID"
// @Param teamId path string true "Team ID"
// @Param result body models.CreateResultDto true "Result object that needs to be created"
// @Success 201 {object} models.ResultResponse
// @Failure 400 {object} models.BadRequestResponse
// @Failure 422 {object} models.UnprocessableEntityResponse
// @Failure 502 {object} models.BadGatewayResponse
// @Router /seasons/{seasonId}/teams/{teamId}/results [post]
func CreateResult(c *gin.Context) {
	seasonID := c.Param("seasonId")
	teamID := c.Param("teamId")
	var bodyDto models.CreateResultDto

	if err := c.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(c, err.Error())
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(c, err.Error())
		return
	}

	if fmt.Sprint(bodyDto.OpponentTeamID) == teamID {
		r.UnprocessableEntity(c, "Opponent Team ID cannot be the same as Team ID")
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
		r.BadGateway(c, err.Error())
		return
	}

	r.Created(c, models.ResultResponse{Result: newResult.ToDto()})
}

// @Summary Update a result
// @Description Update a result
// @Tags results
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID"
// @Param teamId path string true "Team ID"
// @Param resultId path string true "Result ID"
// @Param result body models.UpdateResultDto true "Result object that needs to be updated"
// @Success 200 {object} models.ResultResponse
// @Failure 400 {object} models.BadRequestResponse
// @Failure 404 {object} models.NotFoundResponse
// @Failure 502 {object} models.BadGatewayResponse
// @Router /seasons/{seasonId}/teams/{teamId}/results/{resultId} [put]
func UpdateResult(c *gin.Context) {
	seasonID := c.Param("seasonId")
	teamID := c.Param("teamId")
	resultID := c.Param("resultId")
	var bodyDto models.UpdateResultDto

	if err := c.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(c, err.Error())
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(c, err.Error())
		return
	}

	resultModel, err := services.GetResultByID(seasonID, teamID, resultID)
	if err != nil {
		r.NotFound(c, "Result with id "+resultID+" not found in team with id "+teamID+" in season with id "+seasonID)
		return
	}

	updatedResult, err := services.UpdateResult(resultModel, &bodyDto)
	if err != nil {
		r.BadGateway(c, err.Error())
		return
	}

	r.OK(c, models.ResultResponse{Result: updatedResult.ToDto()})
}

// @Summary Get a result
// @Description Get a result
// @Tags results
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID"
// @Param teamId path string true "Team ID"
// @Param resultId path string true "Result ID"
// @Success 200 {object} models.ResultResponse
// @Failure 404 {object} models.NotFoundResponse
// @Router /seasons/{seasonId}/teams/{teamId}/results/{resultId} [get]
func GetResult(c *gin.Context) {
	seasonID := c.Param("seasonId")
	teamID := c.Param("teamId")
	resultID := c.Param("resultId")

	resultModel, err := services.GetResultByID(seasonID, teamID, resultID)
	if err != nil {
		r.NotFound(c, "Result with id "+resultID+" not found in team with id "+teamID+" in season with id "+seasonID)
		return
	}

	r.OK(c, models.ResultResponse{Result: resultModel.ToDto()})
}

// @Summary Get all results
// @Description Get all results
// @Tags results
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID"
// @Param teamId path string true "Team ID"
// @Success 200 {object} models.ResultsResponse
// @Failure 502 {object} models.BadGatewayResponse
// @Router /seasons/{seasonId}/teams/{teamId}/results [get]
func GetAllResults(c *gin.Context) {
	seasonID := c.Param("seasonId")
	teamID := c.Param("teamId")

	resultModels, err := services.GetAllResults(seasonID, teamID)
	if err != nil {
		r.BadGateway(c, err.Error())
		return
	}

	r.OK(c, models.ResultsResponse{Results: models.ToResultDtos(resultModels)})
}

// @Summary Delete a result
// @Description Delete a result
// @Tags results
// @Accept json
// @Produce json
// @Param seasonId path string true "Season ID"
// @Param teamId path string true "Team ID"
// @Param resultId path string true "Result ID"
// @Success 204
// @Failure 404 {object} models.NotFoundResponse
// @Failure 502 {object} models.BadGatewayResponse
// @Router /seasons/{seasonId}/teams/{teamId}/results/{resultId} [delete]
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
		r.BadGateway(c, err.Error())
		return
	}

	r.Deleted(c)
}
