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
// @Success 200 {object} responses.UsersResponse
// @Router /users [get]
func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.OK(ctx, r.UsersResponse{Users: models.ToUserDtos(users)})
}

// @Summary Get user by ID
// @Description Get user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} responses.UserResponse
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

// @Summary Update user
// @Description Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.UpdateUserDto true "User"
// @Success 200 {object} responses.UserResponse
// @Router /users/{id} [patch]
func (c *UserController) UpdateUser(ctx *gin.Context) {
	userID := utils.GetParamString(ctx, "userId")

	var bodyDto models.UpdateUserDto
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

	currentUser, err := c.userService.GetUserByID(userID)
	if err != nil {
		r.NotFound(ctx, "User not found")
		return
	}

	if user.Role != models.AdminRole && user.UUID != currentUser.UUID {
		r.Forbidden(ctx, "Forbidden")
		return
	}

	updatedUser, err := c.userService.UpdateUser(currentUser, &bodyDto)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.OK(ctx, updatedUser.ToDto())
}
