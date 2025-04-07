package controllers

import (
	"MatchManiaAPI/models"
	r "MatchManiaAPI/responses"
	"MatchManiaAPI/services"
	"MatchManiaAPI/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{userService: userService}
}

// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} []models.UserDto
// @Router /users [get]
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.OK(ctx, models.ToUserDtos(users))
}

// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.UserDto
// @Router /users/{id} [get]
func (c *UserController) GetUserByID(ctx *gin.Context) {
	userID := utils.GetParamString(ctx, "userId")

	user, err := c.userService.GetUserByID(userID)
	if err != nil {
		r.NotFound(ctx, err.Error())
		return
	}

	r.OK(ctx, user.ToDto())
}
