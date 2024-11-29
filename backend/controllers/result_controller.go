package controllers

import (
	"MatchManiaAPI/models"
	"MatchManiaAPI/utils"

	r "MatchManiaAPI/responses"
	"MatchManiaAPI/services"

	"github.com/gin-gonic/gin"
)

type ResultController struct {
	teamService   services.TeamService
	resultService services.ResultService
}

func NewResultController(teamService services.TeamService, resultService services.ResultService) ResultController {
	return ResultController{teamService: teamService, resultService: resultService}
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
func (c *ResultController) GetAllResults(ctx *gin.Context) {
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

	resultModels, err := c.resultService.GetAllResults(seasonID, teamID)
	if err != nil {
		r.BadGateway(ctx, err.Error())
		return
	}

	r.OK(ctx, r.ResultsResponse{Results: models.ToResultDtos(resultModels)})
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
func (c *ResultController) GetResult(ctx *gin.Context) {
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

	resultID, err := utils.GetParamUint(ctx, "resultId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	resultModel, err := c.resultService.GetResultByID(seasonID, teamID, resultID)
	if err != nil {
		r.NotFound(ctx, "Result not found in given team and season")
		return
	}

	r.OK(ctx, r.ResultResponse{Result: resultModel.ToDto()})
}

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
func (c *ResultController) CreateResult(ctx *gin.Context) {
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

	var bodyDto models.CreateResultDto
	if err := ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	if bodyDto.OpponentTeamID == teamID {
		r.UnprocessableEntity(ctx, "Opponent Team ID cannot be the same as Team ID")
		return
	}

	team, err := c.teamService.GetTeamByID(seasonID, teamID)
	if err != nil {
		r.NotFound(ctx, "Team not found in given season")
		return
	}

	_, err = c.teamService.GetTeamByID(seasonID, bodyDto.OpponentTeamID)
	if err != nil {
		r.NotFound(ctx, "Opponent Team not found in given season")
		return
	}

	user := utils.GetAuthUser(ctx)
	if user == nil {
		r.Unauthorized(ctx, "User not found")
		return
	}

	newResult, err := c.resultService.CreateResult(&bodyDto, team.SeasonID, team.ID, user.UUID)
	if err != nil {
		r.BadGateway(ctx, err.Error())
		return
	}

	r.Created(ctx, r.ResultResponse{Result: newResult.ToDto()})
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
// @Router /seasons/{seasonId}/teams/{teamId}/results/{resultId} [patch]
func (c *ResultController) UpdateResult(ctx *gin.Context) {
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

	resultID, err := utils.GetParamUint(ctx, "resultId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	var bodyDto models.UpdateResultDto
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

	result, err := c.resultService.GetResultByID(seasonID, teamID, resultID)
	if err != nil {
		r.NotFound(ctx, "Result not found in given team and season")
		return
	}

	if user.Role != models.AdminRole && user.UUID != result.UserUUID {
		r.Forbidden(ctx, "This action is forbidden")
		return
	}

	updatedResult, err := c.resultService.UpdateResult(resultID, &bodyDto)
	if err != nil {
		r.BadGateway(ctx, err.Error())
		return
	}

	r.OK(ctx, r.ResultResponse{Result: updatedResult.ToDto()})
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
func (c *ResultController) DeleteResult(ctx *gin.Context) {
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

	resultID, err := utils.GetParamUint(ctx, "resultId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	user := utils.GetAuthUser(ctx)
	if user == nil {
		r.Unauthorized(ctx, "User not found")
		return
	}

	resultModel, err := c.resultService.GetResultByID(seasonID, teamID, resultID)
	if err != nil {
		r.NotFound(ctx, "Result not found in given team and season")
		return
	}

	if user.Role != models.AdminRole && user.UUID != resultModel.UserUUID {
		r.Forbidden(ctx, "This action is forbidden")
		return
	}

	err = c.resultService.DeleteResult(resultModel)
	if err != nil {
		r.BadGateway(ctx, err.Error())
		return
	}

	r.Deleted(ctx)
}
