package controllers

import (
	"MatchManiaAPI/config"
	"MatchManiaAPI/constants"
	"MatchManiaAPI/models"
	r "MatchManiaAPI/responses"
	"MatchManiaAPI/services"
	"MatchManiaAPI/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthController struct {
	authService services.AuthService
	userService services.UserService
	env         *config.Env
}

func NewAuthController(
	authService services.AuthService,
	userService services.UserService,
	env *config.Env,
) AuthController {
	return AuthController{
		authService: authService,
		userService: userService,
		env:         env,
	}
}

// @Summary Sign up
// @Description Sign up
// @Tags auth
// @Accept json
// @Produce json
// @Param signUpDto body models.SignUpDto true "Sign up details"
// @Success 201 {object} models.UserDto
// @Failure 400 {object} models.ErrorDto
// @Failure 422 {object} models.ErrorDto
// @Router /auth/signup [post]
func (c *AuthController) SignUp(ctx *gin.Context) {
	var bodyDto models.SignUpDto

	if err := ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	user, err := c.userService.CreateUser(&bodyDto)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	r.Created(ctx, user.ToDto())
}

// @Summary Log in
// @Description Log in
// @Tags auth
// @Accept json
// @Produce json
// @Param loginDto body models.LoginDto true "Log in details"
// @Success 200 {object} models.UserDto
// @Failure 400 {object} models.ErrorDto
// @Failure 422 {object} models.ErrorDto
// @Router /auth/login [post]
func (c *AuthController) LogIn(ctx *gin.Context) {
	var bodyDto models.LoginDto

	if err := ctx.ShouldBindJSON(&bodyDto); err != nil {
		r.BadRequest(ctx, err.Error())
		return
	}

	if err := bodyDto.Validate(); err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	user, err := c.userService.GetUserByEmail(bodyDto.Email)
	if err != nil {
		r.UnprocessableEntity(ctx, "Username or password was incorrect.")
		return
	}

	if !user.ComparePassword(bodyDto.Password) {
		r.UnprocessableEntity(ctx, "Username or password was incorrect.")
		return
	}

	sessionUUID := uuid.New()

	accessToken, err := c.authService.CreateAccessToken(user)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	refreshToken, err := c.authService.CreateRefreshToken(sessionUUID.String(), user)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	err = c.authService.CreateSession(sessionUUID, user.UUID, refreshToken)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	c.authService.SetCookies(ctx, accessToken, refreshToken)

	r.OK(ctx, user.ToDto())
}

// @Summary Log out
// @Description Log out
// @Tags auth
// @Success 204
// @Failure 422 {object} models.ErrorDto
// @Router /auth/logout [post]
func (c *AuthController) LogOut(ctx *gin.Context) {
	tokenString, err := ctx.Cookie(constants.RefreshTokenName)
	if err != nil {
		r.UnprocessableEntity(ctx, "Already logged out")
		c.authService.DeleteCookies(ctx)
		return
	}

	_, sessionId, err := c.authService.VerifyRefreshToken(tokenString)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		c.authService.DeleteCookies(ctx)
		return
	}

	if !c.authService.IsSessionValid(sessionId, tokenString) {
		r.UnprocessableEntity(ctx, "Session is not valid")
		c.authService.DeleteCookies(ctx)
		return
	}

	err = c.authService.InvalidateSession(sessionId)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		c.authService.DeleteCookies(ctx)
		return
	}

	c.authService.DeleteCookies(ctx)

	r.NoContent(ctx)
}

// @Summary Refresh token
// @Description Refresh token
// @Tags auth
// @Success 204
// @Failure 422 {object} models.ErrorDto
// @Router /auth/refresh [post]
func (c *AuthController) RefreshToken(ctx *gin.Context) {
	tokenString, err := ctx.Cookie(constants.RefreshTokenName)
	if err != nil {
		r.UnprocessableEntity(ctx, "Refresh token not found")
		return
	}

	user, sessionId, err := c.authService.VerifyRefreshToken(tokenString)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	if !c.authService.IsSessionValid(sessionId, tokenString) {
		r.UnprocessableEntity(ctx, "Session is not valid")
		return
	}

	accessToken, err := c.authService.CreateAccessToken(user)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	refreshToken, err := c.authService.CreateRefreshToken(sessionId, user)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	err = c.authService.ExtendSession(sessionId, refreshToken)
	if err != nil {
		r.UnprocessableEntity(ctx, err.Error())
		return
	}

	c.authService.SetCookies(ctx, accessToken, refreshToken)

	r.NoContent(ctx)
}

// @Summary Get current user
// @Description Get current user
// @Tags auth
// @Success 200 {object} models.UserDto
// @Failure 422 {object} models.ErrorDto
// @Router /auth/me [get]
func (c *AuthController) GetMe(ctx *gin.Context) {
	fmt.Println("GetMe called")

	user := utils.GetAuthUser(ctx)
	if user == nil {
		r.Unauthorized(ctx, "User not found")
		return
	}

	r.OK(ctx, user.ToDto())
}
