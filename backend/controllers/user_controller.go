package controllers

import (
	responses "MatchManiaAPI/models/dtos/responses/users"
	"MatchManiaAPI/services"
	"MatchManiaAPI/utils"
	r "MatchManiaAPI/utils/httpResponses"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
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
// @Success 200 {object} []responses.UserMinimalDto
// @Router /users [get]
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	var usersDto []responses.UserMinimalDto
	copier.Copy(&usersDto, users)

	r.OK(ctx, usersDto)
}

// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} responses.UserMinimalDto
// @Router /users/{id} [get]
func (c *UserController) GetUserById(ctx *gin.Context) {
	userId, err := utils.GetParamId(ctx, "userId")
	if err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	user, err := c.userService.GetUserById(userId)
	if err != nil {
		r.NotFound(ctx, err.Error())
		return
	}

	var userDto responses.UserMinimalDto
	copier.Copy(&userDto, user)

	r.OK(ctx, userDto)
}
