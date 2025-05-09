package controllers

import (
	"MatchManiaAPI/services"
	"MatchManiaAPI/utils"
	r "MatchManiaAPI/utils/httpresponses"

	"github.com/gin-gonic/gin"
)

type MatchController struct {
	matchService services.MatchService
}

func NewMatchController(
	matchService services.MatchService,
) MatchController {
	return MatchController{
		matchService: matchService,
	}
}

// @Summary Get all matches
// @Description Get all matches
// @Tags matches
// @Accept json
// @Produce json
// @Success 200 {object} []responses.MatchDto
// @Failure 422 {object} responses.ErrorDto
// @Router /matchmaking/matches [get]
func (c *MatchController) GetAllMatches(ctx *gin.Context) {
	matchesDto, err := c.matchService.GetAllMatches()

	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.OK(ctx, matchesDto)
}

// @Summary End a match
// @Description End a match
// @Tags matches
// @Accept json
// @Produce json
// @Success 204
// @Failure 422 {object} responses.ErrorDto
// @Router /matchmaking/matches/end [post]
func (c *MatchController) EndMatch(ctx *gin.Context) {
	userId := utils.MustGetRequestingUserId(ctx)

	err := c.matchService.EndMatch(userId)

	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.NoContent(ctx)
}
