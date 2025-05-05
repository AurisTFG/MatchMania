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
// @Success 200 {object} []responses.ResultDto
// @Failure 400 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /results [get]
func (c *ResultController) GetAllResults(ctx *gin.Context) {
	results, err := c.resultService.GetAllResults()
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	dto := utils.MustCopy[[]responses.ResultDto](results)

	r.OK(ctx, dto)
}

// @Summary Get a result
// @Description Get a result
// @Tags results
// @Accept json
// @Produce json
// @Param resultId path string true "Result ID" default(0deecf6a-289b-49a0-8f1b-9bc4185f99df)
// @Success 200 {object} []responses.ResultDto
// @Failure 400 {object} responses.ErrorDto
// @Failure 404 {object} responses.ErrorDto
// @Router /results/{resultId} [get]
func (c *ResultController) GetResult(ctx *gin.Context) {
	resultId, err := utils.GetParamId(ctx, "resultId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	result, err := c.resultService.GetResultById(resultId)
	if err != nil {
		r.NotFound(ctx, "Result not found in given team and league")
		return
	}

	dto := utils.MustCopy[responses.ResultDto](result)

	r.OK(ctx, dto)
}

// @Summary Create a result
// @Description Create a result
// @Tags results
// @Accept json
// @Produce json
// @Param result body requests.CreateResultDto true "Result object that needs to be created"
// @Success 201
// @Failure 400 {object} responses.ErrorDto
// @Failure 401 {object} responses.ErrorDto
// @Failure 404 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /results [post]
func (c *ResultController) CreateResult(ctx *gin.Context) {
	var bodyDto requests.CreateResultDto
	if err := ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err := validators.Validate(&bodyDto); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	if bodyDto.TeamId == bodyDto.OpponentTeamId {
		r.UnprocessableEntity(ctx, "Team and Opponent Team cannot be the same")
		return
	}

	userId := utils.MustGetRequestingUserId(ctx)

	if err := c.resultService.CreateResult(&bodyDto, userId); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.Created(ctx)
}

// @Summary Update a result
// @Description Update a result
// @Tags results
// @Accept json
// @Produce json
// @Param resultId path string true "Result ID" default(0deecf6a-289b-49a0-8f1b-9bc4185f99df)
// @Param result body requests.UpdateResultDto true "Result object that needs to be updated"
// @Success 204
// @Failure 400 {object} responses.ErrorDto
// @Failure 401 {object} responses.ErrorDto
// @Failure 403 {object} responses.ErrorDto
// @Failure 404 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /results/{resultId} [patch]
func (c *ResultController) UpdateResult(ctx *gin.Context) {
	resultId, err := utils.GetParamId(ctx, "resultId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	var bodyDto requests.UpdateResultDto
	if err = ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err = validators.Validate(&bodyDto); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	currentResult, err := c.resultService.GetResultById(resultId)
	if err != nil {
		r.NotFound(ctx, "Result not found in given team and league")
		return
	}

	if err = c.resultService.UpdateResult(currentResult, &bodyDto); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.NoContent(ctx)
}

// @Summary Delete a result
// @Description Delete a result
// @Tags results
// @Accept json
// @Produce json
// @Param resultId path string true "Result ID" default(0deecf6a-289b-49a0-8f1b-9bc4185f99df)
// @Success 204
// @Failure 400 {object} responses.ErrorDto
// @Failure 401 {object} responses.ErrorDto
// @Failure 403 {object} responses.ErrorDto
// @Failure 404 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /results/{resultId} [delete]
func (c *ResultController) DeleteResult(ctx *gin.Context) {
	resultId, err := utils.GetParamId(ctx, "resultId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	result, err := c.resultService.GetResultById(resultId)
	if err != nil {
		r.NotFound(ctx, "Result not found in given team and league")
		return
	}

	err = c.resultService.DeleteResult(result)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.NoContent(ctx)
}
