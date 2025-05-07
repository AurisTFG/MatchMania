package controllers

import (
	"MatchManiaAPI/models/dtos/requests"
	"MatchManiaAPI/services"
	r "MatchManiaAPI/utils/httpresponses"

	"github.com/gin-gonic/gin"
)

type QueueController struct {
	queueService services.QueueService
}

func NewQueueController(
	queueService services.QueueService,
) QueueController {
	return QueueController{
		queueService: queueService,
	}
}

// @Summary Get all queues
// @Description Get all queues
// @Tags queues
// @Accept json
// @Produce json
// @Success 200 {object} []responses.QueueDto
// @Failure 422 {object} responses.ErrorDto
// @Router /matchmaking/queues [get]
func (c *QueueController) GetAllQueues(ctx *gin.Context) {
	queuesDto, err := c.queueService.GetAllQueues()

	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.OK(ctx, queuesDto)
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
// @Router /matchmaking/queues/join [post]
func (c *QueueController) JoinQueue(ctx *gin.Context) {
	var bodyDto requests.JoinQueueDto
	if err := ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	err := c.queueService.JoinQueue(bodyDto.LeagueId, bodyDto.TeamId)
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
// @Router /matchmaking/queues/leave [post]
func (c *QueueController) LeaveQueue(ctx *gin.Context) {
	var bodyDto requests.LeaveQueueDto
	if err := ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	err := c.queueService.LeaveQueue(bodyDto.LeagueId, bodyDto.TeamId)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.NoContent(ctx)
}
