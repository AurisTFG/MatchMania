package controllers

import (
	"MatchManiaAPI/models/dtos/responses"
	"MatchManiaAPI/services"
	"MatchManiaAPI/utils"
	r "MatchManiaAPI/utils/httpresponses"

	"github.com/gin-gonic/gin"
)

type PlayerController struct {
	playerService services.PlayerService
}

func NewPlayerController(playerService services.PlayerService) PlayerController {
	return PlayerController{playerService: playerService}
}

// @Summary Get all players
// @Description Get all players
// @Tags players
// @Accept json
// @Produce json
// @Success 200 {object} []responses.PlayerMinimalDto
// @Failure 422 {object} responses.ErrorDto
// @Router /players [get]
func (c *PlayerController) GetAllPlayers(ctx *gin.Context) {
	players, err := c.playerService.GetAllPlayers()

	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	dto := utils.MustCopy[[]responses.PlayerMinimalDto](players)

	r.OK(ctx, dto)
}
