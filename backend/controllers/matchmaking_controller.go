package controllers

import (
	"MatchManiaAPI/models/dtos/requests"
	"MatchManiaAPI/models/dtos/responses"
	"MatchManiaAPI/services"
	"MatchManiaAPI/utils"
	r "MatchManiaAPI/utils/httpresponses"

	"github.com/gin-gonic/gin"
)

type MatchmakingController struct {
	matchmakingService services.MatchmakingService
	userService        services.UserService
	teamService        services.TeamService
}

func NewMatchmakingController(
	matchmakingService services.MatchmakingService,
	userService services.UserService,
	teamService services.TeamService,
) MatchmakingController {
	return MatchmakingController{
		matchmakingService: matchmakingService,
		userService:        userService,
		teamService:        teamService,
	}
}

// @Summary Join matchmaking queue
// @Description Join matchmaking queue
// @Tags matchmaking
// @Accept json
// @Produce json
// @Param result body requests.JoinQueueDto true "Join Queue DTO"
// @Success 204
// @Failure 400 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /matchmaking/queue/join [post]
func (c *MatchmakingController) JoinQueue(ctx *gin.Context) {
	var bodyDto requests.JoinQueueDto
	if err := ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	err := c.matchmakingService.JoinQueue(bodyDto.LeagueId, bodyDto.TeamId)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.NoContent(ctx)
}

// @Summary Leave matchmaking queue
// @Description Leave matchmaking queue
// @Tags matchmaking
// @Accept json
// @Produce json
// @Param result body requests.LeaveQueueDto true "Leave Queue DTO"
// @Success 204
// @Failure 400 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /matchmaking/queue/leave [post]
func (c *MatchmakingController) LeaveQueue(ctx *gin.Context) {
	var bodyDto requests.LeaveQueueDto
	if err := ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	err := c.matchmakingService.LeaveQueue(bodyDto.LeagueId, bodyDto.TeamId)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.NoContent(ctx)
}

// @Summary Get matchmaking queue
// @Description Get matchmaking queue
// @Tags matchmaking
// @Accept json
// @Produce json
// @Success 200 {object} responses.QueueTeamsCountDto
// @Failure 400 {object} responses.ErrorDto
// @Failure 422 {object} responses.ErrorDto
// @Router /matchmaking/queue/teams-count/{leagueId} [get]
func (c *MatchmakingController) GetQueueTeamsCount(ctx *gin.Context) {
	leagueId, err := utils.GetParamId(ctx, "leagueId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	count, err := c.matchmakingService.GetQueueTeamsCount(leagueId)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	dto := responses.QueueTeamsCountDto{
		TeamsCount: count,
	}

	r.OK(ctx, dto)
}

// @Summary Check match status
// @Description Check match status
// @Tags matchmaking
// @Accept json
// @Produce json
// @Param teamId path string true "Team ID" default(0deecf6a-289b-49a0-8f1b-9bc4185f99df)
// @Success 200 {object} responses.MatchStatusDto
// @Failure 400 {object} responses.ErrorDto
// @Failure 404 {object} responses.ErrorDto
// @Router /matchmaking/queue/status/{teamId} [get]
func (c *MatchmakingController) CheckMatchStatus(ctx *gin.Context) {
	teamId, err := utils.GetParamId(ctx, "teamId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	isInMatch := c.matchmakingService.IsInMatch(teamId)

	dto := responses.MatchStatusDto{
		IsInMatch: isInMatch,
	}

	r.OK(ctx, dto)
}
